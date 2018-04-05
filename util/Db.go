package util

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var (
	MongoDBHosts = "mongodb1318.aws-us-east-1-portal.19.dblayer.com:11318"
	Database     = "jbndevdb"
	DBUserName   = "jbn"
	DBPassword   = "jbn12345$"

//	AuthDatabase = "authdb"
//	AuthUserName = "jbnroot"
//	AuthPassword = "jbnpass"
)

var (
	mainSession *mgo.Session
	mainDb      *mgo.Database
)

type MgoDb struct {
	Session *mgo.Session
	Db      *mgo.Database
	Col     *mgo.Collection
}

func init() {
	setDBEnv()

	if mainSession == nil {

		var err error
		//mainSession, err = mgo.Dial(Host)
		// We need this object to establish a session to our MongoDB.
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{MongoDBHosts},
			Timeout:  60 * time.Second,
			Database: Database,
			Username: DBUserName,
			Password: DBPassword,
		}
		// Create a session which maintains a pool of socket connections
		// to our MongoDB.
		mainSession, err := mgo.DialWithInfo(mongoDBDialInfo)
		if err != nil {
			log.Fatalf("CreateSession: %s\n", err)
		}

		mainSession.SetMode(mgo.Monotonic, true)
		mainDb = mainSession.DB(Database)
	}

}

func setDBEnv() {
	//	env = os.GetEnv("ENV")
}

func (this *MgoDb) Init() *mgo.Session {

	this.Session = mainSession.Copy()
	this.Db = this.Session.DB(Database)

	return this.Session
}

func (this *MgoDb) C(collection string) *mgo.Collection {
	this.Col = this.Session.DB(Database).C(collection)
	return this.Col
}

func (this *MgoDb) Close() bool {
	defer this.Session.Close()
	return true
}

func (this *MgoDb) DropoDb() {
	err := this.Session.DB(Database).DropDatabase()
	if err != nil {
		panic(err)
	}
}

func (this *MgoDb) RemoveAll(collection string) bool {
	this.Session.DB(Database).C(collection).RemoveAll(nil)

	this.Col = this.Session.DB(Database).C(collection)
	return true
}

func (this *MgoDb) Index(collection string, keys []string) bool {

	index := mgo.Index{
		Key:        keys,
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := this.Db.C(collection).EnsureIndex(index)
	if err != nil {
		panic(err)

		return false
	}

	return true
}

func (this *MgoDb) IsDup(err error) bool {

	if mgo.IsDup(err) {
		return true
	}

	return false
}

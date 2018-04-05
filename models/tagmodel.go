package models

import (
	"encoding/json"
	"fmt"
	//	"strings"
	//	"time"
	"strconv"

	"github.com/joybynature/jbnserverapp/util"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type QATags struct {
	TAGID         bson.ObjectId `json:"_id" bson:"_id"`
	TAGNAME       string        `json:"tagname" bson:"tagname"`
	TAGCODE       string        `json:"tagcode" bson:"tagcode"`
	TAGLEVEL      int           `json:"taglevel" bson:"taglevel"`
	TAGPARENTCODE string        `json:"tagparentcode" bson:"tagparentcode"`
	// ImageURL          string               `json:"imageurl" bson:"imageurl"`
}

type SHOPTags struct {
	TAGID         bson.ObjectId `json:"_id" bson:"_id"`
	TAGNAME       string        `json:"tagname" bson:"tagname"`
	TAGCODE       string        `json:"tagcode" bson:"tagcode"`
	TAGLEVEL      int           `json:"taglevel" bson:"taglevel"`
	TAGPARENTCODE string        `json:"tagparentcode" bson:"tagparentcode"`
	COLLECTIONID  int           `json:"cid" bson:"cid"`
	// ImageURL          string               `json:"imageurl" bson:"imageurl"`
}

type ShopTagCode struct {
	TAGCODE string `json:"tagcode" bson:"tagcode"`
}

type TagsPool struct {
	TagID   bson.ObjectId `json:"_id" bson:"_id"`
	Tagname string        `json:"tagname" bson:"tagname"`
}

func GetTagCategory(c *iris.Context, mongoSession *mgo.Session) {
	tagcode := c.Param("tagparentcode")
	taglevel, _ := strconv.Atoi(c.Param("taglevel"))

	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var qatag []QATags
	collection := sessionCopy.DB(util.DATABASENAME).C("qatags")

	err := collection.Find(bson.M{"$and": []interface{}{bson.M{"tagparentcode": tagcode}, bson.M{"taglevel": taglevel}}}).All(&qatag)
	if err != nil {
		fmt.Println("Get qatag  err " + err.Error())
	}
	tagCategoryResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": qatag}

	resp, err := json.Marshal(tagCategoryResult)
	if err != nil {
		fmt.Println("Get qatag category REsult " + err.Error())
	}

	c.Write(string(resp))
}

func GetShopTag(c *iris.Context, mongoSession *mgo.Session) {
	tagParentName := c.Param("tagparentname")
	taglevel, _ := strconv.Atoi(c.Param("taglevel"))

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	shoptag := SHOPTags{}
	var shoptags []SHOPTags

	collection := sessionCopy.DB(util.DATABASENAME).C("shoptags")

	//err := collection.Find(bson.M{"$and":[]interface{}{bson.M{"tagname":tagParentName},bson.M{"taglevel":1}}}).One(&shoptag)
	err := collection.Find(bson.M{"tagname": tagParentName}).One(&shoptag)
	if err != nil {
		fmt.Println("Get shoptag  err " + err.Error())
	}
	shop_tag := shoptag.TAGCODE
	fmt.Println("shop tag", shop_tag)
	err = collection.Find(bson.M{"$and": []interface{}{bson.M{"tagparentcode": shop_tag}, bson.M{"taglevel": taglevel}}}).All(&shoptags)
	if err != nil {
		fmt.Println("Get shoptag  err " + err.Error())
	}
	ShopTagsResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": shoptags}
	resp, err := json.Marshal(ShopTagsResult)
	if err != nil {
		fmt.Println("Get shoptag category REsult " + err.Error())
	}

	c.Write(string(resp))

}

func GetPoolTags(c *iris.Context, mongoSession *mgo.Session) (tags string) {
	tagName := c.Param("tagname")

	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var tagspool []TagsPool
	collection := sessionCopy.DB(util.DATABASENAME).C("tagspool")

	//The Tag Pool Query to get tags based on the search critera
	stag := bson.M{"$regex": bson.RegEx{Pattern: tagName, Options: "i"}}
	opMatch := bson.M{"$match": bson.M{"tagname": stag}}
	opProject := bson.M{"$project": bson.M{"_id": 0, "tagname": 1}}

	opMain := []bson.M{opMatch, opProject}
	pipe := collection.Pipe(opMain)
	err := pipe.All(&tagspool)
	if err != nil {
		fmt.Println("Error Searching TAG POOL" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_QA502, "message": "Error", "result": "Error find the tags from the tag pool"}
		errResp, _ := json.Marshal(errRsponse)

		return string(errResp)
	}

	sTagResult := map[string]interface{}{"code": util.CODE_TAG001, "message": "Success", "result": tagspool}
	resp, err := json.Marshal(sTagResult)
	if err != nil {
		fmt.Println("Error in creating JSON format for TAG POOL " + err.Error())
	}

	return string(resp)
}

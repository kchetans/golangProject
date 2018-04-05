package models

import (
	 "bytes"
    "net/http"
    "strconv"
	"encoding/json"
	"fmt"
//	"strings"
	 "net/url"
	//"strings"
	//"crypto/md5"
	"encoding/hex"
	"regexp"
	//"strconv"
	"time"
	"math/rand"
	"github.com/dgrijalva/jwt-go"
	"github.com/joybynature/jbnserverapp/util"
	"github.com/joybynature/jbnserverapp/config"
	"github.com/kataras/iris"
	"golang.org/x/crypto/scrypt"
	//	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    //"github.com/NaySoftware/go-fcm"
//	"os/user"
 "github.com/NaySoftware/go-fcm"
)

type Author struct {
	UserID         bson.ObjectId `json:"userid" bson:"userid"`
	Name           string        `json:"name" bson:"name"`
	ImageURL       string        `json:"imageurl" bson:"imageurl"`
	Followers      int           `json:"followers" bson:"followers"`
	SingleLineDesc string        `json:"singlelinedesc" bson:"singlelinedesc"`
}

type User struct {
	UserID       bson.ObjectId `json:"_id" bson:"_id"`
	UserName     string        `json:"username" bson:"username"`
	FirstName    string        `json:"firstname" bson:"firstname"`
	LastName     string        `json:"lastname" bson:"lastname"`
	Pictureurl   string        `json:"pictureurl" bson:"pictureurl"`
	Description  string        `json:"description" bson:"description"`
	AboutUS      string        `json:"aboutus" bson:"aboutus"`
	EmailId      string        `json:"emailid" bson:"emailid"`
	Password     string        `json:"password" bson:"password"`
	Gender       string        `json:"gender" bson:"gender"`
	Dob          time.Time     `json:"dob" bson:"dob`
	Phone        string        `json:"phone" bson:"phone"`
	Scope        string        `json:"scope" bson:"scope"`
	ShrinkStatus bool			`json:"shrinkstatus" bson:"shrinkstatus"`
	Seed		 bool			`json:"seed" bson:"seed"`
	Location     []Location    `json:"location" bson:"location"`
	SubscriberID string        `json:"subscriberid" bson:"subscriberid"`
	MyQuestions  int	   	  `json:"myquestions" bson:"myquestions"`
	MyAnswers    int		    `json:"myanswers" bson:"myanswers"`
	FollowersCount  int  	   `json:"followerscount" bson:"followerscount"`
	MyStories    string        `json:"mystories" bson:"mystories"`
	Status       string        `json:"status" bson:"status"`
	DateOfJoin   time.Time     `json:"dateofjoin" bson:"dateofjoin"`
	DateRemoved  time.Time     `json:"dateremoved" bson:"dateremoved"`
}
type EditProfileForm struct{
	UserName     string      
	ImageURL   	 string        
	Description  string 
	AboutUS		 string       
	Gender       string       
	Dob          time.Time    
	//Phone        string       
//	Location     []Location    `json:"location" bson:"location"`
}

type UserPublicProfile struct{
	UserID       bson.ObjectId `json:"_id" bson:"_id"`
	UserName     string        `json:"username" bson:"username"`
	FirstName    string        `json:"firstname" bson:"firstname"`
	LastName     string        `json:"lastname" bson:"lastname"`
	AboutUS      string        `json:"aboutus" bson:"aboutus"`
	Pictureurl   string        `json:"pictureurl" bson:"pictureurl"`
	Description  string        `json:"description" bson:"description"`
	Location     []Location    `json:"location" bson:"location"`
	MyQuestions  int  		`json:"myquestions" bson:"myquestions"`
	MyAnswers    int    		`json:"myanswers" bson:"myanswers"`
	FollowersCount  int  	   `json:"followerscount" bson:"followerscount"`
	
}
type UserPrivateProfile struct{
	UserID       bson.ObjectId `json:"_id" bson:"_id"`
	UserName     string        `json:"username" bson:"username"`
	FirstName    string        `json:"firstname" bson:"firstname"`
	LastName     string        `json:"lastname" bson:"lastname"`
	Pictureurl   string        `json:"pictureurl" bson:"pictureurl"`
	Description  string        `json:"description" bson:"description"`
	AboutUS      string        `json:"aboutus" bson:"aboutus"`
	Location     []Location    `json:"location" bson:"location"`
	MyQuestions  int  			`json:"myquestions" bson:"myquestions"`
	MyAnswers    int   		   `json:"myanswers" bson:"myanswers"`
	FollowersCount  int  	   `json:"followerscount" bson:"followerscount"`
	EmailId      string        `json:"emailid" bson:"emailid"`
	//Password     string        `json:"password" bson:"password"`
	Gender       string        `json:"gender" bson:"gender"`
	Dob          time.Time     `json:"dob" bson:"dob`
	Phone        string        `json:"phone" bson:"phone"`
	//Scope        string        `json:"scope" bson:"scope"`
}
//type Location struct {
//	Address     string `json:"address" bson:"address"`
//	City        string `json:"city" bson:"city"`
//	State       string `json:"state" bson:"state"`
//	Zip         string `json:"zip" bson:"zip"`
//	CountryCode string `json:"countrycode" bson:"countrycode"`
//	Pincode     string `json:"pincode" bson:"pincode"`
//	Latitude    string `json:"latitude" bson:"latitude"`
//	Longitude   string `json:"Longitude" bson:"Longitude"`
//}
type Location struct {
	AddressID	bson.ObjectId `json:"addressid" bson:"addressid"`
	Address1    string 		  `json:"address1" bson:"address1"`
	Address2    string		  `json:"address2" bson:"address2"`
	City        string		  `json:"city" bson:"city"`
	State       string 		  `json:"state" bson:"state"`
	Zip         string 		  `json:"zip" bson:"zip"`
	Country	    string 		  `json:"country" bson:"country"`
	CountryCode string 		  `json:"countrycode" bson:"countrycode"`
	DefaultAddress string 	  `json:"defaultaddress" bson:"defaultaddress"`	
	Province 	string 	  	  `json:"province" bson:"province"`	
	ProvinceCode string 	  `json:"provincecode" bson:"provincecode"`	
}

type LocationForm struct {
	AddressID		string
	Address1    	string 		  
	Address2    	string		 
	City        	string		 
	State       	string 		  
	Zip         	string 		
	Country	    	string 		
	CountryCode		string 		
	DefaultAddress 	string 	  
	Province 		string 	  	  
	ProvinceCode 	string 	  
}

type MyQuestion struct {
	QuestionId bson.ObjectId `json:"questionid" bson:"questionid"`
}

type MyAnswer struct {
	QuestionId bson.ObjectId `json:"questionid" bson:"questionid"`
	AnswerId   bson.ObjectId `json:"answerid" 	bson:"answerid"`
}

//type MyFollower struct {
//	MyFollowerUsers	[]MyFollowerUsers `json:"followerid"	bson:"followerid"`
//}
//type MyFollowerUsers struct{
//	UserId bson.ObjectId `json:"userid"	bson:"userid"`
//}
type OTPS struct {
	OtpId          bson.ObjectId `json:"_id"	bson:"_id"`
	OTP            string	        `json:"otp" bson:"otp"`
	CreateDateTime time.Time     `json:"createdatetime" bson:"createdatetime"`
	Mobile    		string 		 `json:"mobile" form:"mobile"`
}

type PubUsers struct {
	UserId         bson.ObjectId `json:"userid" bson:"userid"`
	PubUserId      bson.ObjectId `json:"pubuserid" bson:"pubuserid"`
	CreateDateTime time.Time     `json:"ceratedatetime" bson:"createdatetime"`
}

type UserLoginForm struct {
	EmailID  string `json:"emailid" form:"emailid"`
	Password string `json:"password" form:"password"`
}

type SignUpForm struct {
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	Mobile    string `json:"mobile" form:"mobile"`
	EmailID   string `json:"emailid" form:"emailid"`
	Password  string `json:"password" form:"password"`
	SubscriberID string        `json:"subscriberid" bson:"subscriberid"`

}
type OtpForm struct{
	Mobile    string `json:"mobile" form:"mobile"`
}

type UserLoginInfo struct {
	UserID      string `json:"useid"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Description string `json:"description"`
	Mobile      string `json:"mobile"`
	AvatarURL   string `json:"avatarurl"`
	EmailID     string `json:"emailid"`
	Scope       string `json:"scope"`
}

type TokenAccess struct {
	Accesstoken string `json:"accesstoken"`
}

type NewUserStatus struct {
	Status string `json:"status"`
	Otp    string `json:"otp"`
}

type VerifyOTPForm struct{
	Otp    		string 
	Mobile    	string 		
	EmailID		string
}
type Followers struct{
	ID					 bson.ObjectId 			`json:"_id" bson:"_id"	`
	UserId				 bson.ObjectId 			`json:"userid" bson:"userid"`
	FollowerSubscriber	[]FollowerSubscriberID  `json:"followersubscriberid" bson:"followersubscriberid"`
	Name         		  string        		`json:"name" bson:"name"`
	ImageURL      		  string       			`json:"imageurl" bson:"imageurl"`
	SubscriberID 		 string        			`json:"subscriberid" bson:"subscriberid"`

}

type FollowerSubscriberID struct{
	FollowerID		 	 bson.ObjectId  		`json:"followerid" bson:"followerid"`
	SubscriberID 		 string        			`json:"subscriberid" bson:"subscriberid"`
	Name          		 string       			 `json:"name" bson:"name"`
	ImageURL      		 string       			 `json:"imageurl" bson:"imageurl"`
}


type FollowersForm struct{
	FollowerId		 string
	FollowerType		 string
}	

type EmailForm struct{
	EmailID	string
}
type ForgetPassword struct{
	Password 	string
}
func LoginUser(c *iris.Context, mongoSession *mgo.Session) (loginUser string) {
	//Get login form data
	u := UserLoginForm{}

	err := c.ReadForm(&u)
	if err != nil {
		fmt.Println("Error reading login form" + err.Error())
	}

	//Vaidate the login form data
	isValid, errMsg := validateLoginUser(u)

	if !isValid {
		loginValidResult := map[string]interface{}{"code": util.CODE_USR102, "message": "Error", "result": errMsg}
		loginResp, err := json.Marshal(loginValidResult)
		if err != nil {
			fmt.Println("Error in login form data" + err.Error())
		}
		return string(loginResp)
	}
	
	//Check if valid user
//	//check user status
//	userExistResult,userStatus	:= isUserExists(u.EmailID, mongoSession)
//	fmt.Println("userExistResult",userExistResult)
//	fmt.Println("userStatus",userStatus)
//	//Check if user exists
//	if userExistResult != false {
//		userExistsResult := map[string]interface{}{"code": util.CODE_USR106, "message": "Error", "result": userStatus}
//		userResp, err := json.Marshal(userExistsResult)
//		if err != nil {
//			fmt.Println("Error - The user already exists" + err.Error())
//		}
//		return string(userResp)
//	}
	encryptPass, err := scrypt.Key([]byte(u.Password), util.SALT, 16384, 8, 1, 32)
	if err != nil {
		fmt.Println("Error using scrypt.key" + err.Error())
		return
	}
	fmt.Println("PASSWORD - PLAIN:>", u.Password)
	fmt.Println("PASSWORD - SCRYPT:>", encryptPass)

	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var user []UserLoginInfo

	collection := sessionCopy.DB(util.DATABASENAME).C("users")
	opMatch := bson.M{"$match": bson.M{"status": "Active", "emailid": u.EmailID, "password": hex.EncodeToString(encryptPass)}}

	opProject := bson.M{
		"$project": bson.M{"userid": "$_id", "firstname": 1, "lastname": 1, "emailid": 1, "avatarurl": "$pictureurl", "description": 1, "scope": 1, "mobile": "$phone"},
	}

	opMain := []bson.M{opMatch, opProject}
	pipe := collection.Pipe(opMain)
	err = pipe.All(&user)

	if err != nil {
		fmt.Println("Error Getting user details" + err.Error())
	}

	if len(user) == 0 {
		//	//check user status
	userExistResult,userStatus	:= isUserExists(u.EmailID, mongoSession)
	fmt.Println("userExistResult",userExistResult)
	fmt.Println("userStatus",userStatus)
	//Check if user exists
	if userExistResult != false {
		userExistsResult := map[string]interface{}{"code": util.CODE_USR106, "message": "Error", "result": userStatus}
		userResp, err := json.Marshal(userExistsResult)
		if err != nil {
			fmt.Println("Error - The user already exists" + err.Error())
		}
		return string(userResp)
	}
		fmt.Println("EMAIL NOT FOUND")
		loginValResult := map[string]interface{}{"code": util.CODE_USR1001, "message": "Error", "result": "Wrong EmailID or password"}
		loginResp, err := json.Marshal(loginValResult)
		if err != nil {
			fmt.Println("Error login user" + err.Error())
		}
		return string(loginResp)
	} else {
		claims := util.JBNClaims{}
		expireToken := time.Now().Add(time.Hour * 24 * 30).Unix()

		for _, ur := range user {
			uid := hex.EncodeToString([]byte(ur.UserID))
			fmt.Println("USER DETAILS-", ur.EmailID)
			// Create the Joybynature Claims
			claims = util.JBNClaims{
				uid,
				ur.FirstName,
				ur.LastName,
				ur.Description,
				ur.Mobile,
				ur.AvatarURL,
				ur.EmailID,
				ur.Scope, // "admin, can-read, can-write, can-delete",
				jwt.StandardClaims{
					ExpiresAt: expireToken,
					Issuer: "joybynature.com",
				},
			}

		}
		//Create JWT token
		token := util.GenerateToken(claims)
		userLoginResult := map[string]interface{}{"code": util.CODE_USR104, "message": "Success", "result": token}
		resp, err := json.Marshal(userLoginResult)
		if err != nil {
			fmt.Println("Error login user" + err.Error())
		}

		return string(resp)
	}
}

//Validate user login form data
func validateLoginUser(logindata UserLoginForm) (isValid bool, errorMessage string) {
	vErrors := []string{
		"User Email cannot be blank",
		"Enter valid Email",
	}
	emailid := logindata.EmailID
	EmailRegx := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	EmailMatch := EmailRegx.MatchString(emailid)

	//Check user email
	if emailid == "" { //checkIsBlank
		return false, vErrors[0]
	} else if EmailMatch != true {
		return false, vErrors[1]
	}
	//Valid form data
	return true, ""
}

func AddUser(c *iris.Context, mongoSession *mgo.Session) (addUser string) {
	//Get signup form data
	u := SignUpForm{}

	err := c.ReadForm(&u)
	if err != nil {
		fmt.Println("Error reading signup form" + err.Error())
	}

	//Vaidate the signup form data
	isValid, errMsg := validateAddUser(u)

	if !isValid {
		sigupValidResult := map[string]interface{}{"code": util.CODE_USR108, "message": "Error", "result": errMsg}
		signupResp, err := json.Marshal(sigupValidResult)
		if err != nil {
			fmt.Println("Error in signup form data" + err.Error())
		}
		return string(signupResp)
	}

	userExistResult,userStatus	:= isUserExists(u.EmailID, mongoSession)
	fmt.Println("userExistResult",userExistResult)
	fmt.Println("userStatus",userStatus)
	//Check if user exists
	if userExistResult != false {
		userExistsResult := map[string]interface{}{"code": util.CODE_USR106, "message": "Error", "result": userStatus}
		userResp, err := json.Marshal(userExistsResult)
		if err != nil {
			fmt.Println("Error - The user already exists" + err.Error())
		}
		return string(userResp)
	}

	//Create new user
	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(util.DATABASENAME).C("users")

	//Create password
	encryptPass, err := scrypt.Key([]byte(u.Password), util.SALT, 16384, 8, 1, 32)
	if err != nil {
		fmt.Println("Error using scrypt.key" + err.Error())
		return
	}

	//Add User Data
	userData := User{
		UserID:       bson.NewObjectId(),
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		EmailId:      u.EmailID,
		Password:     hex.EncodeToString(encryptPass),
		Phone:        u.Mobile,
		DateOfJoin:   time.Now(),
		Status:       "otp",
		ShrinkStatus:  true,
		Seed:		   false,		
		Scope:        "can-read, can-write, can-delete",
		SubscriberID: "AndroidGCMnumber",
	}
	errUser := collection.Insert(&userData)

	if errUser != nil {
		fmt.Println("ERROR adding new user" + errUser.Error())
		errResponse := map[string]interface{}{"code": util.CODE_USR102, "message": "Error", "result": "Error registering new users"}
		errResp, _ := json.Marshal(errResponse)

		return string(errResp)
	}

	//Added user successfully , respond success json to the client
	userAddResult := map[string]interface{}{"code": util.CODE_ST409, "message": "Success", "result": "User Added Successfully"}
	resp, err := json.Marshal(userAddResult)
	if err != nil {
		fmt.Println("ERROR creating JSON data for adding user" + err.Error())
	}else{
//	phone	:= u.Mobile
	msg := SendOtp(u.Mobile,c,mongoSession)
	if msg == true{
			return string(resp)	
	}else{
		userAddResult := map[string]interface{}{"code": util.CODE_USR101, "message": "Fail", "result": "Otp send issue"}
	resp, err := json.Marshal(userAddResult)
	if err != nil {
		fmt.Println("ERROR sending otp to user" + err.Error())
		}
		return string(resp)	
		}
	}
			return string(resp)	
}

func isUserExists(emailid string, mongoSession *mgo.Session) (isavailable bool,userStatus string) {
	//Get a mongo connection
	fmt.Println("emailid",emailid)
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	 user := User{}

	collection := sessionCopy.DB(util.DATABASENAME).C("users")
	//opMatch := bson.M{"$match": bson.M{"status": "Active", "emailid": emailid}}
	opMatch := bson.M{"$match": bson.M{"emailid": emailid}}

	opProject := bson.M{
		"$project": bson.M{"_id": 0, "emailid": 1,"status":1},
	}

	opMain := []bson.M{opMatch, opProject}
	pipe := collection.Pipe(opMain)
	err := pipe.One(&user)

	if err != nil {
		fmt.Println("Error Getting user details" + err.Error())
		return false,""

	}

		return true,user.Status

	
}

//Validate Add user signup form data
func validateAddUser(adduserdata SignUpForm) (isValid bool, errorMessage string) {
	vErrors := []string{
		"User Email cannot be blank",
		"Enter valid Email",
		"User Phone number cannot be blank",
		"Enter valid Phone Number",
	}
	emailid := adduserdata.EmailID
	EmailRegx := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	EmailMatch := EmailRegx.MatchString(emailid)

	phone := adduserdata.Mobile
	PhoneRegx := regexp.MustCompile(`[0-9]+`)
	PhoneMatch := PhoneRegx.MatchString(phone)

	//Check user email
	if emailid == "" { //checkIsBlank
		return false, vErrors[0]
	} else if EmailMatch != true {
		return false, vErrors[1]
	} else if phone == "" {
		return false, vErrors[2]
	} else if PhoneMatch != true {
		return false, vErrors[3]
	}
	//Valid form data
	return true, ""
}

func GetUserProfile(c *iris.Context, mongoSession *mgo.Session)(getUserProfileUser string){
	authorInfo := getUserJWT(c,"")
	fmt.Println("name",authorInfo.UserID)
	UserID := c.Param("userid")
	uid := bson.ObjectIdHex(UserID)
	if authorInfo.UserID != uid{
	GetUserProfilePublicResult	:= GetUserProfilePublic(uid,c,mongoSession)
	return GetUserProfilePublicResult
	}
	GetUserProfilePrivateResult	:= GetUserProfilePrivate(uid,c,mongoSession)
	return GetUserProfilePrivateResult
}

func GetUserProfilePublic(uid bson.ObjectId,c *iris.Context, mongoSession *mgo.Session) (userProfile string) {
	 userPublic := UserPublicProfile{}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	
	fmt.Println("uid",uid)
	collection := mongoSession.DB(util.DATABASENAME).C("users")
	opMatch := bson.M{"$match":bson.M{"_id":uid}}
	opProject1 := bson.M{"$project":bson.M{"_id":1,"firstname":1,"lastname":1,"pictureurl":1,"description":1,"aboutus":1,"followerscount":1,"myquestions":1,"myanswers":1}}
	//opProject2 := bson.M{"$project":bson.M{"_id":1,"firstname":1,"lastname":1,"pictureurl":1,"description":1,"followerscount":1,"myquestions":1}}
	operations := []bson.M{opMatch,opProject1}
	pipe := collection.Pipe(operations)
	//err := pipe.All(&userPublic)
	err := pipe.One(&userPublic)


	if err != nil{
		fmt.Println("get user public err",err.Error())
	userProfileerr := map[string]interface{}{"code": util.CODE_USR101, "message": "Fail", "result": err.Error()}
	resp, err := json.Marshal(userProfileerr)
	if err != nil {
		fmt.Println("get user public err" + err.Error())
		}
		return string(resp)	
		}
	
	userProfileResult := map[string]interface{}{"code": util.CODE_USR101, "message": "success", "result": userPublic}
	resp, err := json.Marshal(userProfileResult)
	if err != nil {
		fmt.Println("ERROR sending otp to user" + err.Error())
		}
	
	fmt.Println("fname ji",userPublic.FirstName,userPublic.Pictureurl)	
		return string(resp)	
}

func GetUserProfilePrivate(uid bson.ObjectId,c *iris.Context, mongoSession *mgo.Session) (userProfile string) {
	user := UserPrivateProfile{}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := mongoSession.DB(util.DATABASENAME).C("users")
	opMatch := bson.M{"$match":bson.M{"_id":uid}}
	//opProject := bson.M{"$project":bson.M{"_id":1,"firstname":1,"lastname":1,"pictureurl":1,"description":1,"aboutus":1,"followerscount":1}}
	operations := []bson.M{opMatch/*,opProject*/}
	pipe := collection.Pipe(operations)
	err := pipe.One(&user)
	if err != nil{
		fmt.Println("get user private err",err.Error())
	userProfileerr := map[string]interface{}{"code": util.CODE_USR101, "message": "Fail", "result": err.Error()}
	resp, err := json.Marshal(userProfileerr)
	if err != nil {
		fmt.Println("ERROR sending otp to user" + err.Error())
		}
		return string(resp)	
		}
	
	userProfileResult := map[string]interface{}{"code": util.CODE_USR101, "message": "success", "result": user}
	resp, err := json.Marshal(userProfileResult)
	if err != nil {
		fmt.Println("ERROR sending otp to user" + err.Error())
		}
		return string(resp)	
}

func ForgotPasswordUser(c *iris.Context, mongoSession *mgo.Session) (forgotPassword string) {
	authorInfo := getUserJWT(c,"")
	fmt.Println("name",authorInfo.UserID)
	userID := authorInfo.UserID
	FPassword := ForgetPassword{}
	err := c.ReadForm(&FPassword)
	if err != nil{
	fmt.Println("form read forgot password err",err.Error())
	readFormerr := map[string]interface{}{"code": util.CODE_USR101, "message": "fail", "result": err.Error()}
	resp, err := json.Marshal(readFormerr)
	if err != nil {
		fmt.Println("form read forgot password err" + err.Error())
		}
		return string(resp)
	}
	fmt.Println("password",FPassword.Password)
	
	fmt.Println("uid",userID)
	encryptPass, err := scrypt.Key([]byte(FPassword.Password), util.SALT, 16384, 8, 1, 32)
	if err != nil {
		fmt.Println("Error using scrypt.key" + err.Error())
		return
	}

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(util.DATABASENAME).C("users")
	opMatch := bson.M{"$and":[]interface{}{bson.M{"_id": userID},bson.M{"status":"Active"}}}
	opChange := bson.M{"$set": bson.M{"password": hex.EncodeToString(encryptPass)}}

	err = collection.Update(opMatch,opChange)
	
	if err != nil {
		fmt.Println("Error Getting Forgot Password User" + err.Error())
	ForgotPasswordUserErr := map[string]interface{}{"code": util.CODE_USR101, "message": "fail", "result": err.Error()}
	resp, err := json.Marshal(ForgotPasswordUserErr)
	if err != nil {
		fmt.Println("ERROR form read Forgot Password User" + err.Error())
		}
		return string(resp)
	}
	ForgotPasswordUserResult := map[string]interface{}{"code": util.CODE_USR101, "message": "success", "result": "forgot password change successfully"}
	resp, err := json.Marshal(ForgotPasswordUserResult)
	if err != nil {
		fmt.Println("ERROR form read Forgot Password User" + err.Error())
		}
		
			return string(resp)
}
func CheckEmailGetMobileNumber(c *iris.Context,mongoSession *mgo.Session) (MobileNumber string,EmailExist string) {
	Email := EmailForm{}
	err := c.ReadForm(&Email)
	if err != nil{
	fmt.Println("form read CheckEmailGetMobileNumber err",err.Error())
	readFormerr := map[string]interface{}{"code": util.CODE_USR101, "message": "fail", "result": err.Error()}
	resp, err := json.Marshal(readFormerr)
	if err != nil {
		fmt.Println("ERROR form read CheckEmailGetMobileNumber" + err.Error())
		}
		return "",string(resp)
	}
	emailid := Email.EmailID
	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()


	 user := User{}

	collection := sessionCopy.DB(util.DATABASENAME).C("users")
	opMatch := bson.M{"$match": bson.M{"status": "Active", "emailid": emailid}}

	opProject := bson.M{
		"$project": bson.M{"_id": 0, "emailid": 1,"phone":1},
	}

	opMain := []bson.M{opMatch, opProject}
	pipe := collection.Pipe(opMain)
	err = pipe.One(&user)

	if err != nil {
		fmt.Println("Error Getting CheckEmailGetMobileNumber" + err.Error())
	CheckEmailGetMobileNumberErr := map[string]interface{}{"code": util.CODE_USR101, "message": "fail", "result": err.Error()}
	resp, err := json.Marshal(CheckEmailGetMobileNumberErr)
	if err != nil {
		fmt.Println("ERROR form read CheckEmailGetMobileNumber" + err.Error())
		}
		return "",string(resp)
	}
	CheckEmailGetMobileNumberResult := map[string]interface{}{"code": util.CODE_USR101, "message": "success", "result": true}
	resp, err := json.Marshal(CheckEmailGetMobileNumberResult)
	if err != nil {
		fmt.Println("ERROR form read CheckEmailGetMobileNumber" + err.Error())
		}
		msg := SendOtp(user.Phone,c,mongoSession)
		if msg == true{
			return user.Phone,string(resp)
	}
	CheckEmailGetMobileNumberResult = map[string]interface{}{"code": util.CODE_USR101, "message": "success", "result": "otp send problem"}
	resp, err = json.Marshal(CheckEmailGetMobileNumberResult)
	if err != nil {
		fmt.Println("ERROR form read CheckEmailGetMobileNumber" + err.Error())
		}
					return user.Phone,string(resp)
}

func GetAllUsersDataWeb(c *iris.Context, mongoSession *mgo.Session) {
	c.SetHeader("Access-Control-Allow-Origin", "*")
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	var user []User

	collection := sessionCopy.DB(util.DATABASENAME).C("users")
	err := collection.Find(bson.M{}).All(&user)
	if err != nil {
		fmt.Println("List user data err " + err.Error())
	}
	userResult := map[string]interface{}{"code": util.CODE_QA501, "message": "null", "result": user}
	resp, err := json.Marshal(userResult)
	if err != nil {
		fmt.Println("List user REsult " + err.Error())
	}

	//c.JSON(iris.StatusOK, resp)
	c.Write(string(resp))

}
func GenerateOTP(c *iris.Context) (otpVal int) {
	otpVal = rand.Intn(99999)
	return otpVal
}

func AddOTP(p string,c *iris.Context,mongoSession *mgo.Session) (resultAddOTP bool,otp string){
	fmt.Println("Is otp added")
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	// Get a collection to execute the query against.
	collection := sessionCopy.DB(util.DATABASENAME).C("otps")
	mobile := p
	otpVal	:= GenerateOTP(c)
	otpS := strconv.Itoa(otpVal)
	
	otps := OTPS{
		OtpId:       	bson.NewObjectId(),
		OTP:    		otpS,
		Mobile:  	    mobile,
		CreateDateTime: time.Now(),
		       
	}
	//insert otp
	errOtp := collection.Insert(&otps)

	if errOtp != nil {
		fmt.Println("ERROR adding new otp" + errOtp.Error())
		errResponse := map[string]interface{}{"code": util.CODE_USR102, "message": "Error", "result": "Error registering otps"}
		errResp, _ := json.Marshal(errResponse)
		fmt.Println(errResp)
		return false,""
	}
		otpAddResponse := map[string]interface{}{"code": util.CODE_USR102, "message": "Success", "result": "Otp Added successfully"}
		resp, _ := json.Marshal(otpAddResponse)
		fmt.Println(resp)
		return true,otpS
}


func SendOtp(phone string,c *iris.Context,mongoSession *mgo.Session)(SendOtpResult bool){
	fmt.Println("send otp")
	resultAddOTP,otpS	:= AddOTP(phone,c,mongoSession)
	fmt.Println("resultAddOTP",resultAddOTP)
	fmt.Println("otpS",otpS)
	mobile := phone
	if resultAddOTP == true{
		dataSms := url.Values{}
		SmsSetup :=	config.OtpSms()
		msg := otpS+" "+"is your login otp. Please use it to verify your account with Joybynature.com"
   	 	dataSms.Add("From", "JBN-123")
    	dataSms.Add("To", mobile)
    	dataSms.Add("Body",msg)
 		client := &http.Client{}

    r, _ := http.NewRequest("POST","https://"+SmsSetup.ExotelSid+":"+SmsSetup.ExotelToken+"@twilix.exotel.in/v1/Accounts/"+SmsSetup.ExotelSid+"/Sms/send",bytes.NewBufferString(dataSms.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    r.Header.Add("Content-Length", strconv.Itoa(len(dataSms.Encode())))
    resp, err := client.Do(r)
	if err != nil{
		fmt.Println("err")
	}else{
		 fmt.Println(resp.Status)
		if resp.StatusCode == 200{
		otpSendResponse := map[string]interface{}{"code": util.CODE_USR102, "message": "Success", "result": "Otp Send Sms successfully"}
		resp, _ := json.Marshal(otpSendResponse)
		fmt.Println(resp)
			return true
		}else{
		otpSendResponse := map[string]interface{}{"code": util.CODE_USR102, "message": "Error", "result": "Otp Send Sms server Failed"}
		resp, _ := json.Marshal(otpSendResponse)
		fmt.Println(resp)
			return false
			}	
		 }
	}
		otpAddResponse := map[string]interface{}{"code": util.CODE_USR102, "message": "Success", "result": "Otp Added Failed"}
		resp, _ := json.Marshal(otpAddResponse)
		fmt.Println(resp)
		return false
	
}

func ReSendOtp(c *iris.Context,mongoSession *mgo.Session)(SendOtpResult string){
	u := SignUpForm{}
	err := c.ReadForm(&u)
	if err != nil{
		fmt.Println("Error in otp resend",err.Error())
	}
	fmt.Println("send otp")
	resultAddOTP,otpS	:= AddOTP(u.Mobile,c,mongoSession)
	fmt.Println("resultAddOTP",resultAddOTP)
	fmt.Println("otpS",otpS)
	mobile := u.Mobile
	if resultAddOTP == true{
		dataSms := url.Values{}
		SmsSetup :=	config.OtpSms()
		msg := otpS+" "+"is your login otp. Please use it to verify your account with Joybynature.com"
   	 	dataSms.Add("From", "JBN-123")
    	dataSms.Add("To", mobile)
    	dataSms.Add("Body",msg)
 		client := &http.Client{}

    r, _ := http.NewRequest("POST","https://"+SmsSetup.ExotelSid+":"+SmsSetup.ExotelToken+"@twilix.exotel.in/v1/Accounts/"+SmsSetup.ExotelSid+"/Sms/send",bytes.NewBufferString(dataSms.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    r.Header.Add("Content-Length", strconv.Itoa(len(dataSms.Encode())))
    resp, err := client.Do(r)
	if err != nil{
		fmt.Println("err")
	}else{
		 fmt.Println(resp.Status)
		if resp.StatusCode == 200{
		otpSendResponse := map[string]interface{}{"code": util.CODE_USR102, "message": "Success", "result": "Otp Send Sms successfully"}
		resp, _ := json.Marshal(otpSendResponse)
		fmt.Println(resp)
		ResultChangeUserStatus := ChangeUserPhone(u,c,mongoSession)
		if ResultChangeUserStatus == true{
							return string(resp)			
		}

		}else{
		otpSendResponse := map[string]interface{}{"code": util.CODE_USR102, "message": "Error", "result": "Otp Send Sms server Failed"}
		resp, _ := json.Marshal(otpSendResponse)
		fmt.Println(resp)
			return string(resp)
			}	
		 }
	}
		otpAddResponse := map[string]interface{}{"code": util.CODE_USR102, "message": "Success", "result": "Otp Added Failed"}
		resp, _ := json.Marshal(otpAddResponse)
		fmt.Println(resp)
	return string(resp)	
}

func VerifyOTP(c *iris.Context, mongoSession *mgo.Session) (verifyOTP string) {
	Votp :=	VerifyOTPForm{}
	err := c.ReadForm(&Votp)
	if err != nil{
		fmt.Println(" otp form data err " + err.Error())
		}
	fmt.Println("entered otp",Votp.Otp,Votp.EmailID,Votp.Mobile)
	otps := OTPS{}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(util.DATABASENAME).C("otps")

	err = collection.Find(bson.M{"$and":[]interface{}{bson.M{"mobile":Votp.Mobile},bson.M{"otp": Votp.Otp}}}).One(&otps)
	//err = collection.Find(bson.M{"mobile":Votp.Mobile,"otp": Votp.Otp}).One(&otps)
	if err != nil {
		fmt.Println("otp Find err" + err.Error())
		verifyotperr := map[string]interface{}{"code": util.CODE_OTP1101, "message": "Fail", "result": err.Error()}
		resp, err := json.Marshal(verifyotperr)
		if err != nil {
			fmt.Println(" otp err " + err.Error())
		}
			return string(resp)
	} 
			ResultChangeUserStatus := ChangeUserStatus(Votp,c,mongoSession)
			if ResultChangeUserStatus == true{
			ResultOtpVerifiedToken := OtpVerifiedToken(Votp,c,mongoSession)
			return ResultOtpVerifiedToken
			}
		verifyotperr := map[string]interface{}{"code": util.CODE_OTP1101, "message": "Fail", "result": err.Error()}
		resp, err := json.Marshal(verifyotperr)
		if err != nil {
			fmt.Println(" otp err " + err.Error())
		}
			return string(resp)

}

func OtpVerifiedToken(Votp VerifyOTPForm,c *iris.Context,mongoSession *mgo.Session) (ResultOtpVerifiedToken string){
	 	var user []UserLoginInfo
		sessionCopy := mongoSession.Copy()
		defer sessionCopy.Close()
		collection := sessionCopy.DB(util.DATABASENAME).C("users")
		
		opMatch := bson.M{"$match": bson.M{"status": "Active", "emailid": Votp.EmailID}}
		opProject := bson.M{
		"$project": bson.M{"userid": "$_id", "firstname": 1, "lastname": 1, "emailid": 1, "avatarurl": "$pictureurl", "description": 1, "scope": 1, "mobile": "$phone"},
	}

	opMain := []bson.M{opMatch,opProject}
	pipe := collection.Pipe(opMain)
	err := pipe.All(&user)
	
		if err != nil {
			fmt.Println("user iNFORMATION err" + err.Error())
			userErr := map[string]interface{}{"code": util.CODE_OTP1101, "message": "Fail", "result": err.Error()}
			resp, err := json.Marshal(userErr)
			if err != nil {
				fmt.Println(" user iNFORMATION err " + err.Error())
			}
			return string(resp)
			}
			fmt.Println(user)
		claims := util.JBNClaims{}
		expireToken := time.Now().Add(time.Hour * 1).Unix()

		for _, ur := range user {
			uid := hex.EncodeToString([]byte(ur.UserID))
			fmt.Println("USER DETAILS-", ur.EmailID)
			// Create the Joybynature Claims
			claims = util.JBNClaims{
				uid,
				ur.FirstName,
				ur.LastName,
				ur.Description,
				ur.Mobile,
				ur.AvatarURL,
				ur.EmailID,
				ur.Scope, // "admin, can-read, can-write, can-delete",
				jwt.StandardClaims{
					ExpiresAt: expireToken,
					Issuer: "joybynature.com",
				},
			}

		}
		//Create JWT token
		token := util.GenerateToken(claims)
		userLoginResult := map[string]interface{}{"code": util.CODE_USR104, "message": "Success", "result": token}
		resp, err := json.Marshal(userLoginResult)
		if err != nil {
			fmt.Println("Error login user" + err.Error())
		}

		return string(resp)
}

func ChangeUserStatus(Votp VerifyOTPForm,c *iris.Context,mongoSession *mgo.Session) (ResultChangeUserStatus bool){
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(util.DATABASENAME).C("users")
		Match := bson.M{"$and":[]interface{}{bson.M{"emailid": Votp.EmailID},bson.M{"phone":Votp.Mobile}}}
		Change := bson.M{"$set": bson.M{"status": "Active"}}
		err := collection.Update(Match, Change)
		if err != nil {
			fmt.Println("user  status update err" + err.Error())
			otpstatusupdateerr := map[string]interface{}{"code": util.CODE_OTP1101, "message": "Fail", "result": err.Error()}
			resp, err := json.Marshal(otpstatusupdateerr)
			if err != nil {
				fmt.Println(" otp status update err " + err.Error())
			}
			fmt.Println(resp)
			return false
		} 
	
		return true
}
func ChangeUserPhone(u SignUpForm,c *iris.Context,mongoSession *mgo.Session) (ResultChangeUserStatus bool){
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(util.DATABASENAME).C("users")
		Match := bson.M{"emailid": u.EmailID}
		Change := bson.M{"$set": bson.M{"phone": u.Mobile}}
		err := collection.Update(Match, Change)
		if err != nil {
			fmt.Println("user  status update err" + err.Error())
			otpstatusupdateerr := map[string]interface{}{"code": util.CODE_OTP1101, "message": "Fail", "result": err.Error()}
			resp, err := json.Marshal(otpstatusupdateerr)
			if err != nil {
				fmt.Println(" otp status update err " + err.Error())
			}
			fmt.Println(resp)
			return false
		} 
	
		return true
}


func VerifyUser(c *iris.Context,mongoSession *mgo.Session)(VerifyUserResult string){
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	var user []User
	authorInfo := getUserJWT(c,"")
	fmt.Println("name",authorInfo.UserID)
	//userid := bson.ObjectIdHex(authorInfo.UserID)
	fmt.Println(util.DATABASENAME)
	collections := sessionCopy.DB(util.DATABASENAME).C("users")
	err := collections.Find(bson.M{"_id":authorInfo.UserID,"status":"Active"}).All(&user)
	if err != nil{
		fmt.Println("error in verify user",err.Error())
	}
	fmt.Println(user)
	if user != nil{
	UserResult := map[string]interface{}{"code": util.CODE_OTP1101, "message": "success", "result": true}
			resp, err := json.Marshal(UserResult)
			if err != nil {
				fmt.Println(" verify user Result err " + err.Error())
			}
			fmt.Println(resp)
	return string(resp)
	}
	UserResult := map[string]interface{}{"code": util.CODE_OTP1101, "message": "success", "result": false}
			resp, err := json.Marshal(UserResult)
			if err != nil {
				fmt.Println(" verify user Result err " + err.Error())
			}
			fmt.Println(resp)
	return string(resp)
}


func AddRemoveFollowers(c *iris.Context,mongoSession *mgo.Session)(AddRemoveFollowersResult string){
	f := FollowersForm{}
	err := c.ReadForm(&f)
	if err != nil{
		fmt.Println("add remove followers data not found err",err.Error())
	}
	authorInfo := getUserJWT(c, "")
	uid := authorInfo.UserID 
	fmt.Println("uid",uid)
	fmt.Println("fid",f.FollowerId)
	fmt.Println("ftype",f.FollowerType)
	//chek that user follow this user or not
	followerExist := CheckFollowerExist(f,c,mongoSession)
	if f.FollowerType == "Follow"{
		//if follower not follow this user
		if followerExist != true{
			CheckUserNewOld	:= CheckUserFollowersExists(uid,c,mongoSession)
			//chek user have already follower or fist time follow any follower
			//if user have no follower that one is first follower then
			if CheckUserNewOld != true{
				//add new user with follower
			AddNewFollowerResult	:= AddNewFollower(f,c,mongoSession)
			//increment follower count in user
			IncrementFollowercountResult	:= IncrementFollowercount(f,c,mongoSession)
			fmt.Println("followcount increment",IncrementFollowercountResult)
				return AddNewFollowerResult	

			}else{
				//update user and add the follower
			UpdateNewFollowerResult	:= UpdateNewFollower(f,c,mongoSession)
						//increment follower count in user
			IncrementFollowercountResult	:= IncrementFollowercount(f,c,mongoSession)
			fmt.Println("followcount increment",IncrementFollowercountResult)
			return UpdateNewFollowerResult

			}	

		}
		fmt.Println("U already follow this User")
		addfollowererr := map[string]interface{}{"code": util.CODE_USR107, "message": "Success", "result": "U already follow this User"}
			resp, err := json.Marshal(addfollowererr)
			if err != nil {
				fmt.Println(" addremove follower err " + err.Error())
			}
			return string(resp)
		
	
				}else{
					if followerExist != false{
						//unfollow
					RemoveFollowersResult := RemoveFollowers(f,c,mongoSession)
					fmt.Println("U r UNfollow ")
					DecrementFollowercountResult :=	DecrementFollowercount(f,c,mongoSession)
					fmt.Println("Decrement Follower",DecrementFollowercountResult)
					return RemoveFollowersResult	
					}
					fmt.Println("U r not follow this user")
					addfollowererr := map[string]interface{}{"code": util.CODE_USR107, "message": "Success", "result": "U r not follow this user"}
			resp, err := json.Marshal(addfollowererr)
			if err != nil {
				fmt.Println(" addremove follower err " + err.Error())
			}
			return string(resp)
					
				}
	return ""
}

func GetUserById(uid bson.ObjectId,c *iris.Context, mongoSession *mgo.Session) (name string,image string,subscriberid string) {
	 user := User{}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	
	fmt.Println("uid",uid)
	collection := mongoSession.DB(util.DATABASENAME).C("users")
	opMatch := bson.M{"$match":bson.M{"_id":uid}}
	opProject1 := bson.M{"$project":bson.M{"_id":1,"firstname":1,"pictureurl":1,"subscriberid":1}}
	//opProject2 := bson.M{"$project":bson.M{"_id":1,"firstname":1,"lastname":1,"pictureurl":1,"description":1,"followerscount":1,"myquestions":1}}
	operations := []bson.M{opMatch,opProject1}
	pipe := collection.Pipe(operations)
	//err := pipe.All(&userPublic)
	err := pipe.One(&user)


	if err != nil{
		fmt.Println("get user  err",err.Error())
	userProfileerr := map[string]interface{}{"code": util.CODE_USR101, "message": "Fail", "result": err.Error()}
	resp, err := json.Marshal(userProfileerr)
	if err != nil {
		fmt.Println("get user  err" + err.Error())
		}
		return string(resp),"",""	
		}
	
	
	fmt.Println("fname ji",user.FirstName)	
		return user.FirstName,user.Pictureurl,user.SubscriberID
	
}
func AddNewFollower(f FollowersForm,c *iris.Context,mongoSession *mgo.Session)(AddNewFollowerResult string){
	mongoCopy := mongoSession.Copy()
	defer mongoCopy.Clone()
//	SubscriberId := GetSubscriberID(f,c,mongoSession)
//	fmt.Println("sid",SubscriberId)
	authorInfo := getUserJWT(c, "")
	
	uid := authorInfo.UserID
	UserSubscriberId := GetSubscriberID(uid,c,mongoSession)
	name := authorInfo.Name
	image := authorInfo.ImageURL
	fname,fimage,SubscriberId	:= GetUserById(bson.ObjectIdHex(f.FollowerId),c,mongoSession)
	 fmt.Println(fname,fimage)
	followersubscriber := []FollowerSubscriberID{FollowerSubscriberID{
				FollowerID: bson.ObjectIdHex(f.FollowerId),
				Name:				fname,
				ImageURL:			fimage,
				SubscriberID:		SubscriberId,
				}}
	follower := &Followers{
				ID:					 bson.NewObjectId(),
				UserId:				 uid,
				Name:				name,
				ImageURL:			image,
				SubscriberID:		UserSubscriberId,
				FollowerSubscriber:	followersubscriber,
				}
	collection := mongoCopy.DB(util.DATABASENAME).C("followers")
	
	err := collection.Insert(follower)
	if err != nil{
			fmt.Println("follower insert err"+err.Error())
			addfollowererr := map[string]interface{}{"code": util.CODE_USR107, "message": "Fail", "result": err.Error()}
			resp, err := json.Marshal(addfollowererr)
			if err != nil {
				fmt.Println(" add follower err " + err.Error())
			}
			return string(resp)
		}
	
		addfollowerresult := map[string]interface{}{"code": util.CODE_USR107, "message": "Success", "result":"Follower added successfully"}
		resp, err := json.Marshal(addfollowerresult)
		if err != nil {
			fmt.Println(" add follower err " + err.Error())
		}
		return string(resp)
							
}

func UpdateNewFollower(f FollowersForm,c *iris.Context,mongoSession *mgo.Session)(UpdateNewFollowerResult string){
	mongoCopy := mongoSession.Copy()
	defer mongoCopy.Clone()
//	SubscriberId := GetSubscriberID(f,c,mongoSession)
//	fmt.Println("sid",SubscriberId)
	collection := mongoCopy.DB(util.DATABASENAME).C("followers")
	fname,fimage,SubscriberId	:= GetUserById(bson.ObjectIdHex(f.FollowerId),c,mongoSession)
	 fmt.Println(fname,fimage)
	followersubscriber := []FollowerSubscriberID{FollowerSubscriberID{
				FollowerID: bson.ObjectIdHex(f.FollowerId),
				SubscriberID:SubscriberId,
				Name:				fname,
				ImageURL:			fimage,
				}}
	authorInfo := getUserJWT(c, "")
	uid := authorInfo.UserID 
	matchQueri := bson.M{"userid": uid}
		change := bson.M{"$set":bson.M{"imageurl":authorInfo.ImageURL},"$push": bson.M{"followersubscriberid": bson.M{"$each": followersubscriber}}}

	//change := bson.M{"$push": bson.M{"followersubscriberid": bson.M{"$each": followersubscriber}}}
	err := collection.Update(matchQueri, change)
	if err != nil{
	    fmt.Println("follower insert err"+err.Error())
		addfollowererr := map[string]interface{}{"code": util.CODE_USR107, "message": "Fail", "result": err.Error()}
		resp, err := json.Marshal(addfollowererr)
		if err != nil {
			fmt.Println(" add follower err " + err.Error())
		}
			return string(resp)
		}
														//increment followercount 
		addfollowerresult := map[string]interface{}{"code": util.CODE_USR107, "message": "Success", "result":"Follower added successfully"}
		resp, err := json.Marshal(addfollowerresult)
		if err != nil {
			fmt.Println(" add follower err " + err.Error())
		}
		return string(resp)
}

func RemoveFollowers(f FollowersForm,c *iris.Context,mongoSession *mgo.Session)(RemoveFollowersResult string){
	mongoCopy := mongoSession.Copy()
	defer mongoCopy.Clone()
//	SubscriberId := GetSubscriberID(f,c,mongoSession)
//	fmt.Println("sid",SubscriberId)
	collection := mongoCopy.DB(util.DATABASENAME).C("followers")

	authorInfo := getUserJWT(c, "")
	uid := authorInfo.UserID 
	matchQueri := bson.M{"userid": uid}
	change := bson.M{"$pull": bson.M{"followersubscriberid":bson.M{"followerid": bson.ObjectIdHex(f.FollowerId)}}}
	err := collection.Update(matchQueri, change)
	if err != nil{
	    fmt.Println("follower remove err"+err.Error())
		addfollowererr := map[string]interface{}{"code": util.CODE_USR107, "message": "Fail", "result": err.Error()}
		resp, err := json.Marshal(addfollowererr)
		if err != nil {
			fmt.Println(" remove follower err " + err.Error())
		}
			return string(resp)
		}
														
		addfollowerresult := map[string]interface{}{"code": util.CODE_USR107, "message": "Success", "result":"Follower remove successfully"}
		resp, err := json.Marshal(addfollowerresult)
		if err != nil {
			fmt.Println(" remove follower err " + err.Error())
		}
		return string(resp)
	
}

func DecrementFollowercount(f FollowersForm,c *iris.Context,mongoSession *mgo.Session)(DecrementFollowercountResult bool){
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	collection := sessionCopy.DB(util.DATABASENAME).C("users")
	opMatchQuery := bson.M{"_id":bson.ObjectIdHex(f.FollowerId) }
	opUpdate := bson.M{"$inc": bson.M{"followerscount": -1}}
	err := collection.Update(opMatchQuery, opUpdate)
	if err != nil{
		return false
	}
	return true
}

func GetSubscriberID(uid bson.ObjectId,c *iris.Context,mongoSession *mgo.Session)(GetSubscriberIDResult string){
	fmt.Println("id",uid)
	 user := User{}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	
	collection := sessionCopy.DB(util.DATABASENAME).C("users")
	err :=	collection.Find(bson.M{"_id":uid,"status":"Active"}).One(&user)
	if err != nil {
		fmt.Println("Get subscriberid err",err.Error())
		SubscriberIderr := map[string]interface{}{"code": util.CODE_USR107, "message": "Fail", "result": err.Error()}
		resp, err := json.Marshal(SubscriberIderr)
		if err != nil {
			fmt.Println(" Get subscriberid err " + err.Error())
		}
		return string(resp)

	}
	fmt.Println("subscriberid",user.SubscriberID)
	fmt.Println("user id",user.UserID)	
	return user.SubscriberID
}

func CheckUserFollowersExists(uid bson.ObjectId,c *iris.Context,mongoSession *mgo.Session)(CheckUserFollowersExistsResult bool){
	 follower := Followers{}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	
	collection := sessionCopy.DB(util.DATABASENAME).C("followers")
	err :=	collection.Find(bson.M{"userid":uid}).One(&follower)
	if err != nil {
		fmt.Println("Check User Followers Exists err",err.Error())
		
		return false

	}
		return true	
}


func IncrementFollowercount(f FollowersForm,c *iris.Context,mongoSession *mgo.Session)(IncrementFollowercountResult bool){
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	collection := sessionCopy.DB(util.DATABASENAME).C("users")
	opMatchQuery := bson.M{"_id":bson.ObjectIdHex(f.FollowerId) }
	opUpdate := bson.M{"$inc": bson.M{"followerscount": 1}}
	err := collection.Update(opMatchQuery, opUpdate)
	if err != nil{
		return false
	}
	return true
}

func CheckFollowerExist(f FollowersForm,c *iris.Context,mongoSession *mgo.Session)(CheckFollowerExistResult bool){
	fmt.Println("fid",f.FollowerId)
	 authorInfo := getUserJWT(c, "")
	uid := authorInfo.UserID 
	follower := Followers{}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	
	collection := sessionCopy.DB(util.DATABASENAME).C("followers")
	err :=	collection.Find(bson.M{"userid":uid,"followersubscriberid.followerid":bson.ObjectIdHex(f.FollowerId)}).One(&follower)
	if err != nil {
		fmt.Println("Get subscriberid err",err.Error())
		return false
	}
	fmt.Println("user",follower.UserId)
	fmt.Println("subscriberid id",follower.FollowerSubscriber)	
	return true
}

func GetFollowerExist(c *iris.Context,mongoSession *mgo.Session)(GetFollowerExist string){
	if bson.IsObjectIdHex(c.Param("uid")) == true{
		fmt.Println("not a valid id")
	}
	FollowerId := bson.ObjectIdHex(c.Param("fid"))
	
	fmt.Println("fid",bson.ObjectIdHex(c.Param("fid")))
	 authorInfo := getUserJWT(c, "")
	uid := authorInfo.UserID 
	follower := Followers{}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	
	collection := sessionCopy.DB(util.DATABASENAME).C("followers")
	err :=	collection.Find(bson.M{"userid":uid,"followersubscriberid.followerid":FollowerId}).One(&follower)
	if err != nil {
		fmt.Println(" Get GetWhoIFollowers result err" + err.Error())
		GetWhoIFollowersResult := map[string]interface{}{"code": util.CODE_USR107, "message": "Fail", "result":false}
		resp, err := json.Marshal(GetWhoIFollowersResult)
		if err != nil {
			fmt.Println(" Get GetWhoIFollowers result err" + err.Error())
		}
		return string(resp)
	}
	fmt.Println("user",follower.UserId)
	fmt.Println("subscriberid id",follower.FollowerSubscriber)	
	GetWhoIFollowersResult := map[string]interface{}{"code": util.CODE_USR107, "message": "Success", "result":true}
		resp, err := json.Marshal(GetWhoIFollowersResult)
		if err != nil {
			fmt.Println(" Get GetWhoIFollowers result err" + err.Error())
		}
		return string(resp)
}

func GetWhoIFollowers(c *iris.Context,mongoSession *mgo.Session)(GetWhoIFollowers string){
	fmt.Println("GetWhoIFollowers")
	follower :=  Followers{}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	
	collection := sessionCopy.DB(util.DATABASENAME).C("followers")
	opMatch := bson.M{}
	if bson.IsObjectIdHex(c.Param("uid")) == true{
		uid := bson.ObjectIdHex(c.Param("uid"))
		fmt.Println("uid ",uid)
		opMatch = bson.M{"$match":bson.M{"userid":uid}}
	}else{
		authorInfo := getUserJWT(c, "")
		uid := authorInfo.UserID 
		fmt.Println("uid auth",uid)
		opMatch = bson.M{"$match":bson.M{"userid":uid}}
	}
	opProject := bson.M{"$project":bson.M{"followersubscriberid":1}}
	operations := []bson.M{opMatch,opProject}
	pipe := collection.Pipe(operations)
	err := pipe.One(&follower)
	fmt.Println("pipe",pipe)
	if err != nil {
		fmt.Println("Get GetWhoIFollowers err",err.Error())
		GetWhoIFollowersErr := map[string]interface{}{"code": util.CODE_USR107, "message": "Fail", "result": err.Error()}
		resp, err := json.Marshal(GetWhoIFollowersErr)
		if err != nil {
			fmt.Println(" Get GetWhoIFollowers err " + err.Error())
		}
		return string(resp)
	}
	GetWhoIFollowersResult := map[string]interface{}{"code": util.CODE_USR107, "message": "Success", "result":follower.FollowerSubscriber}
		resp, err := json.Marshal(GetWhoIFollowersResult)
		if err != nil {
			fmt.Println(" Get GetWhoIFollowers result err" + err.Error())
		}
		return string(resp)
	
}

func GetWhoFollowersMe(c *iris.Context,mongoSession *mgo.Session)(GetWhoFollowersMe string){
	fmt.Println("Get Who Followers Me")
	var follower  []Followers
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	
	collection := sessionCopy.DB(util.DATABASENAME).C("followers")
	opMatch := bson.M{}
	if bson.IsObjectIdHex(c.Param("uid")) == true{
		uid := bson.ObjectIdHex(c.Param("uid"))
		fmt.Println("uid ",uid)
		opMatch = bson.M{"$match":bson.M{"followersubscriberid.followerid":uid}}
	}else{
		authorInfo := getUserJWT(c, "")
		uid := authorInfo.UserID 
		fmt.Println("uid auth",uid)
		opMatch = bson.M{"$match":bson.M{"followersubscriberid.followerid":uid}}
	}
	//opMatch := bson.M{"$match":bson.M{"userid":uid}}
	opProject := bson.M{"$project":bson.M{"userid":1,"name":1,"imageurl":1}}
	operations := []bson.M{opMatch,opProject}
	pipe := collection.Pipe(operations)
	err := pipe.All(&follower)
	fmt.Println("pipe",pipe)
	if err != nil {
		fmt.Println("Get Get Who Followers Me err",err.Error())
		GetWhoFollowersMeErr := map[string]interface{}{"code": util.CODE_USR107, "message": "Fail", "result": err.Error()}
		resp, err := json.Marshal(GetWhoFollowersMeErr)
		if err != nil {
			fmt.Println(" Get GetWhoFollowersMe err " + err.Error())
		}
		return string(resp)
	}
	GetWhoFollowersMeResult := map[string]interface{}{"code": util.CODE_USR107, "message": "Success", "result":follower}
		resp, err := json.Marshal(GetWhoFollowersMeResult)
		if err != nil {
			fmt.Println(" Get GetWhoFollowersMe result err" + err.Error())
		}
		return string(resp)
	
}

func UserQuestionCount(c *iris.Context,mongoSession *mgo.Session)(UserQuestionCountResult bool){
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	authorInfo := getUserJWT(c, "")
	collection := sessionCopy.DB(util.DATABASENAME).C("users")
	opMatchQuery := bson.M{"_id":authorInfo.UserID }
	opUpdate := bson.M{"$inc": bson.M{"myquestions": 1}}
	err := collection.Update(opMatchQuery, opUpdate)
	if err != nil{
		return false
	}
	return true
}

func UserAnswerCount(c *iris.Context,mongoSession *mgo.Session)(UserAnswerCountResult bool){
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	authorInfo := getUserJWT(c, "")
	collection := sessionCopy.DB(util.DATABASENAME).C("users")
	opMatchQuery := bson.M{"_id":authorInfo.UserID }
	opUpdate := bson.M{"$inc": bson.M{"myanswers": 1}}
	err := collection.Update(opMatchQuery, opUpdate)
	if err != nil{
		return false
	}
	return true
}


func EditProfile(c *iris.Context,mongoSession *mgo.Session)(EditProfileResult string){
	eprofile := EditProfileForm{}
	err :=	c.ReadForm(&eprofile)
	if err != nil{
		fmt.Println("edit user profile",err.Error())
	}
	 authorInfo := getUserJWT(c, "")
	uid := authorInfo.UserID 
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	imageType := "user"
	imagename := UploadImage(c,imageType)
	fmt.Println("imagename",imagename)
	
	collection := sessionCopy.DB(util.DATABASENAME).C("users")
		Match := bson.M{"_id": uid}
		Change := bson.M{}
	if imagename != ""{
		Change = bson.M{"$set": bson.M{"pictureurl":imagename,"description":eprofile.Description,"aboutus":eprofile.AboutUS,
		"gender":     eprofile.Gender,
		"dob":      eprofile.Dob,
		"shrinkstatus":  false,
		"username":   eprofile.UserName,
		}}
	}else{
		Change = bson.M{"$set": bson.M{"description":eprofile.Description,"aboutus":eprofile.AboutUS,
		"gender":     eprofile.Gender,
		"dob":      eprofile.Dob,
		"username":   eprofile.UserName,
		}}
	}
		err = collection.Update(Match, Change)
		if err != nil {
			fmt.Println("edit profile err" + err.Error())
			editProfileErr := map[string]interface{}{"code": util.CODE_OTP1101, "message": "Fail", "result": err.Error()}
			resp, err := json.Marshal(editProfileErr)
			if err != nil {
				fmt.Println(" edit profile err " + err.Error())
			}
			fmt.Println(resp)
			return string(resp)
		} 
		token := GenerateNewToken(uid,c,mongoSession)

			fmt.Println(token)
			return string(token)
	
}


func FindSubscriberID(c *iris.Context,mongoSession *mgo.Session)(SubscriberId string){
	//find SubscriberId by User
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	userid := bson.ObjectIdHex(c.Param("uid"))
	fmt.Println("userid",userid)
	var user []User
	authorInfo := getUserJWT(c,"")
	fmt.Println("name",authorInfo.UserID)
	//userid := bson.ObjectIdHex(authorInfo.UserID)
	fmt.Println(util.DATABASENAME)
	collections := sessionCopy.DB(util.DATABASENAME).C("users")
	opMatch := bson.M{"_id":userid}
	opProject := bson.M{"$project":bson.M{"subscriberid":1}}
	operations := []bson.M{opMatch,opProject}
	pipe := collections.Pipe(operations)
	err := pipe.All(&user)
	
	if err != nil{
		fmt.Println("error Find SubscriberID",err.Error())
	}
	fmt.Println(user)

	
	UserResult := map[string]interface{}{"code": util.CODE_OTP1101, "message": "success", "result": user}
			resp, err := json.Marshal(UserResult)
			if err != nil {
				fmt.Println("error Find SubscriberID " + err.Error())
			}
			fmt.Println(resp)
	return string(resp)
}

//func SendNotification(message string,subscriberid string){
//	//send notification 
//	     serverKey := "YOUR-KEY"

//    data := map[string]string{
//        "msg": "Hello World1",
//        "sum": "Happy Day",
//    }

//  ids := []string{
//      "token1",
//  }


//  xds := []string{
//      "token5",
//      "token6",
//      "token7",
//  }

//    c := fcm.NewFcmClient(serverKey)
//    c.NewFcmRegIdsMsg(ids, data)
//    c.AppendDevices(xds)

//    status, err := c.Send()


//    if err == nil {
//    status.PrintResults()
//    } else {
//        fmt.Println(err)
//    }

//}

//type LocationForm struct {
//	AddressID		string
//	Address1    	string 		  
//	Address2    	string		 
//	City        	string		 
//	State       	string 		  
//	Zip         	string 		
//	Country	    	string 		
//	CountryCode		string 		
//	DefaultAddress 	bool 	  
//	Province 		string 	  	  
//	ProvinceCode 	string 	  
//}

func AddLocation(c *iris.Context,mongoSession *mgo.Session)(AddLocationResult string){
	location := LocationForm{}
	err := c.ReadForm(&location)
	if err != nil{
		fmt.Println("location form read",err.Error())
	}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	collections := sessionCopy.DB(util.DATABASENAME).C("users")
	  sessionCopy.SetMode(mgo.Monotonic, true)

	//db.somecollection.update({name: "some name"},{ $set: {"lastseen": "2012-12-28"}}, {upsert:true})
	fmt.Println("location.Address1",location.Address1)
	fmt.Println("location.Address1",location.Address2)
	fmt.Println("location.Address1",location.City)
	fmt.Println("location.Address1",location.Country)
	fmt.Println("location.Address1",location.CountryCode)
	fmt.Println("location.Address1",location.DefaultAddress)
//	Match := bson.M{"_id": uid}
//		Change := bson.M{"$set": bson.M{"pictureurl":imagename,"description":eprofile.Description,"aboutus":eprofile.AboutUS,
//		"gender":     eprofile.Gender,
//		"dob":      eprofile.Dob,
//		"username":   eprofile.UserName,
//		}}
//		err = collection.Update(Match, Change)
	authorInfo := getUserJWT(c,"")
	fmt.Println("name",authorInfo.UserID)
	//opMatchQuery1 = bson.M{"$match":bson.M{"_id":authorInfo.UserID}}
	opMatch := bson.M{}
	if bson.IsObjectIdHex(location.AddressID) == true{
		lid := bson.ObjectIdHex(location.AddressID)
		fmt.Println("lid ",lid)
	opMatch = bson.M{"$and":[]interface{}{bson.M{"_id":authorInfo.UserID},bson.M{"location.addressid": lid}}}	
	
	}else{
		
	opMatch = bson.M{"$match":bson.M{"_id":authorInfo.UserID}}	
	fmt.Println("opmatch",opMatch)
	}
		Change := bson.M{"$set": bson.M{
		"addressid":bson.NewObjectId(),	
		"address1":	location.Address1,
		"address2":	location.Address2,
		"city":		location.City,
		"state":     location.State,
		"zip":      location.Zip,
		"country":   location.Country,
		"countrycode": location.CountryCode,
		"defaultaddress": location.DefaultAddress,
		}}
	//upsertall :=	bson.M{"upsert":true}
	err = collections.Update(opMatch,Change)
	if err != nil{
		fmt.Println("AddEditLocation",err.Error())
		UserResult := map[string]interface{}{"code": util.CODE_USR102, "message": "success", "result": err.Error()}
			resp, err := json.Marshal(UserResult)
			if err != nil {
				fmt.Println("error AddEditLocation " + err.Error())
			}
			fmt.Println(resp)
			return string(resp)
	}
	UserResult := map[string]interface{}{"code": util.CODE_USR102, "message": "success", "result": "add location successfully"}
			resp, err := json.Marshal(UserResult)
			if err != nil {
				fmt.Println("error AddEditLocation " + err.Error())
			}
			fmt.Println(resp)
	return string(resp)
}

func GenerateNewToken(uid bson.ObjectId,c *iris.Context,mongoSession *mgo.Session)(token string){
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var user []User

	collection := mongoSession.DB(util.DATABASENAME).C("users")
	opMatch := bson.M{"$match":bson.M{"_id":uid}}
	opProject1 := bson.M{"$project":bson.M{"_id":1,"firstname":1,"lastname":1,"pictureurl":1,"description":1,"followerscount":1,"myquestions":1,"myanswers":1}}
	//opProject2 := bson.M{"$project":bson.M{"_id":1,"firstname":1,"lastname":1,"pictureurl":1,"description":1,"followerscount":1,"myquestions":1}}
	operations := []bson.M{opMatch,opProject1}
	pipe := collection.Pipe(operations)
	//err := pipe.All(&userPublic)
	err := pipe.All(&user)
	if err != nil {
		fmt.Println("Error Getting user details" + err.Error())
	}

	if len(user) == 0 {
		fmt.Println("EMAIL NOT FOUND")
		loginValResult := map[string]interface{}{"code": util.CODE_USR1001, "message": "Error", "result": "Wrong EmailID or password"}
		loginResp, err := json.Marshal(loginValResult)
		if err != nil {
			fmt.Println("Error login user" + err.Error())
		}
		return string(loginResp)
	} else {
		claims := util.JBNClaims{}
		expireToken := time.Now().Add(time.Hour * 24 * 30).Unix()

		for _, ur := range user {
			uid := hex.EncodeToString([]byte(ur.UserID))
			fmt.Println("USER DETAILS-", ur.EmailId)
			// Create the Joybynature Claims
			claims = util.JBNClaims{
				uid,
				ur.FirstName,
				ur.LastName,
				ur.Description,
				ur.Phone,
				ur.Pictureurl,
				ur.EmailId,
				ur.Scope, // "admin, can-read, can-write, can-delete",
				jwt.StandardClaims{
					ExpiresAt: expireToken,
					Issuer: "joybynature.com",
				},
			}

		}
		//Create JWT token
		token := util.GenerateToken(claims)
		userLoginResult := map[string]interface{}{"code": util.CODE_USR104, "message": "Success", "result": token}
		resp, err := json.Marshal(userLoginResult)
		if err != nil {
			fmt.Println("Error login user" + err.Error())
		}

		return string(resp)
	}
}

type SubscriberID struct{
		SubscriberId		string		`json:"subscriberid" bson:"subscriberid"`	
}

func GetSubscriberListMYFollowers(c *iris.Context,mongoSession *mgo.Session)(GetSubscriberListMYFollowers string){
	fmt.Println("Get Who Followers Me")
	var subscriberid  []SubscriberID
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	
	collection := sessionCopy.DB(util.DATABASENAME).C("followers")
	opMatch := bson.M{}
	if bson.IsObjectIdHex(c.Param("uid")) == true{
		uid := bson.ObjectIdHex(c.Param("uid"))
		fmt.Println("uid ",uid)
		opMatch = bson.M{"$match":bson.M{"followersubscriberid.followerid":uid}}
	}else{
		authorInfo := getUserJWT(c, "")
		uid := authorInfo.UserID 
		fmt.Println("uid auth",uid)
		opMatch = bson.M{"$match":bson.M{"followersubscriberid.followerid":uid}}
	}
	//opMatch := bson.M{"$match":bson.M{"userid":uid}}
	opProject := bson.M{"$project":bson.M{"subscriberid":1}}
	operations := []bson.M{opMatch,opProject}
	pipe := collection.Pipe(operations)
	err := pipe.All(&subscriberid)
	fmt.Println("pipe",pipe)
	if err != nil {
		fmt.Println("Get Get Who Followers Me err",err.Error())
		GetWhoFollowersMeErr := map[string]interface{}{"code": util.CODE_USR107, "message": "Fail", "result": err.Error()}
		resp, err := json.Marshal(GetWhoFollowersMeErr)
		if err != nil {
			fmt.Println(" Get GetWhoFollowersMe err " + err.Error())
		}
		return string(resp)
	}
	fmt.Println("myfollowerlist",subscriberid)
		fmt.Println("myfollowerlist len",len(subscriberid))
	 //GetUserSubscriberList(myfollowerlist,c,mongoSession)
	

	GetWhoFollowersMeResult := map[string]interface{}{"code": util.CODE_USR107, "message": "Success", "result":subscriberid}
		resp, err := json.Marshal(GetWhoFollowersMeResult)
		if err != nil {
			fmt.Println(" Get GetWhoFollowersMe result err" + err.Error())
		}
		
			return string(resp)

}

 type ListSubID struct{
	//ListSubscriptionID[]		string
	SubscriberID 	[]string       `json:"subscriberid" bson:"subscriberid"`

}

//func GetAllSubscriberList(c *iris.Context,mongoSession *mgo.Session)(GetAllSubscriberList string){
//	fmt.Println("Get Who Followers Me")
//	 subscriberid  := ListSubID{}
//	sessionCopy := mongoSession.Copy()
//	defer sessionCopy.Clone()
	
//	collection := sessionCopy.DB(util.DATABASENAME).C("notificationsubscribers")
//	err := collection.Find(bson.M{"_id":bson.ObjectIdHex("584fee2daa1d91161e0541fc")}).One(&subscriberid)

//	if err != nil {
//		fmt.Println("Get Get Who Followers Me err",err.Error())
//		GetWhoFollowersMeErr := map[string]interface{}{"code": util.CODE_USR107, "message": "Fail", "result": err.Error()}
//		resp, err := json.Marshal(GetWhoFollowersMeErr)
//		if err != nil {
//			fmt.Println(" Get GetWhoFollowersMe err " + err.Error())
//		}
//		return string(resp)
//	}
//	fmt.Println("myfollowerlist",subscriberid.SubscriberID)
//	allsubscriber := subscriberid.SubscriberID
////	ShopNotificationAll(allsubscriber,c)
	

//	GetWhoFollowersMeResult := map[string]interface{}{"code": util.CODE_USR107, "message": "Success", "result":subscriberid}
//		resp, err := json.Marshal(GetWhoFollowersMeResult)
//		if err != nil {
//			fmt.Println(" Get GetWhoFollowersMe result err" + err.Error())
//		}

//			return string(resp)

//}
type NotificationForm struct{
	Title string
	Body  string
	Icon  string
}
func GetAllSubscriberList(c *iris.Context,mongoSession *mgo.Session)(GetAllSubscriberListstring string,GetAllSubscriberList []string){
	 subscriberid  := ListSubID{}
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Clone()
	
	collection := sessionCopy.DB(util.DATABASENAME).C("notificationsubscribers")
	err := collection.Find(bson.M{"_id":bson.ObjectIdHex("584fee2daa1d91161e0541fc")}).One(&subscriberid)

	if err != nil {
		fmt.Println("Get Get Who Followers Me err",err.Error())
		GetWhoFollowersMeErr := map[string]interface{}{"code": util.CODE_USR107, "message": "Fail", "result": err.Error()}
		resp, err := json.Marshal(GetWhoFollowersMeErr)
		if err != nil {
			fmt.Println(" Get GetWhoFollowersMe err " + err.Error())
		}
		return string(resp),subscriberid.SubscriberID
	}
	fmt.Println("myfollowerlist",subscriberid.SubscriberID)
	allsubscriber := subscriberid.SubscriberID

GetWhoFollowersMeResult := map[string]interface{}{"code": util.CODE_USR107, "message": "Success", "result":subscriberid}
		resp, err := json.Marshal(GetWhoFollowersMeResult)
		if err != nil {
			fmt.Println(" Get GetWhoFollowersMe result err" + err.Error())
		}

			return string(resp),allsubscriber
}
func SendNotificationToAll(c *iris.Context,mongoSession *mgo.Session){
	notificationmsg := NotificationForm{}
	err := c.ReadForm(&notificationmsg)
	if err != nil{
		fmt.Println("notification form err",err)
	}
	fmt.Println("Title",notificationmsg.Title)
	fmt.Println("Body",notificationmsg.Body)
	GetAllSubscriberListstring,GetAllSubscriberList := GetAllSubscriberList(c,mongoSession)
	fmt.Println("GetAllSubscriberListstring",GetAllSubscriberListstring)
	serverKey := "AAAAWQ5h-cc:APA91bFqSLXZvkgiuWKo8PVOfjdpsajd5gv_PKi04dQvEfL5B_0Y9RudCuhD1W-eZURY6L7tAH37fYgEXatnA8mXIuq7P4OjS4lmDjeXHQUswAnGSmOIo0679mpaRoBl_HyhAc6a7kUPAxTGD5uSo2v9uNj8nJ06ag"

    data := map[string]string{
        "msg": "Hello how r u",
        "sum": "Happy Day",
		"image": "https://s3.ap-south-1.amazonaws.com/joybynaturedev/user/IMG-20161118-WA0007.jpg",
    }

    notificationPayload := fcm.NotificationPayload{
        Title: notificationmsg.Title,
        Body: notificationmsg.Body,
		//Icon: "https://s3.ap-south-1.amazonaws.com/joybynaturedev/user/IMG-20161118-WA0007.jpg",
	}

    fcm := fcm.NewFcmClient(serverKey)

    fcm.SetNotificationPayload(&notificationPayload)

    fcm.NewFcmRegIdsMsg(GetAllSubscriberList, data)

    status, err := fcm.Send()

    if err == nil {
        status.PrintResults()
		c.Write("notification send successfully")
    } else {
        fmt.Println(err)
    }
}

type AllSubscriberForm struct{
	SubscriberID	[]string
}
type OneSubscriberForm struct{
	SubscriberID	string
}

func InsertAllSubscriberList(c *iris.Context,mongoSession *mgo.Session)(InsertAllSubscriberListResult string){
	subId :=	AllSubscriberForm{}
	err := c.ReadForm(&subId)
	if err != nil{
		fmt.Println("err in insert all sublist",err)
	}
	mongoCopy := mongoSession.Copy()
	defer mongoCopy.Clone()
	collection := mongoCopy.DB(util.DATABASENAME).C("notificationsubscribers")
	
	matchQueri := bson.M{"_id": bson.ObjectIdHex("584fee2daa1d91161e0541fc")}
	change := bson.M{"$push": bson.M{"subscriberid": bson.M{"$each": subId.SubscriberID}}}
	err = collection.Update(matchQueri, change)
	if err != nil{
	    fmt.Println("InsertAllSubscriberList insert err"+err.Error())
		AllSubscriberListerr := map[string]interface{}{"code": util.CODE_USR107, "message": "Fail", "result": err.Error()}
		resp, err := json.Marshal(AllSubscriberListerr)
		if err != nil {
			fmt.Println("  Insert All Subscriber List err " + err.Error())
		}
			return string(resp)
		}
		AllSubscriberListresult := map[string]interface{}{"code": util.CODE_USR107, "message": "Success", "result":"SubscriberId added successfully"}
		resp, err := json.Marshal(AllSubscriberListresult)
		if err != nil {
			fmt.Println(" add AllSubscriberListresult err " + err.Error())
		}
		return string(resp)
}

func UpdateUserSubscriberId(c *iris.Context,mongoSession *mgo.Session)(UpdateSubscriberIdResult string){
			subId :=	OneSubscriberForm{}
				err := c.ReadForm(&subId)
				if err != nil{
					fmt.Println("err in insert one sublist",err)
				}
			authorInfo := getUserJWT(c,"")
			fmt.Println("name",authorInfo.UserID)
			UserID := authorInfo.UserID

			sessionCopy := mongoSession.Copy()
			defer sessionCopy.Close()

	collection := sessionCopy.DB(util.DATABASENAME).C("users")
	opMatch := bson.M{"$and":[]interface{}{bson.M{"_id": UserID},bson.M{"status":"Active"}}}
	opChange := bson.M{"$set": bson.M{"subscriberid": subId.SubscriberID}}

	err = collection.Update(opMatch,opChange)
	
	if err != nil {
		fmt.Println("Error UpdateUserSubscriberId" + err.Error())
	UpdateUserSubscriberIdErr := map[string]interface{}{"code": util.CODE_USR101, "message": "fail", "result": err.Error()}
	resp, err := json.Marshal(UpdateUserSubscriberIdErr)
	if err != nil {
		fmt.Println("ERROR UpdateUserSubscriberId" + err.Error())
		}
		return string(resp)
	}

UpdateUserSubscriberIdResult := map[string]interface{}{"code": util.CODE_USR101, "message": "success", "result": "Update User SubscriberId Successfully"}
	resp, err := json.Marshal(UpdateUserSubscriberIdResult)
	if err != nil {
		fmt.Println("ERROR UpdateUserSubscriberId" + err.Error())
		}
		return string(resp)

}
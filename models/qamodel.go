package models

/**
 * ============================================
 * Date | Author & Changes
 * =============================================
 * 05-Sept-2016 | Ramesh Kunhiraman
 *  -Add QA tags
 *	-Remove non key words from question
 *	-Get QA tags
 *
 * 27-Sept-2016 | Ramesh Kunhiraman
 * -Add QA Emox
 * -Change struct fields
 * -Add QA comments
 *
 */


import (
    //"github.com/disintegration/imaging"
	"encoding/json"
//	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
	//"net/url"
//	"encoding/base64"
	 "image"
    //"os"
     //"image/jpeg"
    _ "image/png"
//	"net/http"
    //"github.com/nfnt/resize"

//	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
//	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
//	  "github.com/aws/aws-sdk-go/service/s3"
	"github.com/dgrijalva/jwt-go"
	"github.com/joybynature/jbnserverapp/util"
	"github.com/joybynature/jbnserverapp/config"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
 	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
)

var (
	errQAAdd = errors.New("Unable to post a question.")
)

type Answer struct {
	AnswerID     bson.ObjectId  `json:"answerid" bson:"answerid"`
	Text         string         `json:"text" bson:"text"`
	ImageURL     string         `json:"imageurl" bson:"imageurl"`
	VideoURL     string         `json:"videourl" bson:"videourl"`
	Author       Author         `json:"author" bson:"author"`
	PostDateTime time.Time      `json:"postdatetime" bson:"postdatetime"`
	EmoxScale    EmoactionScale `json:"emoxscale" bson:"emoxscale"`
	Status       string         `json:"status" bson:"status"`
	CommentCount int            `json:"commentcount" bson:"commentcount"`
	EmoxCount    int            `json:"emoxcount" bson:"emoxcount"`
	PopularCount int            `json:"popularcount" bson:"popularcount"`
}

type EmoactionScale struct {
	Anger   int `json:"anger" bson:"anger"`
	Sad     int `json:"sad" bson:"sad"`
	OK      int `json:"ok" bson:"ok"`
	Happy   int `json:"happy" bson:"happy"`
	Awesome int `json:"awesome" bson:"awesome"`
}

type QAEmox struct {
	QuestionID   bson.ObjectId `json:"questionid" bson:"questionid"`
	AnswerID     bson.ObjectId `json:"answerid" bson:"answerid"`
	EmoxText     string        `json:"emoxtext" bson:"emoxtext"`
	ByUser       bson.ObjectId `json:"userid" bson:"userid"`
	EmoxDateTime time.Time     `json:"emoxdatetime" bson:"emoxdatetime"`
}

//for handling refreshs
type QAEmoxScale struct {
	QuestionID bson.ObjectId  `json:"questionid" bson:"questionid"`
	AnswerID   bson.ObjectId  `json:"answerid" bson:"answerid"`
	EmoxScale  EmoactionScale `json:"emoxscale" bson:"emoxscale"`
}

//type AnswerWithEmox struct{
//	AnswerId	  bson.ObjectId  `json:"answerid" bson:"answerid"`
//	EmoxScale    EmoactionScale `json:"emoxscale" bson:"emoxscale"`
//}
type QAComment struct {
	CommentID    bson.ObjectId `json:"_id" bson:"_id"`
	QuestionID   bson.ObjectId `json:"questionid" bson:"questionid"`
	AnswerID     bson.ObjectId `json:"answerid" bson:"answerid"`
	PostText     string        `json:"posttext" bson:"posttext"`
	Author       Author        `json:"author" bson:"author"`
	PostDateTime time.Time     `json:"postdatetime" bson:"postdatetime"`
}

type QACommentForm struct {
	//CommentID  string
	QuestionID string
	AnswerID   string
	PostText   string
	Author     string
}

type QuestionRevision struct {
	RevisionText     string    `json:"revisiontext" bson:"revisiontext"`
	RevisionDateTime time.Time `json:"revisiondatetime" bson:"revisiondatetime"`
}

type QuestionAnswer struct {
	QAID         bson.ObjectId      `json:"_id" bson:"_id"`
	Title        string             `json:"title" bson:"title"`
	Revision     []QuestionRevision `json:"revision" bson:"revision"`
	Author       Author             `json:"author" bson:"author"`
	PostDateTime time.Time          `json:"postdatetime" bson:"postdatetime"`
	Status       string             `json:"status" bson:"status"`
	QATags       []string           `json:"qatags" bson:"qatags"`
	Category     string             `json:"category" bson:"category"`
	Answer       []Answer           `json:"answer" bson:"answer"`
	AnswerCount  int                `json:"answercount" bson:"answercount"`
	QAViewsCount int                `json:"qaviewscount" bson:"qaviewscount"`

}

type QAFilter struct {
	QAID         bson.ObjectId      `json:"_id" bson:"_id"`
	Title        string             `json:"title" bson:"title"`
	Author       Author             `json:"author" bson:"author"`
	PostDateTime time.Time          `json:"postdatetime" bson:"postdatetime"`
	Status       string             `json:"status" bson:"status"`
	QATags       []string           `json:"qatags" bson:"qatags"`
	Category     string             `json:"category" bson:"category"`
	Answer       []Answer           `json:"answer" bson:"answer"`
	AnswerCount  int                `json:"answercount" bson:"answercount"`
	QAViewsCount int                `json:"qaviewscount" bson:"qaviewscount"`
	Time	    bool                `json:"time" bson:"time"`

}

type RelatedProduct struct {
	SKU         string `json:"sku" bson:"sku"`
	ProductName string `json:"productname" bson:"productname"`
	ImageURL    string `json:"imageurl" bson:"imageurl"`
	Rating      string `json:"rating" bson:"rating"`
}

type RelatedQuestion struct {
	QuestionID   bson.ObjectId `json:"_id" bson:"_id"`
	Tile         string        `json:"title" bson:"title"`
	Author       Author        `json:"author" bson:"author"`
	Postdatetime time.Time     `json:"postdatetime" bson:"postdatetime"`
	QAViewsCount int           `json:"qaviewscount" bson:"qaviewscount"`
	Answercount  int           `json:"answercount" bson:"answercount"`
}

type UnwindQuestionAnswer struct {
	QAID         bson.ObjectId      `json:"_id" bson:"_id"`
	Title        string             `json:"title" bson:"title"`
	Revision     []QuestionRevision `json:"revision" bson:"revision"`
	Author       Author             `json:"author" bson:"author"`
	PostDateTime time.Time          `json:"postdatetime" bson:"postdatetime"`
	Status       string             `json:"status" bson:"status"`
	QATags       []string           `json:"qatags" bson:"qatags"`
	Category     string             `json:"category" bson:"category"`
	Answer       Answer             `json:"answer" bson:"answer"`
	AnswerCount  int                `json:"answercount" bson:"answercount"`
	QAViewsCount int                `json:"qaviewscount" bson:"qaviewscount"`
}

type IgnoreWord struct {
	IId      bson.ObjectId `json:"_id" bson:"_id"`
	Wordname string        `json:"wordname" bson:"wordname"`
}

type QAKeywordTags struct {
	KID      bson.ObjectId `json:"_id" bson:"_id"`
	Keyword  []string      `json:"keyword" bson:"keyword"`
	Category string        `json:"category" bson:"category"`
}

type Question struct {
	Title  string
	Author string //can be anonymous
	Tags   []string
}

type QuestionWebData struct {
	Title             string
	UserID            string
	QuestionAddedDate string
	Name              string
	ImageURL          string
	Followers         int
	SingleLineDesc    string
}
type AnswerWebData struct {
	Qid             string
	Text            string
	AnswerImageURL  string
	VideoURL        string
	AnswerAddedDate string
	UserID          string
	Name            string
	ImageURL        string
	Followers       int
	SingleLineDesc  string
}

type AddAnswerData struct {
	Qid      string
	Text     string
	ImageURL string
	VideoURL string
	Author   string
}

type AddCommentData struct {
	Qid      string
	Aid      string
	PostText string
	Author   string
}

type AddRevisionData struct {
	Qid          string
	RevisionText string
	Author       string
}

type AddEmoxData struct {
	Qid      string
	Aid      string
	EmoxText string
	Author   string
}

type Response struct {
	Text string
}

type IsOwner struct {
	QuestionID string `json:"questionid"`
	AnswerID   string `json:"answerid"`
}


type UserAnswerEmoxCount struct{
	ID					bson.ObjectId	 		`json:"_id" bson:"_id"`
	UserID				bson.ObjectId			`json:"userid" bson:"userid"`
	QuestionID			bson.ObjectId			`json:"questionid" bson:"questionid"`
	AnswerID			bson.ObjectId			`json:"answerid" bson:"answerid"`
	EmoxCount			int					`json:"emoxcount" bson:"emoxcount"`
}


func AddQuestion(c *iris.Context, mongoSession *mgo.Session) (addQuestionResponse string) {
	fmt.Println("inside Add Question")
	loc, _ := time.LoadLocation("Asia/Kolkata")
    fmt.Println("Now: ", loc)
	indiaTime := time.Now().In(loc)
	fmt.Println("time",indiaTime)
	question := Question{}
	err := c.ReadForm(&question)
	fmt.Println("question " + question.Title)
	if err != nil {
		fmt.Println("Error when reading Add question form:" + err.Error())
	}

	//Validate Form data
	isValid, errMsg := validateAddQuestion(question)

	if !isValid {
		qaValidationResult := map[string]interface{}{"code": util.CODE_QA502, "message": "Error", "result": errMsg}
		validationResp, err := json.Marshal(qaValidationResult)
		if err != nil {
			fmt.Println("Add Question Validation Error " + err.Error())
		}
		return string(validationResp)
	} else { //Valida Form data
		//Get a mongo db connection
		sessionCopy := mongoSession.Copy()
		defer sessionCopy.Close()

		// Get a collection to execute the query against.
		collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")

		index := mgo.Index{
			Key:        []string{"title"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}

		err = collection.EnsureIndex(index)
		if err != nil {
			fmt.Println("Error Creating index" + err.Error())
		}
		authorInfo := getUserJWT(c, question.Author)
		fmt.Println("+++++AUTHOR---", authorInfo)

		//Add QA Tags
		qaRelatedTags, err := GenerateQATags(question.Title, c, mongoSession)
		if err != nil {
			fmt.Println("GENERATETAG ERRORS" + err.Error())
		}

		fmt.Println("AUTHOR-" + authorInfo.Name)

		//Add Question
		qDataInfo := &QuestionAnswer{QAID: bson.NewObjectId(), Title: question.Title, PostDateTime: time.Now(), Status: "Active", Author: authorInfo, QATags: qaRelatedTags[0].Keyword, Category: qaRelatedTags[0].Category}
		data, _ := json.MarshalIndent(qDataInfo, "", "    ")

		fmt.Println("DATA -" + string(data))
		err = collection.Insert(qDataInfo)

		if err != nil {
			fmt.Println("Adding question err " + err.Error())
			qaAddQuestionResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Fail", "result": err.Error()}
			resp, err := json.Marshal(qaAddQuestionResult)
			if err != nil {
				fmt.Println("Add Question REsult err " + err.Error())
			}
			return string(resp)
		}

		//Question added successfully , respond success json to the client
		qaAddQuestionResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": "Question Added Successfully"}
		resp, err := json.Marshal(qaAddQuestionResult)
		if err != nil {
			fmt.Println("Add Question REsult " + err.Error())
		}
		//question count increment in user
		UserQuestionCountResult := UserQuestionCount(c,mongoSession)
		fmt.Println(UserQuestionCountResult)
		return string(resp)
	}
}

//Validate Add question form data
func validateAddQuestion(question Question) (isValid bool, errorMessage string) {
	vErrors := []string{
		"Question title cannot be blank",
		"Question title cannot be more than 375 character",
	}
	questiontitlestring := question.Title
	qCharacterCount := strings.Count(questiontitlestring, "")
	fmt.Println(qCharacterCount)
	/*qWordList := strings.Split(question.Title, " ")
	qWordCount := len(qWordList)*/

	//Check question title
	if question.Title == "" { //checkIsBlank
		return false, vErrors[0]
	} else if qCharacterCount >= 375 { //checkWordCount
		return false, vErrors[1]
	}
	//Valid form data
	return true, ""
}

func GenerateQATags(question string, c *iris.Context, mongoSession *mgo.Session) (qatags []QAKeywordTags, err error) {
	fmt.Println("===START: GENERATE QA TAGS===")

	//Remove all the special characters from question
	removeSpecialChrExp := regexp.MustCompile("[^A-Za-z]+")
	questionWordList := removeSpecialChrExp.Split(question, -1)
	fmt.Println("QUESTION - ", questionWordList)

	//Get associated tags referenced by wellbeing related keywords
	QATags, err := getQATags(questionWordList, c, mongoSession)
	if err != nil {
		fmt.Println("GET QA TAGS ERROR", err)
		return QATags, errors.New("GET QA TAGS ERROR - Unable to get key QA tags")
	}

	fmt.Println("===END: GENERATE QA TAGS===", QATags)

	return QATags, nil
}

func getQATags(questionWordList []string, c *iris.Context, mongoSession *mgo.Session) (qaTags []QAKeywordTags, err error) {
	fmt.Println("===START: GET QA TAGS===")

	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	//FIND AND REMOVE NON-KEY WORDS
	// Get a ignore word collection to execute the query against.
	ignoreColl := sessionCopy.DB(util.DATABASENAME).C("ignorewords")

	//Setup to remove non core words
	questionWordsQueryFmt := make([]bson.RegEx, len(questionWordList))
	for iword := range questionWordList {
		questionWordsQueryFmt[iword] = bson.RegEx{"^" + questionWordList[iword] + "$", "i"}
	}
	fmt.Println("QUESTION REGEX FORMAT-", questionWordsQueryFmt)
	//Setup to remove non core words
	var iw []IgnoreWord

	OpMatch := bson.M{"$match": bson.M{"wordname": bson.M{"$in": questionWordsQueryFmt}}}
	//OpUnwind := bson.M{"$unwind": "$wordname"}
	//OpGroup := bson.M{"$group": bson.M{"_id": bson.M{"ignorewords": "$wordname"}, "count": bson.M{"$sum": 1}}}
	OpProject := bson.M{"$project": bson.M{"wordname": 1}}

	operations := []bson.M{OpMatch, OpProject}
	pipe := ignoreColl.Pipe(operations)
	err = pipe.All(&iw)

	if err != nil {
		fmt.Println("Ignore Word query issue " + err.Error())
	}
	fmt.Println("IGNORE WORDS-", iw)

	//determine core tags
	questionWordMap := map[string]bool{}

	//Create a question
	for q := range questionWordList {
		questionWordMap[questionWordList[q]] = true
	}

	for i := range iw {
		delete(questionWordMap, iw[i].Wordname)
	}

	fmt.Println("Core Words -", questionWordMap)

	//GET QA TAGS
	//Get a qa keyword collection to execute the query against.
	qaKeyColl := sessionCopy.DB(util.DATABASENAME).C("qakeywordtags")

	keyIndex := mgo.Index{
		Key:        []string{"keyword"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = qaKeyColl.EnsureIndex(keyIndex)
	if err != nil {
		fmt.Println("CREATE INDEX ERROR-QAKEYWORDS", err)
	}

	//GET associated Category
	var i int
	coreWordsQueryFmt := make([]string, len(questionWordMap))
	for cwKey := range questionWordMap {
		coreWordsQueryFmt[i] = cwKey
		i = i + 1
	}

	var qakeytags []QAKeywordTags

	QAOpMatch := bson.M{"$match": bson.M{"keyword": bson.M{"$all": coreWordsQueryFmt}}}
	QAOpProject := bson.M{"$project": bson.M{"category": 1, "_id": 0}}
	QAOpLimit := bson.M{"$limit": 1}
	QAOperations := []bson.M{QAOpMatch, QAOpProject, QAOpLimit}
	fmt.Println("QA TAG LIST-", QAOperations)
	pipe = qaKeyColl.Pipe(QAOperations)
	err = pipe.All(&qakeytags)

	if qakeytags == nil { //default tag(s)
		qakeytags = []QAKeywordTags{
			QAKeywordTags{
				Keyword:  coreWordsQueryFmt,
				Category: "health",
			},
		}
	}

	//	resp, err := json.Marshal(qakeytags)
	//	if err != nil {
	//		fmt.Println("Add Question Result " + err.Error())
	//	}

	fmt.Println("===END GET KEY TAGS===", qakeytags)

	return qakeytags, nil
}

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

func GetQuestionsList(c *iris.Context, mongoSession *mgo.Session) (qaList string) {
	c.SetHeader("Access-Control-Allow-Origin", "*")

	qaTags := c.Param("tags")
	qaTagsList := strings.Split(qaTags, ",")
	//	pageno, _ := strconv.Atoi(c.Param("pageno"))
	//qid :=
	questionID := c.Param("qid")

	//fmt.Println("PAGE NO-", pageno)

	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var qa []QuestionAnswer

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	tagQueryFmt := make([]bson.RegEx, len(qaTagsList))
	opMatch := bson.M{}
	opMatch1 := bson.M{}

	if qaTags == "" || qaTags == "\"\"" {
		opMatch = bson.M{"$match": bson.M{"status": "Active"}}
	} else {
		for tw := range qaTagsList {
			tagQueryFmt[tw] = bson.RegEx{qaTagsList[tw], "i"}
		}

		opMatch = bson.M{"$match": bson.M{"qatags": bson.M{"$in": tagQueryFmt}, "status": "Active"}}
	}

	if questionID == "" || questionID == "\"\"" {
		opMatch1 = bson.M{"$match": bson.M{"status": "Active"}}
		fmt.Println("if questionid")
	} else {
		lastQuestionId := bson.ObjectIdHex(questionID)
		fmt.Println(lastQuestionId)
		//opMatch1 = bson.M{"$match":bson.M{"$gt":[]interface{}{"$_id",lastQuestionId}}}
		opMatch1 = bson.M{"$match": bson.M{"_id": bson.M{"$lt": lastQuestionId}}}

		fmt.Println("else questionid")
	}

	opProject := bson.M{"$project": bson.M{"_id": 1, "title": 1, "author": 1, "postdatetime": 1, "status": 1, "qatags": 1, "category": 1, "answer": 1}}
	opProject1 := bson.M{"$sort": bson.M{"_id": -1}}
	opProject2 := bson.M{"$limit": 30}

	//	opMain := []bson.M{opMatch, opProject}
	opMain := []bson.M{opMatch, opProject, opProject1, opMatch1, opProject2}

	pipe := collection.Pipe(opMain)
	err := pipe.All(&qa)

	if err != nil {
		fmt.Println("Question with top answer result err " + err.Error())
	}

	qatopanswerResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": qa}
	resp, err := json.Marshal(qatopanswerResult)
	if err != nil {
		fmt.Println("Top QA REsult " + err.Error())
	}

	return string(resp)
}

func GetQADetail(c *iris.Context, mongoSession *mgo.Session) (qaDetail string) {
	questionIdRequestData := bson.ObjectIdHex(c.Param("qid"))

	//Get Mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var qaData []QuestionAnswer

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")

	//Get question details
	opMatch := bson.M{"$match": bson.M{"_id": questionIdRequestData, "status": "Active"}}
	opProject := bson.M{
		"$project": bson.M{"_id": 1, "title": 1, "revision": 1, "author": 1, "postdatetime": 1, "qatags": 1, "answercount": 1, "answer": 1},
	}

	opSort := bson.M{"$sort": bson.M{"postdatetime": -1}}
	opLimit := bson.M{"$limit": 30} //can change later

	opMain := []bson.M{opMatch, opProject, opSort, opLimit}
	pipe := collection.Pipe(opMain)
	err := pipe.All(&qaData)
	if err != nil {
		fmt.Println("Error in finding qa details" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_QA104, "message": "Error", "result": "Error in finding QA Details"}
		errResp, _ := json.Marshal(errRsponse)

		return string(errResp)
	}

	//Increment the QAViews
	opMatchQuery := bson.M{"_id": questionIdRequestData}
	opUpdate := bson.M{"$inc": bson.M{"qaviewscount": 1}}
	err = collection.Update(opMatchQuery, opUpdate)

	if err != nil {
		fmt.Println("Not able to increment the QAViews" + err.Error())
	}

	qaDetailResult := map[string]interface{}{"code": util.CODE_QA103, "message": "Success", "result": qaData}
	resp, err := json.Marshal(qaDetailResult)
	if err != nil {
		fmt.Println("Error in creating JSON format for qa details" + err.Error())
	}

	return string(resp)
}

func GetRelatedQuestions(c *iris.Context, mongoSession *mgo.Session) (relatedQuestions string) {
	//Get qatag
	qaTagsRequestData := strings.Split(c.Param("tags"), ",")

	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var relatedQuestionResponseData []RelatedQuestion

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")

	//Set the regular expression to find related questions
	qaTagsQueryFmt := make([]bson.RegEx, len(qaTagsRequestData))
	for tg := range qaTagsRequestData {
		qaTagsQueryFmt[tg] = bson.RegEx{"^" + qaTagsRequestData[tg] + "$", "i"}
	}

	//Find the related questions based on qatags
	opMatch := bson.M{"$match": bson.M{"qatags": bson.M{"$in": qaTagsQueryFmt}}}
	opProject := bson.M{
		"$project": bson.M{"_id": 1, "title": 1, "author": 1, "postdatetime": 1, "qaviewscount": 1, "answercount": 1},
	}

	opMain := []bson.M{opMatch, opProject}
	pipe := collection.Pipe(opMain)
	err := pipe.All(&relatedQuestionResponseData)
	if err != nil {
		fmt.Println("Error in finding related questions" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_QA106, "message": "Error", "result": "Error in finding the related stories"}
		errResp, _ := json.Marshal(errRsponse)

		return string(errResp)
	}

	qaTagJsonResult := map[string]interface{}{"code": util.CODE_QA105, "message": "Success", "result": relatedQuestionResponseData}
	resp, err := json.Marshal(qaTagJsonResult)
	if err != nil {
		fmt.Println("Error in creating JSON format for RELATED questions" + err.Error())
	}

	return string(resp)
}

func DeleteQuestion(c *iris.Context, mongoSession *mgo.Session) (deleteQuestionResult string) {
	questionID := bson.ObjectIdHex(c.Param("qid"))

	//Get user from Token
	authorInfo := getUserJWT(c, "question")

	//Check if owner
	isOwner := IsQuestionOwner(questionID.String(), authorInfo.UserID.String(), c, mongoSession)

	if !isOwner { //Cannot delete the question if not the owner
		qaDeleteQuestionAccessResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "Result": "Do not have permissions to delete the question"}
		resp, err := json.Marshal(qaDeleteQuestionAccessResult)
		if err != nil {
			fmt.Println("" + err.Error())
		}
		return string(resp)
	}

	//Get mongodb connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	err := collection.Update(bson.M{"_id": questionID}, bson.M{"$set": bson.M{"status": "InActive"}})
	if err != nil {
		fmt.Println("Error deleting the question" + err.Error())
	}

	qaDeleteQuestionResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "Result": "Question Deleted successfully"}
	resp, err := json.Marshal(qaDeleteQuestionResult)
	if err != nil {
		fmt.Println("Error in creating JSON format for Delete Question" + err.Error())
	}

	return string(resp)
}

func SearchQA(c *iris.Context, mongoSession *mgo.Session) (searchQAResult string) {
	searchText := c.Param("text")
	//c.Write(searchText)
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	var qa []QuestionAnswer

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")

	searchbychar := bson.M{"$regex": bson.RegEx{Pattern: searchText, Options: "i"}}
	operation1 := bson.M{"$match": bson.M{"title": searchbychar, "status": "Active"}}
	operation2 := bson.M{"$project": bson.M{"_id": 1, "title": 1}}	
	//operation2 := bson.M{"$project": bson.M{"title": 1}}
//	operation2 := bson.M{"$project": bson.M{"_id": 1, "title": 1, "bgimageurl": 1, "author": 1, "postdatetime": 1, "status": 1, "qatags": 1, "questionfollowers": 1, "qaviews": 1, "answercount": bson.M{"$size": "$answer"}, "oneanswer": bson.M{"$slice": []interface{}{"$answer", -1}}}}
//	operation3 := bson.M{"$project": bson.M{"_id": 1, "title": 1, "bgimageurl": 1, "author": 1, "postdatetime": 1, "status": 1, "qatags": 1, "questionfollowers": 1, "qaviews": 1, "answercount": 1, "answer": "$oneanswer"}}
//	operation4 := bson.M{"$project": bson.M{"_id": 1, "title": 1, "bgimageurl": 1, "author": 1, "postdatetime": 1, "status": 1, "qatags": 1, "questionfollowers": 1, "qaviews": 1, "answercount": 1, "answer": bson.M{"answerid": 1, "text": 1, "author": 1, "imageurl": 1, "videourl": 1, "postdatetime": 1, "emoactions": 1, "emoactionscale": 1, "status": 1, "commentcount": 1}}}
	//operation5 := bson.M{"$sort": bson.M{"postdatetime": -1}}
	//operations := []bson.M{operation1, operation2, operation3, operation4, operation5}

	operations := []bson.M{operation1, operation2}
	pipe := collection.Pipe(operations)
	err := pipe.All(&qa)
	if err != nil {
		fmt.Println("Question search result err " + err.Error())
	}
	searchQaResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": qa}
	resp, err := json.Marshal(searchQaResult)
	if err != nil {
		fmt.Println("Search Qa REsult " + err.Error())
	}

	//c.JSON(iris.StatusOK, resp)
	return string(resp)

}

func AddAnswerToQuestion(c *iris.Context, mongoSession *mgo.Session) (AddAnswerToQuestionResult string) {
	fmt.Println("inside Add Answer")
	//Retrieve all the form data
//	answerResult := Answer{}
	answer := AddAnswerData{}
	err := c.ReadForm(&answer)
	fmt.Println("answer " + answer.Text)
	fmt.Println("QuestionId" +answer.Qid)
	if err != nil {
		fmt.Println("Error when reading Add answer form: " + err.Error())
	}
//	uplod function call bt not no due to comment
	imageType := "qa"
	imagename := UploadImage(c,imageType)
	fmt.Println("imagename",imagename)
	
	
	//Validate Form data
	isValid, errMsg := validateAddAnswer(answer)

	if !isValid {
		qaValidationResult := map[string]interface{}{"code": util.CODE_QA502, "message": "Error", "result": errMsg}
		validationResp, err := json.Marshal(qaValidationResult)
		if err != nil {
			fmt.Println("Add Answer Validation Error " + err.Error())
		}
		return string(validationResp)
	} else { //Valida Form data
		//Get a mongo db connection
		sessionCopy := mongoSession.Copy()
		defer sessionCopy.Close()

		// Get a collection to execute the query against.
		collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")

		//Get user details
		authorInfo := getUserJWT(c, answer.Author)
		//Add answer
		answerinfo := []Answer{Answer{AnswerID: bson.NewObjectId(), Text: answer.Text, ImageURL: imagename, VideoURL: answer.VideoURL, Author: authorInfo, Status: "Active", PostDateTime: time.Now()}}
		qa := &QuestionAnswer{Answer: answerinfo}
		fmt.Println("QuestionId",answer.Qid)
		if bson.IsObjectIdHex(answer.Qid) != true{
			fmt.Println("qid not found")
		}
		qid := bson.ObjectIdHex(answer.Qid)	
		
		
		fmt.Println("qId",qid)
		data, _ := json.MarshalIndent(qa, "", "    ")
		fmt.Println("DATA -" + string(data))
		matchQueri := bson.M{"_id": qid}
		change := bson.M{"$push": bson.M{"answer": bson.M{"$each": answerinfo}}}
		err = collection.Update(matchQueri, change)

		if err != nil {
			fmt.Println("Adding Answer " + err.Error())
			errRsponse := map[string]interface{}{"code": util.CODE_QA102, "message": "Error", "result": "Error adding answer to qa"}
			errResp, _ := json.Marshal(errRsponse)

			return string(errResp)
		}
		collection = sessionCopy.DB(util.DATABASENAME).C("questionanswers")

		opMatchQuery := bson.M{"_id": qid}
		opUpdate := bson.M{"$inc": bson.M{"answercount": 1}}
		err = collection.Update(opMatchQuery, opUpdate)

		if err != nil {
			fmt.Println("Add answercount to the qa" + err.Error())
		}

		//Answer added successfully , respond success json to the client
		qaAddAnswerResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": "Answer Added Successfully"}
		res, err := json.Marshal(qaAddAnswerResult)
		if err != nil {
			fmt.Println("Add Answer REsult " + err.Error())
		}
		UserAnswerCountResult := UserAnswerCount(c,mongoSession)
		fmt.Println(UserAnswerCountResult)
		
				return string(res)

	}
}



func getUserJWT(c *iris.Context, auth string) (user Author) {
	//Get the user token
	jwtUserToken := c.Get("jwt").(*jwt.Token)

	//Get claims data
	jbnClaims := jwtUserToken.Claims.(jwt.MapClaims)

	//if claims, ok := userToken.Claims.(jwt.MapClaims); ok && userToken.Valid {
	//			claimTestedValue = claims["foo"].(string)
	//		} else {
	//			claimTestedValue = "Claims Failed"
	//		}
	var authorInfo = Author{}

	//check if anonymous author
	if auth == "anonymous" { //Set Anonymous users
		authorInfo.UserID = bson.ObjectIdHex(jbnClaims["userid"].(string))
		authorInfo.Name = "Anonymous"
		authorInfo.SingleLineDesc = "iam invisible"
	} else { //set the user
		authorInfo.UserID = bson.ObjectIdHex(jbnClaims["userid"].(string))
		authorInfo.Name = jbnClaims["firstname"].(string)
		authorInfo.ImageURL = jbnClaims["avatarurl"].(string)
		authorInfo.SingleLineDesc = jbnClaims["description"].(string)
	}

	return authorInfo
}

//Validate Add answer form data
func validateAddAnswer(answer AddAnswerData) (isValid bool, errorMessage string) {
	vErrors := []string{
		"Answer text cannot be blank",
		"Answer text cannot be more than 10000 characters",
	}
	answerstring := answer.Text
	aCharacterCount := strings.Count(answerstring, " ")
	/*qWordList := strings.Split(answer.Text, " ")
	qWordCount := len(qWordList)*/

	//Check answer text
	if answer.Text == "" { //checkIsBlank
		return false, vErrors[0]
	} else if aCharacterCount >= 10000 { //checkWordCount
		return false, vErrors[1]
	}
	//Valid form data
	return true, ""
}

func GetAllAnswerComments(c *iris.Context, mongoSession *mgo.Session) (getAllAnswerCommentsResult string) {
	fmt.Println("inside GetAllAnswerComments")
	questionId := c.Param("qid")
	answerId := c.Param("aid")
	fmt.Println(questionId)
	fmt.Println(answerId)
	qId := bson.ObjectIdHex(questionId)
	aId := bson.ObjectIdHex(answerId)
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	var qaComment []QAComment

	collection := sessionCopy.DB(util.DATABASENAME).C("qacomments")
	matchQueri1 := bson.M{"$match":bson.M{"questionid": qId}}
	matchQueri2 := bson.M{"$match":bson.M{"answerid": aId}}
	projectQueri1 := bson.M{"$project":bson.M{"posttext":1,"author":1,"postdatetime":1}}	
	projectQueri2 := bson.M{"$sort":bson.M{"postdatetime":-1}}	
	operations := []bson.M{matchQueri1,matchQueri2,projectQueri1,projectQueri2}
	pipe := collection.Pipe(operations)
	err := pipe.All(&qaComment)
	if err != nil {
		fmt.Println(" comments result err " + err.Error())
	}
	qaAllAnswerCommentsResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": qaComment}
	resp, err := json.Marshal(qaAllAnswerCommentsResult)
	if err != nil {
		fmt.Println(" all comments REsult " + err.Error())
	}
	return string(resp)
}

func GetAnswer(c *iris.Context, mongoSession *mgo.Session) (getAnswerResult string) {
	fmt.Println("inside GetAnswer")
	questionId := c.Param("qid")
	answerId := c.Param("aid")
	fmt.Println(questionId)
	fmt.Println(answerId)
	qId := bson.ObjectIdHex(questionId)
	aId := bson.ObjectIdHex(answerId)
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	var qaAnswer []QuestionAnswer

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	matchQueri1 := bson.M{"$match":bson.M{"_id": qId}}
	projectQueri1 := bson.M{"$project":bson.M{"answer":bson.M{"$filter": bson.M{"input": "$answer", "as": "answer", "cond": bson.M{"$eq": []interface{}{"$$answer.answerid",aId}}}}}}
	projectQueri2 := bson.M{"$project":bson.M{"answer":bson.M{"$filter": bson.M{"input": "$answer", "as": "answer", "cond": bson.M{"$eq": []interface{}{"$$answer.status","Active"}}}}}}
	projectQueri3 := bson.M{"$project":bson.M{"answer.answerid":1,"answer.text":1,"answer.imageurl":1,"answer.videourl":1,"answer.author":1,"answer.postdatetime":1,"answer.emoxscale":1,"answer.commentcount":1,"answer.emoxcount":1}}	
	//projectQueri2 := bson.M{"$sort":bson.M{"postdatetime":-1}}	
	operations := []bson.M{matchQueri1,projectQueri1,projectQueri2,projectQueri3}
	pipe := collection.Pipe(operations)
	err := pipe.All(&qaAnswer)
	if err != nil {
		fmt.Println("Answer  result err " + err.Error())
	}
	qaAllAnswerCommentsResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": qaAnswer}
	resp, err := json.Marshal(qaAllAnswerCommentsResult)
	if err != nil {
		fmt.Println("Answer REsult " + err.Error())
	}
	return string(resp)
}

func GetMyQuestions(c *iris.Context, mongoSession *mgo.Session) (getMyQuestion string) {
	fmt.Println("inside GetMyQuestions")
	authorInfo := getUserJWT(c,"")
	userId := authorInfo.UserID
	fmt.Println(userId)
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	var qa []QuestionAnswer

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")

	opMatch := bson.M{
		"$match": bson.M{
			"status":        "Active",
			"author.userid": userId,
		},
	}

	opProject1 := bson.M{
		"$project": bson.M{"_id": 1, "title": 1, "author": 1, "postdatetime": 1, "qaviews": 1, "answercount": 1,
		},
	}
	opProject2 := bson.M{"$sort": bson.M{"postdatetime": -1}}

	operations := []bson.M{opMatch, opProject1,opProject2}
	pipe := collection.Pipe(operations)
	err := pipe.All(&qa)
	if err != nil {
		fmt.Println("My Question result err " + err.Error())
	}
	qaMyQuestionResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": qa}
	resp, err := json.Marshal(qaMyQuestionResult)
	if err != nil {
		fmt.Println("My Question REsult " + err.Error())
	}

	//c.JSON(iris.StatusOK, resp)
	return string(resp)

}

func GetMyQuestionsCount(userId bson.ObjectId,c *iris.Context, mongoSession *mgo.Session) (GetMyQuestionsCountResult int) {
	fmt.Println("inside GetMyQuestions")
//	uid :=	c.Param("uid")
//	userId := bson.ObjectIdHex(uid)
	fmt.Println(userId)
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
//	var qa []QuestionAnswer

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	count,err := collection.Find(bson.M{"status":"Active","author.userid": userId}).Count()
	if err != nil{
		fmt.Println("err in question count",err.Error())
	}
	fmt.Println("qa count",count)
	
	return count

}

func GetMyAnswers(c *iris.Context, mongoSession *mgo.Session) (getMyAnswerResult string) {
	fmt.Println("inside GetMyAnswers")
	authorInfo := getUserJWT(c,"")
	userId := authorInfo.UserID
	fmt.Println(userId)

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	var qa []QuestionAnswer

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")

	opMatch := bson.M{"$match": bson.M{"answer.author.userid":userId}}
	
	opProject1 := bson.M{"$project":bson.M{"_id": 1, "title": 1,"author": 1, "postdatetime": 1,"qaviews": 1, "answercount": 1,"answer":bson.M{"$filter": bson.M{"input": "$answer", "as": "answer", "cond": bson.M{"$eq": []interface{}{"$$answer.status","Active"}}}}}}
	opProject2 := bson.M{"$project":bson.M{"_id": 1, "title": 1,"author": 1, "postdatetime": 1,"qaviews": 1, "answercount": 1,"answer.answerid":1,"answer.text":1,"answer.imageurl":1,"answer.videourl":1,"answer.author":1,"answer.postdatetime":1,"answer.emoxscale":1,"answer.commentcount":1,"answer.emoxcount":1}}	
		operations := []bson.M{opMatch,opProject1,opProject2}
	pipe := collection.Pipe(operations)
	err := pipe.All(&qa)
	if err != nil {
		fmt.Println("My Answer result err " + err.Error())
	}
	qaMyAnswerResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": qa}
	resp, err := json.Marshal(qaMyAnswerResult)
	if err != nil {
		fmt.Println("My Answer REsult " + err.Error())
	}

	//c.JSON(iris.StatusOK, resp)
	return string(resp)

}

func AddComment(c *iris.Context, mongoSession *mgo.Session) (addCommentResp string) {
	//Retrieve all the form data
	commentDataRequest := QACommentForm{}
	err := c.ReadForm(&commentDataRequest)

	if err != nil {
		fmt.Println("Error when reading add comment form: " + err.Error())
	}

	//Validate Comment Form data
	isValid, errMsg := validateQACommentData(commentDataRequest)

	if !isValid {
		qaValidationResult := map[string]interface{}{"code": util.CODE_QA502, "message": "Error", "result": errMsg}
		validationResp, err := json.Marshal(qaValidationResult)
		if err != nil {
			fmt.Println("Add Comment Validation Error " + err.Error())
		}
		return string(validationResp)
	}

	//Get user details
	authorData := getUserJWT(c, commentDataRequest.Author)

	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	// Get access to qa collection
	collection := sessionCopy.DB(util.DATABASENAME).C("qacomments")

	//Add QA comment
	commentData := QAComment{
		CommentID:    bson.NewObjectId(),
		QuestionID:   bson.ObjectIdHex(commentDataRequest.QuestionID),
		AnswerID:     bson.ObjectIdHex(commentDataRequest.AnswerID),
		PostText:     commentDataRequest.PostText,
		Author:       authorData,
		PostDateTime: time.Now(),
	}
	err = collection.Insert(&commentData)

	if err != nil {
		fmt.Println("ERROR while posting comments for an answer" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_QA102, "message": "Error", "result": "Error adding comments for an answer"}
		errResp, _ := json.Marshal(errRsponse)

		return string(errResp)
	}
	//increase comment count collection
	collection = sessionCopy.DB(util.DATABASENAME).C("questionanswers")

	opMatchQuery := bson.M{"answer.answerid": bson.ObjectIdHex(commentDataRequest.AnswerID)}
	opUpdate := bson.M{"$inc": bson.M{"answer.$.commentcount": 1}}
	err = collection.Update(opMatchQuery, opUpdate)

	if err != nil {
		fmt.Println("Add single comment to the qa" + err.Error())
	}

	//QA Comment added successfully , respond success json to the client
	qaAddCommentResult := map[string]interface{}{"code": util.CODE_QA101, "message": "Success", "result": "QA Comment Added Successfully"}
	resp, err := json.Marshal(qaAddCommentResult)
	if err != nil {
		fmt.Println("ERROR creating JSON data for posting qa comments" + err.Error())
	}

	return string(resp)
}

//Validate Add comment form data
func validateQACommentData(comment QACommentForm) (isValid bool, errorMessage string) {
	vErrors := []string{
		"Comment cannot be blank",
		"Comment cannot be more than 5000 characters",
	}
	postText := comment.PostText
	commentChCount := strings.Count(postText, "")

	//Check comment post
	if comment.PostText == "" { //checkIsBlank
		return false, vErrors[0]
	} else if commentChCount >= 5000 { //checkWordCount
		return false, vErrors[1]
	}
	//Valid form data
	return true, ""
}

func ReviseQuestion(c *iris.Context, mongoSession *mgo.Session) (reviseQuestionResult string) {
	fmt.Println("inside Add Revision Question")
	//Retrieve all the form data
	revision := AddRevisionData{}
	err := c.ReadForm(&revision)
	fmt.Println("Question revision " + revision.RevisionText)
	fmt.Println("Question id " + revision.Qid)
	if err != nil {
		fmt.Println("Error when reading Add revision form: " + err.Error())
	}

	//Validate Form data
	isValid, errMsg := validateAddRevision(revision)

	if !isValid {
		qaValidationResult := map[string]interface{}{"code": util.CODE_QA502, "message": "Error", "result": errMsg}
		validationResp, err := json.Marshal(qaValidationResult)
		if err != nil {
			fmt.Println("Add Question Revision Validation Error " + err.Error())
		}
		return string(validationResp)
	} else { //Valida Form data
		//Get a mongo db connection
		sessionCopy := mongoSession.Copy()
		defer sessionCopy.Close()

		// Get a collection to execute the query against.
		collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")

		//authorInfo := getUserJWT(c, revision.Author)

		//Add question revision
		questionrevisiondata := []QuestionRevision{QuestionRevision{RevisionText: revision.RevisionText, RevisionDateTime: time.Now()}}
		qa := &QuestionAnswer{Revision: questionrevisiondata}
		qid := bson.ObjectIdHex(revision.Qid)
		data, _ := json.MarshalIndent(qa, "", "    ")
		fmt.Println("DATA -" + string(data))

		matchQueri := bson.M{"_id": qid}
		change := bson.M{"$push": bson.M{"revision": bson.M{"$each": questionrevisiondata}}}
		err = collection.Update(matchQueri, change)

		if err != nil {
			fmt.Println("Adding Revision " + err.Error())
		}

		//Revision added successfully , respond success json to the client
		qaAddCommentResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": "Question Revision Added Successfully"}
		resp, err := json.Marshal(qaAddCommentResult)
		if err != nil {
			fmt.Println("Add Comment REsult " + err.Error())
		}

		return string(resp)
	}
}

//Validate Add Question Revision form data
func validateAddRevision(revision AddRevisionData) (isValid bool, errorMessage string) {
	vErrors := []string{
		"Revision text cannot be blank",
		"Revision text cannot be more than 500 characters",
	}
	revisiontextstring := revision.RevisionText
	rCharacterCount := strings.Count(revisiontextstring, "")
	fmt.Println("Question Revision character", rCharacterCount)
	//Check Question Revision text
	if revision.RevisionText == "" { //checkIsBlank
		return false, vErrors[0]
	} else if rCharacterCount >= 500 { //checkWordCount
		return false, vErrors[1]
	}
	//Valid form data
	return true, ""
}

func DeleteAnswer(c *iris.Context, mongoSession *mgo.Session) (delAnswerResponse string) {
	questionID := bson.ObjectIdHex(c.Param("qid"))
	answerID := bson.ObjectIdHex(c.Param("aid"))

	//Get user from Token
	authorInfo := getUserJWT(c, "")

	//Check if owner
	isOwner := IsAnswerOwner(answerID.String(), authorInfo.UserID.String(), c, mongoSession)

	if !isOwner { //Cannot delete the answer if not the owner
		qaDeleteAnswerAccessResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "Result": "Do not have permissions to delete the answer"}
		resp, err := json.Marshal(qaDeleteAnswerAccessResult)
		if err != nil {
			fmt.Println("" + err.Error())
		}
		return string(resp)
	}

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	matchQueri := bson.M{"$and": []interface{}{bson.M{"_id": questionID}, bson.M{"answer.answerid": answerID}}}
	change := bson.M{"$set": bson.M{"answer.$.status": "InActive"}}

	err := collection.Update(matchQueri, change)

	if err != nil {
		fmt.Println("Error deleting the answer" + err.Error())
	}
	qaDeleteAnswerResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "Result": "Answer deleted successfully"}
	resp, err := json.Marshal(qaDeleteAnswerResult)
	if err != nil {
		fmt.Println("Error in JSON format for delete answer" + err.Error())
	}

	return string(resp)
}


func GetQAEmox(c *iris.Context, mongoSession *mgo.Session) (qaEmox string) {
	//Get all the comments for a qa
	questionId := bson.ObjectIdHex(c.Param("qid"))

	answerId := bson.ObjectIdHex(c.Param("aid"))

	fmt.Println("QUESTIONID-", questionId)
	fmt.Println("ANSWERID-", answerId)
	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var qaEmoxData []QAEmoxScale
	//var questionanswers []QuestionAnswer
	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")

	
	opProject1 := bson.M{"$unwind": "$answer"}

	opMatch := bson.M{"$match": bson.M{"answer.answerid": answerId}}
	opProject := bson.M{
		"$project": bson.M{"questionid": "$_id", "answerid": "$answer.answerid", "emoxscale": "$answer.emoxscale"},
	}
	opMain := []bson.M{opProject1, opMatch, opProject}
	pipe := collection.Pipe(opMain)
	err := pipe.All(&qaEmoxData)
	fmt.Println("MAIN - ", opMain)
	if err != nil {
		fmt.Println("Error in finding qa emoxs" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_QA504, "message": "Error", "result": "Error in finding emoxs for qa"}
		errResp, _ := json.Marshal(errRsponse)

		return string(errResp)
	}

	qaEmoxResult := map[string]interface{}{"code": util.CODE_QA504, "message": "Success", "result": qaEmoxData}
	//qaEmoxResult := map[string]interface{}{"code": util.CODE_ST415, "message": "Success", "result": questionanswers}
	resp, err := json.Marshal(qaEmoxResult)
	if err != nil {
		fmt.Println("Error in creating JSON format for qa Emox API" + err.Error())
	}

	return string(resp)
}

func GetQAFeed(c *iris.Context, mongoSession *mgo.Session) (qaFeedResponse string) {
	qaTags := c.Param("tags")
	qaTagsList := strings.Split(qaTags, ",")
	
	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var qa []QuestionAnswer
	authu := getUserJWT(c, "")
	fmt.Println("\n\n\nGET USER DETAILS-", authu)

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	tagQueryFmt := make([]bson.RegEx, len(qaTagsList))
	opMatch := bson.M{}
	if qaTags == "" || qaTags == "\"\"" {
		opMatch = bson.M{"$match": bson.M{"$and": []bson.M{
			bson.M{"status": "Active"},
			//bson.M{"$ne": bson.M{"answercount": 0}},
			bson.M{"answercount": bson.M{"$ne": 0}},
		},
		}}
	} else {
		for tw := range qaTagsList {
			tagQueryFmt[tw] = bson.RegEx{qaTagsList[tw], "i"}
		}
		opMatch = bson.M{"$match": bson.M{"qatags": bson.M{"$in": tagQueryFmt}, "status": "Active", "answercount": bson.M{"$ne": 0}}}
	}

	opProject := bson.M{"$project": bson.M{"_id": 1, "title": 1, "postdatetime": 1, "qaviewscount": 1, "answercount": 1, "answer": 1}}
	
	opSort1 := bson.M{"$sort": bson.M{"postdatetime": -1}}
	opSort2 := bson.M{"$sort": bson.M{"_id": -1}}
	opMatch1 := bson.M{}

	questionID := c.Param("qid")

	if bson.IsObjectIdHex(questionID) != true { //not valid object id
		opMatch1 = bson.M{"$match": bson.M{"_id": bson.M{"$exists": true}}}
		fmt.Println("if questionid GetQAFeed")
	} else { //VALID Object id
		lastQuestionId := bson.ObjectIdHex(questionID)
		fmt.Println(lastQuestionId)
		//opMatch1 = bson.M{"$match":bson.M{"$gt":[]interface{}{"$_id",lastQuestionId}}}
		opMatch1 = bson.M{"$match": bson.M{"_id": bson.M{"$lt": lastQuestionId}}}

		fmt.Println("else questionid GetQAFeed")
	}
	opProject2 := bson.M{"$limit": 30}
	

	opMain := []bson.M{opMatch, opProject, opSort1, opSort2, opMatch1, opProject2}

	pipe := collection.Pipe(opMain)
	err := pipe.All(&qa)

	if err != nil {
		fmt.Println("Question with top answer result err " + err.Error())
	}

	qatopanswerResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": qa}
	resp, err := json.Marshal(qatopanswerResult)
	if err != nil {
		fmt.Println("Top QA REsult " + err.Error())
	}

	return string(resp)
}

func GetQAAnswer(c *iris.Context, mongoSession *mgo.Session) (qaAnswerResponse string) {
	qaTags := c.Param("tags")
	qaTagsList := strings.Split(qaTags, ",")
	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var qa []QuestionAnswer

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	tagQueryFmt := make([]bson.RegEx, len(qaTagsList))
	opMatch := bson.M{}

	if qaTags == "" || qaTags == "\"\"" {
		opMatch = bson.M{"$match": bson.M{"status": "Active", "answercount": 0}}
	} else {
		for tw := range qaTagsList {
			tagQueryFmt[tw] = bson.RegEx{qaTagsList[tw], "i"}
		}
		opMatch = bson.M{"$match": bson.M{"qatags": bson.M{"$in": tagQueryFmt}, "status": "Active", "answercount": 0}}
	}

	opProject := bson.M{"$project": bson.M{"_id": 1, "title": 1, "postdatetime": 1, "qaviewscount": 1}}
	opSort1 := bson.M{"$sort": bson.M{"postdatetime": -1}}
	opSort2 := bson.M{"$sort": bson.M{"_id": -1}}

	opMatch1 := bson.M{}

	questionID := c.Param("qid")

	if bson.IsObjectIdHex(questionID) != true { //not valid object id
		//opMatch1 = bson.M{"$match": bson.M{"status": "Active"}}
		opMatch1 = bson.M{"$match": bson.M{"_id": bson.M{"$exists": true}}}
		fmt.Println("if questionid")
	} else {
		lastQuestionId := bson.ObjectIdHex(questionID)
		fmt.Println(lastQuestionId)
		//opMatch1 = bson.M{"$match":bson.M{"$gt":[]interface{}{"$_id",lastQuestionId}}}
		opMatch1 = bson.M{"$match": bson.M{"_id": bson.M{"$lt": lastQuestionId}}}

		fmt.Println("else questionid")
	}
	opProject2 := bson.M{"$limit": 30}
	opMain := []bson.M{opMatch, opProject, opSort1,opSort2, opMatch1, opProject2}

	pipe := collection.Pipe(opMain)
	err := pipe.All(&qa)

	if err != nil {
		fmt.Println("Question with top answer result err " + err.Error())
	}

	qatopanswerResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": qa}
	resp, err := json.Marshal(qatopanswerResult)
	if err != nil {
		fmt.Println("Top QA REsult " + err.Error())
	}

	return string(resp)
}

func GetQATrends(c *iris.Context, mongoSession *mgo.Session) (qaTrendResponse string) {
	qaTags := c.Param("tags")
	qaTagsList := strings.Split(qaTags, ",")
	pageno, _ := strconv.Atoi(c.Param("pageno"))
	limit := 30
	skip := 30 * (pageno - 1)
	fmt.Println("skip", skip)
	fmt.Println("PAGE NO-", pageno)

	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	var qa []QuestionAnswer

	//var qa []UnwindQuestionAnswer

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	tagQueryFmt := make([]bson.RegEx, len(qaTagsList))
	opMatch := bson.M{}

	if qaTags == "" || qaTags == "\"\"" {
		opMatch = bson.M{"$match": bson.M{"status": "Active"}}
	} else {
		for tw := range qaTagsList {
			tagQueryFmt[tw] = bson.RegEx{qaTagsList[tw], "i"}
		}
		opMatch = bson.M{"$match": bson.M{"qatags": bson.M{"$in": tagQueryFmt}, "status": "Active"}}
	}

	opProject2 := bson.M{"$unwind": "$answer"}
	opProject3 := bson.M{"$project": bson.M{"_id": 1, "title": 1, "author": 1, "postdatetime": 1, "answercount": 1, "answer.answerid": "$answer.answerid", "answer.text": "$answer.text", "answer.imageurl": "$answer.imageurl", "answer.author": "$answer.author", "answer.postdatetime": "$answer.postdatetime", "answer.emoxscale": "$answer.emoxscale", "answer.popularcount": bson.M{"$add": []interface{}{bson.M{"$multiply": []interface{}{"$answer.emoxcount", 2}}, bson.M{"$multiply": []interface{}{"$answer.commentcount", 3}}}}}}
	opProject4 := bson.M{"$sort": bson.M{"answer.popularcount": -1}}
	opProject5 := bson.M{"$group": bson.M{"_id": "$_id", "title": bson.M{"$first": "$title"}, "answercount": bson.M{"$first": "$answercount"}, "postdatetime": bson.M{"$first": "$postdatetime"}, "revision": bson.M{"$first": "$revision"}, "author": bson.M{"$first": "$author"}, "answer": bson.M{"$push": "$answer"}}}
	opProject6 := bson.M{"$sort": bson.M{"answercount": -1}}
	opProject7 := bson.M{"$project": bson.M{"_id": 1, "title": 1, "revision": 1, "author": 1, "postdatetime": 1, "status": 1, "qatags": 1, "qaviews": 1, "answercount": 1, "answer": bson.M{"$slice": []interface{}{"$answer", 1}}}}
	opProject8 := bson.M{"$skip": skip}
	opProject9 := bson.M{"$limit": limit}

	opMain := []bson.M{opMatch, opProject2, opProject3, opProject4, opProject5, opProject6, opProject7, opProject8, opProject9}
	pipe := collection.Pipe(opMain)
	err := pipe.All(&qa)

	if err != nil {
		fmt.Println("Question with top answer result err " + err.Error())
	}

	qatopanswerResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": qa}
	resp, err := json.Marshal(qatopanswerResult)
	if err != nil {
		fmt.Println("Top QA REsult " + err.Error())
	}

	return string(resp)
}

func AddQuestionWeb(c *iris.Context, mongoSession *mgo.Session) {
	fmt.Println("inside Add Question through web")
	//Retrieve all the form data
	question := QuestionWebData{}
	err := c.ReadForm(&question)
	fmt.Println("question " + question.Title)
	layout := "2006-01-02T15:04"

	questionDateTime, err := time.Parse(layout, question.QuestionAddedDate)
	if err != nil {
		fmt.Println("answer time issue" + err.Error())
	}
	// p(t1)
	fmt.Println("answer image " + questionDateTime.String())

	if err != nil {
		fmt.Println("Error when reading Add question form:" + err.Error())
	}

	//Validate Form data
	isValid, errMsg := validateAddQuestionWeb(question)

	if !isValid {
		qaValidationResult := map[string]interface{}{"code": util.CODE_QA502, "message": "Error", "result": errMsg}
		validationResp, err := json.Marshal(qaValidationResult)
		if err != nil {
			fmt.Println("Add Question through web Validation Error " + err.Error())
		}
		c.JSON(iris.StatusOK, validationResp)
	} else { //Valida Form data
		//Get a mongo db connection
		sessionCopy := mongoSession.Copy()
		defer sessionCopy.Close()

		// Get a collection to execute the query against.
		collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")

		index := mgo.Index{
			Key:        []string{"title"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}

		err = collection.EnsureIndex(index)
		if err != nil {
			fmt.Println("Error Creating index" + err.Error())
		}
		//Add QA Tags
		qaRelatedTags, err := GenerateQATags(question.Title, c, mongoSession)
		if err != nil {
			fmt.Println("GENERATETAG ERRORS" + err.Error())
		}

		fmt.Println("AUTHOR-" + question.Name)

		qUserID := bson.ObjectIdHex(question.UserID)
		aAuthorInfo := Author{UserID: qUserID, Name: question.Name, ImageURL: question.ImageURL, SingleLineDesc: question.SingleLineDesc, Followers: question.Followers}

		qDataInfo := &QuestionAnswer{QAID: bson.NewObjectId(), Title: question.Title, PostDateTime: questionDateTime, Status: "Active", Author: aAuthorInfo, QATags: qaRelatedTags[0].Keyword, Category: qaRelatedTags[0].Category}
		data, _ := json.MarshalIndent(qDataInfo, "", "    ")

		fmt.Println("DATA -" + string(data))
		err = collection.Insert(qDataInfo)

		if err != nil {
			fmt.Println("Adding question err " + err.Error())
			qaAddQuestionResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Fail", "result": err.Error()}
			resp, err := json.Marshal(qaAddQuestionResult)
			if err != nil {
				fmt.Println("Add Question REsult err " + err.Error())
			}
			c.Write(string(resp))
		}

		//Question added successfully , respond success json to the client
		qaAddQuestionResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": "Question Added Successfully"}
		resp, err := json.Marshal(qaAddQuestionResult)
		if err != nil {
			fmt.Println("Add Question Web REsult " + err.Error())
		}

		//c.JSON(iris.StatusOK, resp)
		c.Write(string(resp))
	}
}

//Validate Add question form data
func validateAddQuestionWeb(question QuestionWebData) (isValid bool, errorMessage string) {
	vErrors := []string{
		"Question title cannot be blank",
		"Question title cannot be more than 375 character",
	}
	questiontitlestring := question.Title
	qCharacterCount := strings.Count(questiontitlestring, "")
	fmt.Println(qCharacterCount)
	/*qWordList := strings.Split(question.Title, " ")
	qWordCount := len(qWordList)*/

	//Check question title
	if question.Title == "" { //checkIsBlank
		return false, vErrors[0]
	} else if qCharacterCount >= 375 { //checkWordCount
		return false, vErrors[1]
	}
	//Valid form data
	return true, ""
}

func AddAnswerToQuestionWeb(c *iris.Context, mongoSession *mgo.Session) (addAnswerToQuestionWebResult string) {
	fmt.Println("inside Add Answer")

	//Retrieve all the form data
	answer := AnswerWebData{}
	err := c.ReadForm(&answer)
	fmt.Println("answer " + answer.Text)
	layout := "2006-01-02T15:04"

	answerDateTime, err := time.Parse(layout, answer.AnswerAddedDate)
	if err != nil {
		fmt.Println("answer time issue" + err.Error())
	}
	// p(t1)
	fmt.Println("answer image " + answerDateTime.String())
	//	fmt.Println("date time"+time.Date(answer.AnswerAddedDate))
	if err != nil {
		fmt.Println("Error when reading Add answer form: " + err.Error())
	}

	//Validate Form data
	isValid, errMsg := validateAddAnswerWeb(answer)

	if !isValid {
		qaValidationResult := map[string]interface{}{"code": util.CODE_QA502, "message": "Error", "result": errMsg}
		validationResp, err := json.Marshal(qaValidationResult)
		if err != nil {
			fmt.Println("Add Answer Validation Error " + err.Error())
		}
		//c.JSON(iris.StatusOK, validationResp)
		return string(validationResp)
	} else { //Valida Form data
		//Get a mongo db connection
		sessionCopy := mongoSession.Copy()
		defer sessionCopy.Close()

		// Get a collection to execute the query against.
		collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
		aUserID := bson.ObjectIdHex(answer.UserID)
		aAuthorInfo := Author{UserID: aUserID, Name: answer.Name, ImageURL: answer.ImageURL, SingleLineDesc: answer.SingleLineDesc, Followers: answer.Followers}

		//Add answer
		answerinfo := []Answer{Answer{AnswerID: bson.NewObjectId(), Text: answer.Text, ImageURL: answer.AnswerImageURL, VideoURL: answer.VideoURL, Author: aAuthorInfo, Status: "Active", PostDateTime: answerDateTime}}
		qa := &QuestionAnswer{Answer: answerinfo}
		fmt.Println("Question Id",answer.Qid)
//	if	bson.IsObjectIdHex(answer.Qid) != true{
//			fmt.Println("plz paas a valid id")
//		}
			qid := bson.ObjectIdHex(answer.Qid)
		fmt.Println("qId",qid)

		data, _ := json.MarshalIndent(qa, "", "    ")
		fmt.Println("DATA -" + string(data))
		matchQueri := bson.M{"_id": qid}
		change := bson.M{"$push": bson.M{"answer": bson.M{"$each": answerinfo}}}
		err = collection.Update(matchQueri, change)

		if err != nil {
			fmt.Println("Adding Answer " + err.Error())
			errRsponse := map[string]interface{}{"code": util.CODE_QA102, "message": "Error", "result": "Error adding answer to qa"}
			errResp, _ := json.Marshal(errRsponse)

			return string(errResp)

		}
		collection = sessionCopy.DB(util.DATABASENAME).C("questionanswers")

		opMatchQuery := bson.M{"_id": qid}
		opUpdate := bson.M{"$inc": bson.M{"answercount": 1}}
		err = collection.Update(opMatchQuery, opUpdate)

		if err != nil {
			fmt.Println("Add answercount to the qa" + err.Error())
		}

		//Answer added successfully , respond success json to the client
		qaAddAnswerResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": "Answer Added Successfully"}
		res, err := json.Marshal(qaAddAnswerResult)
		if err != nil {
			fmt.Println("Add Answer REsult " + err.Error())
		}

		//c.JSON(iris.StatusOK, res)
		return string(res)
	}
}

//Validate Add answer form data
func validateAddAnswerWeb(answer AnswerWebData) (isValid bool, errorMessage string) {
	vErrors := []string{
		"Answer text cannot be blank",
		"Answer text cannot be more than 10000 characters",
	}
	answerstring := answer.Text
	aCharacterCount := strings.Count(answerstring, " ")
	/*qWordList := strings.Split(answer.Text, " ")
	qWordCount := len(qWordList)*/

	//Check answer text
	if answer.Text == "" { //checkIsBlank
		return false, vErrors[0]
	} else if aCharacterCount >= 10000 { //checkWordCount
		return false, vErrors[1]
	}
	//Valid form data
	return true, ""
}

//func UploadImage(c *iris.Context,imageType string) (imageName string) {
//	fmt.Println(imageType)
//	AWSS3setup :=	config.AmazonS3Bucket()
//	creds := credentials.NewStaticCredentials(AWSS3setup.AWSAccessKeyID, AWSS3setup.AWSSecretAccessKey, AWSS3setup.Token)
//	config := &aws.Config{
//		Region:      aws.String("ap-south-1"),
//		Credentials: creds,
//	}
//	sess := session.New(config)
//	uploader := s3manager.NewUploader(sess)
//	t := time.Now()
//	datetime := t.Format("Mon Jan _2 15:04:05 2006")
//	fmt.Println(" time" + datetime)
//	uploadfile,err := c.FormFile("ImageURL")

//	if err != nil {
//		fmt.Println("err uploding file:", err)
//	}
//	fmt.Println("uploadfilevalue",uploadfile)
	
//	if uploadfile !=nil{
//		filename := uploadfile.Filename
//		header := uploadfile.Header
		
//		//fmt.Println("body",body)
//		fmt.Println("header",header)
//		fmt.Println("file name" + filename)
//		//changeFile := datetime + filename
//		changeFile :=  filename
//		fmt.Println("change file" + changeFile)
//		path := "/"+imageType+"/" + changeFile
//		sourceFile, err := uploadfile.Open()
//		if err != nil{
//		     fmt.Println("err open file",err)
            
//		}
//		defer sourceFile.Close()
//		fmt.Println("sourceFile",sourceFile)

//		image1, _, err := image.DecodeConfig(sourceFile)

//    	if err != nil {
//	 	fmt.Println("err in decode image",err.Error())
//    	}
//    fmt.Println("image1",image1) 
////	newImage := resize.Resize(160, 0, image1, resize.Lanczos3)
//    fmt.Println("ColorModel",image1.ColorModel) 
//	fmt.Println("width",image1.Width)
//    fmt.Println("height",image1.Height) 

//	permission := "public-read"

//	result, err := uploader.Upload(&s3manager.UploadInput{
//		Body:   sourceFile,
//		Bucket: aws.String(AWSS3setup.Bucket),
//		Key:    aws.String(path),
//		ACL: aws.String(permission),
//		Metadata: map[string]*string{
//        "Key": aws.String("MetadataValue"),
//                 },
//	})
//	if err != nil {
//		log.Fatalln("Failed to upload", err)
//	}

//	fmt.Printf("response ", result)

//	log.Println("Successfully uploaded to", result.Location)

//	image := result.Location
//	fmt.Println("image" + image)
//	return image
//	}
//	return ""
//}

func UploadImage(c *iris.Context,imageType string) (imageName string) {
	fmt.Println(imageType)
	AWSS3setup :=	config.AmazonS3Bucket()
	creds := credentials.NewStaticCredentials(AWSS3setup.AWSAccessKeyID, AWSS3setup.AWSSecretAccessKey, AWSS3setup.Token)
	config := &aws.Config{
		Region:      aws.String("ap-south-1"),
		Credentials: creds,
	}
	sess := session.New(config)
	uploader := s3manager.NewUploader(sess)
	t := time.Now()
	datetime := t.Format("Mon Jan _2 15:04:05 2006")
	fmt.Println(" time" + datetime)
	uploadfile,err := c.FormFile("ImageURL")

	if err != nil {
		fmt.Println("err uploding file:", err)
	}
	fmt.Println("uploadfilevalue",uploadfile)
	
	if uploadfile !=nil{
		filename := uploadfile.Filename
		header := uploadfile.Header
		
		//fmt.Println("body",body)
		fmt.Println("header",header)
		fmt.Println("file name" + filename)
		//changeFile := datetime + filename
		changeFile :=  filename
		fmt.Println("change file" + changeFile)
		path := "/"+imageType+"/" + changeFile
		sourceFile, err := uploadfile.Open()
		if err != nil{
		     fmt.Println("err open file",err)
            
		}
		fmt.Println("sourceFile",sourceFile)
//		m, _, err := image.Decode(sourceFile)
//				if err != nil {
//					log.Fatal(err)
//				}
//				bounds := m.Bounds()
		image1, _, err := image.DecodeConfig(sourceFile)

    	if err != nil {
	 	fmt.Println("err in decode image",err.Error())
    	}

		
	//m := imaging.Resize(sourceFile, 128, 128, imaging.Lanczos)
//		fmt.Println("m",m)
//		fmt.Println("bounds",bounds)

		defer sourceFile.Close()

    fmt.Println("image1",image1) 
    fmt.Println("ColorModel",image1.ColorModel) 
	fmt.Println("width",image1.Width)
    fmt.Println("height",image1.Height) 

	permission := "public-read"

	result, err := uploader.Upload(&s3manager.UploadInput{
		Body:   sourceFile,
		Bucket: aws.String(AWSS3setup.Bucket),
		Key:    aws.String(path),
		ACL: aws.String(permission),
		Metadata: map[string]*string{
        "Key": aws.String("MetadataValue"),
                 },
	})
	if err != nil {
		log.Fatalln("Failed to upload", err)
	}

	fmt.Printf("response ", result)

	log.Println("Successfully uploaded to", result.Location)

	image := result.Location
	fmt.Println("image" + image)
	return image
	}
	return ""
}

//func UploadImage1(c *iris.Context,imageType string) (imageName string) {
//	fmt.Println(imageType)
//	h := md5.New()
//	content := strings.NewReader("") 
//	content.WriteTo(h)
//	AWSS3setup :=	config.AmazonS3Bucket()
//	creds := credentials.NewStaticCredentials(AWSS3setup.AWSAccessKeyID, AWSS3setup.AWSSecretAccessKey, AWSS3setup.Token)
	
//	config := &aws.Config{
//		Region:      aws.String("ap-south-1"),
//		Credentials: creds,
//	}
//	svc := s3.New(session.New(config))

//	uploadfile,err := c.FormFile("ImageURL")

//	if err != nil {
//		fmt.Println("err uploding file:", err)
//	}
//	fmt.Println("uploadfilevalue",uploadfile)
	
//	if uploadfile !=nil{
//		filename := uploadfile.Filename
//		path := "/"+imageType+"/" + filename
//		sourceFile, err := uploadfile.Open()
//		if err != nil{
//		     fmt.Println("err open file",err)
            
//		}
//		defer sourceFile.Close()
	
//	fileUpload, err := svc.PutObject(&s3.PutObjectInput{
//    		Body:   sourceFile,
//			Bucket: aws.String(AWSS3setup.Bucket),
//			Key:    aws.String(path),
			
//})
//	if err != nil {
//	    log.Printf("Failed to upload data to %s/%s, %s\n", AWSS3setup.Bucket, path, err)
//	    return
//	}


//		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
//		Bucket: aws.String(AWSS3setup.Bucket),
//		Key:    aws.String(path),
//		})
		
//		md5s := base64.StdEncoding.EncodeToString(h.Sum(nil)) 
//		req.HTTPRequest.Header.Set("Content-MD5", md5s) 
//		urlStr, err := req.Presign(15 * time.Minute)  
//		if err != nil { 
//		    fmt.Println("error presigning request", err) 
//		    return 
//		}
		
////		req1, err := http.NewRequest("PUT", urlStr, strings.NewReader("")) 
////		req1.Header.Set("Content-MD5", md5s) 
////		if err != nil { 
////		    fmt.Println("error creating request", urlStr) 
////		    return 
////		} 
		
////		resp, err := http.DefaultClient.Do(req1)
////		fmt.Println("resp",resp, err) 
	
//	fmt.Println("urlstr",urlStr) 

//fmt.Println("uploadResult",fileUpload)
//log.Printf("Successfully created bucket %s and uploaded data with key %s\n", AWSS3setup.Bucket, path)

//	return urlStr
//	}
//	return ""
//}

//func GetImageUrl(c *iris.Context,imageName string) (imageURL string) {
//	fmt.Println(imageName)
//	if imageName != ""{
//	AWSS3setup :=	config.AmazonS3Bucket()
//	creds := credentials.NewStaticCredentials(AWSS3setup.AWSAccessKeyID, AWSS3setup.AWSSecretAccessKey, AWSS3setup.Token)
	
//	config := &aws.Config{
//		Region:      aws.String("ap-south-1"),
//		Credentials: creds,
//	}
//	svc := s3.New(session.New(config))
//	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
//		Bucket: aws.String(AWSS3setup.Bucket),
//		Key:    aws.String(imageName),
//		})
//		urlStr, err := req.Presign(15 * time.Minute)

//		if err != nil {
//		    log.Println("Failed to sign request", err)
//		}
		
//	log.Println("The URL is", urlStr)
	
//log.Printf("Successfully created bucket %s and uploaded data with key %s\n", AWSS3setup.Bucket, imageName)

//	return urlStr
//	}
//	return ""
//}
func GetQuestionsListWeb(c *iris.Context, mongoSession *mgo.Session) (qaList string) {
	c.SetHeader("Access-Control-Allow-Origin", "*")
	qaTags := c.Param("tags")
	qaTagsList := strings.Split(qaTags, ",")
	questionID := c.Param("qid")

	//fmt.Println("PAGE NO-", pageno)

	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var qa []QuestionAnswer

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	tagQueryFmt := make([]bson.RegEx, len(qaTagsList))
	opMatch := bson.M{}
	opMatch1 := bson.M{}

	if qaTags == "" || qaTags == "\"\"" {
		opMatch = bson.M{"$match": bson.M{"status": "Active"}}
	} else {
		for tw := range qaTagsList {
			tagQueryFmt[tw] = bson.RegEx{qaTagsList[tw], "i"}
		}

		opMatch = bson.M{"$match": bson.M{"qatags": bson.M{"$in": tagQueryFmt}, "status": "Active"}}
	}
	if bson.IsObjectIdHex(questionID) != true {
		opMatch1 = bson.M{"$match": bson.M{"status": "Active"}}
		fmt.Println("if questionid")
	} else {
		lastQuestionId := bson.ObjectIdHex(questionID)
		fmt.Println(lastQuestionId)
		//opMatch1 = bson.M{"$match":bson.M{"$gt":[]interface{}{"$_id",lastQuestionId}}}
		opMatch1 = bson.M{"$match": bson.M{"_id": bson.M{"$lt": lastQuestionId}}}

		fmt.Println("else questionid")
	}

	opProject := bson.M{"$project": bson.M{"_id": 1, "title": 1, "author": 1, "postdatetime": 1, "status": 1, "qatags": 1, "category": 1, "answer": 1}}
	opProject1 := bson.M{"$sort": bson.M{"_id": -1}}
	opProject2 := bson.M{"$limit": 3000}

	//	opMain := []bson.M{opMatch, opProject}
	opMain := []bson.M{opMatch, opProject, opProject1, opMatch1, opProject2}

	pipe := collection.Pipe(opMain)
	err := pipe.All(&qa)

	if err != nil {
		fmt.Println("Question with top answer result err " + err.Error())
	}

	qatopanswerResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": qa}
	resp, err := json.Marshal(qatopanswerResult)
	if err != nil {
		fmt.Println("Top QA REsult " + err.Error())
	}

	return string(resp)
}

func GetCheckEmoxExist(c *iris.Context, mongoSession *mgo.Session) (GetCheckEmoxExistResult string) {
	fmt.Println("Check qa emox")
	Qid := bson.ObjectIdHex(c.Param("qid"))
	Aid := bson.ObjectIdHex(c.Param("aid"))
	
	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	
	collection := sessionCopy.DB(util.DATABASENAME).C("qaemoxs")
	authorInfo := getUserJWT(c,"")
	 QaEmox	:=	QAEmox{}
	fmt.Println("questionid",Qid)
	fmt.Println("answerid",Aid)
	fmt.Println("userid",authorInfo.UserID)

	OpMatch := bson.M{"$match": bson.M{"questionid":Qid,"answerid":Aid,"userid":authorInfo.UserID}}
	OpProject := bson.M{"$project": bson.M{"emoxtext": 1}}

	operations := []bson.M{OpMatch, OpProject}
	pipe := collection.Pipe(operations)
	err := pipe.One(&QaEmox)
	//check emox 
	if err != nil {
		fmt.Println("ERROR During check emoxes for qa" + err.Error())
			
	}
	
	getqaEmoxResult := map[string]interface{}{"code": util.CODE_QA503, "message": "Success", "result":QaEmox}
	resp, err := json.Marshal(getqaEmoxResult)
	if err != nil {
		fmt.Println("ERROR creating JSON data for qa emox" + err.Error())
	}

	return string(resp)
	
}

func IsAnswerOwner(answerID string, userID string, c *iris.Context, mongoSession *mgo.Session) (owner bool) {
	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	user := []IsOwner{}

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	opMatch := bson.M{"$match": bson.M{"status": "Active", "answer.answerid": answerID, "answer.author.userid": userID, "answer.status": "Active"}}

	opProject := bson.M{
		"$project": bson.M{"answer.answerid": 1},
	}

	opMain := []bson.M{opMatch, opProject}
	pipe := collection.Pipe(opMain)
	err := pipe.All(&user)

	if err != nil {
		fmt.Println("Error getting answer ownership details" + err.Error())
	}
	if len(user) > 0 {
		return true
	}

	return false
}

func IsQuestionOwner(questionID string, userID string, c *iris.Context, mongoSession *mgo.Session) (owner bool) {
	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	user := []IsOwner{}
	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	opMatch := bson.M{"$match": bson.M{"status": "Active", "questionid": questionID, "author.userid": userID}}

	opProject := bson.M{
		"$project": bson.M{"questionid": 1},
	}

	opMain := []bson.M{opMatch, opProject}
	pipe := collection.Pipe(opMain)
	err := pipe.All(&user)

	if err != nil {
		fmt.Println("Error Getting Question ownership details" + err.Error())
	}
	if len(user) > 0 {
		return true
	}

	return false
}

func GetTimeFilterQA(c *iris.Context,mongoSession *mgo.Session)(GetTimeFilterQAResult string){
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	var	question []QAFilter
	
	currentdate := time.Now()
	nextday := time.Now().Add(time.Hour * 24)
	previousday := time.Now().Add(-time.Hour * 24)
	previousweek := time.Now().Add(-time.Hour * 24*7)
	previousmonth := time.Now().Add(-time.Hour * 24*30)
	fmt.Println("currentdate",currentdate)
	fmt.Println("nextday",nextday)
	fmt.Println("previousday",previousday)
	fmt.Println("previousweek",previousweek)
	fmt.Println("previousmonth",previousmonth)
				
	qaTags := c.Param("tags")
	qaTagsList := strings.Split(qaTags, ",")
	Time := c.Param("time")
	tagQueryFmt := make([]bson.RegEx, len(qaTagsList))
	opMatch := bson.M{}
	if qaTags == "" || qaTags == "\"\"" {
		opMatch = bson.M{"$match": bson.M{"status": "Active"}}
	} else {
		for tw := range qaTagsList {
			tagQueryFmt[tw] = bson.RegEx{qaTagsList[tw], "i"}
		}

		opMatch = bson.M{"$match": bson.M{"qatags": bson.M{"$in": tagQueryFmt}, "status": "Active"}}
	}
		opProject := bson.M{}
	if Time == "day"{
		opProject = bson.M{"$project":bson.M{"_id":1,"title":1,"postdatetime":1,"author":1,"answercount":1,"qaviewscount":1,"time":bson.M{"$gte":[]interface{}{"$postdatetime",previousday}}}}

	}else{
	 if Time == "week"{
		opProject = bson.M{"$project":bson.M{"_id":1,"title":1,"postdatetime":1,"author":1,"answercount":1,"qaviewscount":1,"time":bson.M{"$gte":[]interface{}{"$postdatetime",previousweek}}}}
	}else{
	 if Time == "month"{
		opProject = bson.M{"$project":bson.M{"_id":1,"title":1,"postdatetime":1,"author":1,"answercount":1,"qaviewscount":1,"time":bson.M{"$gte":[]interface{}{"$postdatetime",previousmonth}}}}

			}
		}
	}
	opProject1 := bson.M{"$project":bson.M{"_id":1,"title":1,"postdatetime":1,"author":1,"answercount":1,"qaviewscount":1,"time":1}}
	opProject2 := bson.M{"$match":bson.M{"time":true}}
	opProject3 := bson.M{"$sort":bson.M{"postdatetime":-1}}
	operations := []bson.M{opMatch,opProject,opProject1,opProject2,opProject3}
	pipe := collection.Pipe(operations)
	err := pipe.All(&question)
	if err != nil{
		fmt.Println("get time filter qa",err.Error())
	}
	qaTimeFilterResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": question}
	resp, err := json.Marshal(qaTimeFilterResult)
	if err != nil {
		fmt.Println("Top QA REsult " + err.Error())
	}

	return string(resp)
	
}

func AddEmoxQA(c *iris.Context, mongoSession *mgo.Session) (addEmox string) {
	//Retrieve all the form data
	emox := QAEmox{}
	err := c.ReadForm(&emox)
	if err != nil {
		fmt.Println("Error when reading Add QA Emox form: " + err.Error())
	}

	//Validate Form data
	isValid, errMsg := validateAddQaEmox(emox)

	if !isValid {
		emoxValidationResult := map[string]interface{}{"code": util.CODE_QA503, "message": "Error", "result": errMsg}
		vResp, err := json.Marshal(emoxValidationResult)
		if err != nil {
			fmt.Println("ERROR during Emox QA validation" + err.Error())
		}
		return string(vResp)
	}

	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	// Get access to storyemoxs collection
	collection := sessionCopy.DB(util.DATABASENAME).C("qaemoxs")
	authorInfo := getUserJWT(c, "")
	checkResult ,preEmoxText	:= EmoxExist(emox,c, mongoSession)
	fmt.Println("check exox exist",checkResult)
	fmt.Println("check exox PreText",preEmoxText)
	//check already emox or not if not then add
	if checkResult != true{
	//Add Emox
	emoxData := QAEmox{
		QuestionID:   emox.QuestionID,
		AnswerID:     emox.AnswerID,
		EmoxText:     emox.EmoxText,
		ByUser:       authorInfo.UserID,
		EmoxDateTime: time.Now(),
	}
	//Add emox for user/story
	err = collection.Insert(&emoxData)

	if err != nil {
		fmt.Println("ERROR During add emoxes for qa" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_QA503, "message": "Error", "result": "Error in adding emoxes to qa"}
		errResp, _ := json.Marshal(errRsponse)
	
		return string(errResp)
		
	}
	AddAnswerEmoxCountResult := AddAnswerEmoxCount(emox,c , mongoSession)
		fmt.Println("addansweremoxcount",AddAnswerEmoxCountResult) 
	//QA Emox added successfully , respond success json to the client
	qaAddEmoxResult := map[string]interface{}{"code": util.CODE_QA503, "message": "Success", "result": "QA Emox Added Successfully"}
	resp, err := json.Marshal(qaAddEmoxResult)
	if err != nil {
		fmt.Println("ERROR creating JSON data for qa emox" + err.Error())
	}
	//Increment emox count and emoxScale
		IncrementEmoxCountResult := IncrementEmoxCount(emox,c, mongoSession)
		//emox increment then result true else false
		if IncrementEmoxCountResult != true{
			qaIncEmoxErr:= map[string]interface{}{"code": util.CODE_QA503, "message": "fail", "result": "emox increment err"}
			resp, err := json.Marshal(qaIncEmoxErr)
			if err != nil {
				fmt.Println("ERROR creating JSON data for qa  emox" + err.Error())
			}
					return string(resp)

		}
		return string(resp)

	}
	//if emox already present then update 
	if preEmoxText	!= emox.EmoxText{
			QAEmoxUpdateResult := QAEmoxUpdate(emox,c,mongoSession)
			IncrementDecrementEmoxCountResult  := DecrementIncrementEmoxCount(preEmoxText,emox,c, mongoSession)
			if IncrementDecrementEmoxCountResult != true{
			qaIncDecEmoxErr:= map[string]interface{}{"code": util.CODE_QA503, "message": "fail", "result": "emox increment err"}
			resp, err := json.Marshal(qaIncDecEmoxErr)
			if err != nil {
				fmt.Println("ERROR creating JSON data for qa  emox" + err.Error())
			}
					return string(resp)

		}
			return QAEmoxUpdateResult
		}
	emoxSameResult:= map[string]interface{}{"code": util.CODE_QA503, "message": "Success", "result": "Selected SameEMox"}
			resp, err := json.Marshal(emoxSameResult)
			if err != nil {
				fmt.Println("ERROR creating JSON data for qa  emox" + err.Error())
			}
					return string(resp)
}

//Validate post emox data form data
func validateAddQaEmox(emox QAEmox) (isValid bool, errorMessage string) {
	vErrors := []string{
		"Emox text not found",
	}
	//Check story emox text
	if emox.EmoxText == "" { //checkIfBlank
		return false, vErrors[0]
	}
	//Valid form data
	return true, ""
}

func EmoxExist(valueEmox QAEmox,c *iris.Context, mongoSession *mgo.Session) (checkResult bool,preEmoxText string) {
	fmt.Println("Check qa emox")

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	
	
	collection := sessionCopy.DB(util.DATABASENAME).C("qaemoxs")
	authorInfo := getUserJWT(c,"")
	 QaEmox	:=	QAEmox{}
	fmt.Println("questionid",valueEmox.QuestionID)
	fmt.Println("answerid",valueEmox.AnswerID)
	fmt.Println("userid",authorInfo.UserID)
	fmt.Println("emox text",valueEmox.EmoxText)

	//check emox 
	err := collection.Find(bson.M{"questionid":valueEmox.QuestionID,"answerid":valueEmox.AnswerID,"userid":authorInfo.UserID}).One(&QaEmox)

	if err != nil {
		fmt.Println("ERROR During check emoxes for qa" + err.Error())
		return false,""
	}
	
	fmt.Println("pre emox text",QaEmox.EmoxText)
	fmt.Println("current emox text",valueEmox.EmoxText)
			
	return true,QaEmox.EmoxText
	
}

func QAEmoxUpdate(valueEmox QAEmox,c *iris.Context, mongoSession *mgo.Session)(QAEmoxUpdateResult string){
	fmt.Println("Update qa emox")

	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	
	collection := sessionCopy.DB(util.DATABASENAME).C("qaemoxs")
	authorInfo := getUserJWT(c,"")
//	 QaEmox	:=	QAEmox{}
	fmt.Println("questionid",valueEmox.QuestionID)
	fmt.Println("answerid",valueEmox.AnswerID)
	fmt.Println("userid",authorInfo.UserID)
	fmt.Println("emox text",valueEmox.EmoxText)

	opMatch := bson.M{"questionid":valueEmox.QuestionID,"answerid":valueEmox.AnswerID,"userid":authorInfo.UserID}
	opUpdate := bson.M{"$set":bson.M{"emoxtext":valueEmox.EmoxText}}
	err := collection.Update(opMatch,opUpdate)
	if err != nil{
		fmt.Println("error updating emox in check emox",err.Error())
	}
	qaUpdateEmoxResult := map[string]interface{}{"code": util.CODE_QA503, "message": "Success", "result": "QA Emox Added Successfully"}
	resp, err := json.Marshal(qaUpdateEmoxResult)
	if err != nil {
		fmt.Println("ERROR creating JSON data for qa update emox" + err.Error())
	}

	return string(resp)
}

func IncrementEmoxCount(emox QAEmox,c *iris.Context, mongoSession *mgo.Session)(IncrementEmoxCountResult bool){
	fmt.Println("Increment qa emox")
	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	opMatchQuery := bson.M{"$and": []interface{}{bson.M{"_id": emox.QuestionID}, bson.M{"answer.answerid": emox.AnswerID}}}
	opUpdate := bson.M{"$inc": bson.M{"answer.$.emoxscale." + emox.EmoxText: 1, "answer.$.emoxcount": 1}}
	err := collection.Update(opMatchQuery, opUpdate)

	if err != nil {
		fmt.Println("Increase the emox count for the qa" + err.Error())
		return false
	}
	return true
}

func DecrementIncrementEmoxCount(preEmoxText string,emox QAEmox,c *iris.Context, mongoSession *mgo.Session)(IncrementEmoxCountResult bool){
	fmt.Println("Increment qa emox")
		fmt.Println("preemoxtext",preEmoxText)
	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(util.DATABASENAME).C("questionanswers")
	opMatchQuery := bson.M{"$and": []interface{}{bson.M{"_id": emox.QuestionID}, bson.M{"answer.answerid": emox.AnswerID}}}
	opUpdate := bson.M{"$inc": bson.M{"answer.$.emoxscale." + preEmoxText: -1,"answer.$.emoxscale." + emox.EmoxText: 1}}

	err := collection.Update(opMatchQuery, opUpdate)

	if err != nil {
		fmt.Println("Decrease the emox count for the qa" + err.Error())
		return false
	}
	return true
}

func AddAnswerEmoxCount(valueEmox QAEmox,c *iris.Context, mongoSession *mgo.Session) (AddAnswerEmoxCountResult string){
	fmt.Println("inside addanswerEmox")
	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	// Get access to storyemoxs collection
	collection := sessionCopy.DB(util.DATABASENAME).C("answeremoxcount")
	authorInfo := getUserJWT(c, "")
	CheckAnswerEmoxCountExist := CheckAnswerEmoxCountExist(valueEmox,c, mongoSession)
	fmt.Println("CheckAnswerEmoxCountExist",CheckAnswerEmoxCountExist)
	if CheckAnswerEmoxCountExist != true{
	//Add Emox
	answerEmoxCount := UserAnswerEmoxCount{	
			ID:				bson.NewObjectId(),		
			UserID:			authorInfo.UserID,		
			QuestionID:		valueEmox.QuestionID,	
			AnswerID:		valueEmox.AnswerID,	
			EmoxCount:		1,		
	}
		
	fmt.Println("answerEmoxCount",answerEmoxCount)
	
	//Add emox for user/story
	err := collection.Insert(&answerEmoxCount)

	if err != nil {
		fmt.Println("ERROR During add answeremoxecount for qa" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_QA503, "message": "Error", "result": "Error in adding emoxes to qa"}
		errResp, _ := json.Marshal(errRsponse)

		return string(errResp)
	}
	//QA Emox added successfully , respond success json to the client
	qaAddEmoxResult := map[string]interface{}{"code": util.CODE_QA503, "message": "Success", "result": "QA Answer Emox Count Added Successfully"}
	resp, err := json.Marshal(qaAddEmoxResult)
	if err != nil {
		fmt.Println("ERROR creating JSON data for answeremoxecount for qa" + err.Error())
		}
			return string(resp)
	}
	UpdateAnswerEmoxCountResult := UpdateAnswerEmoxCount(valueEmox,c , mongoSession) 
	
	return UpdateAnswerEmoxCountResult
}
func UpdateAnswerEmoxCount(valueEmox QAEmox,c *iris.Context, mongoSession *mgo.Session) (AddAnswerEmoxCountResult string){
	fmt.Println("answer Emox Update")

	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	fmt.Println("answer id",valueEmox.AnswerID)
	// Get access to storyemoxs collection
	collection := sessionCopy.DB(util.DATABASENAME).C("answeremoxcount")
	authorInfo := getUserJWT(c, "")
	
	opMatchQuery := bson.M{"$and": []interface{}{bson.M{"answerid": valueEmox.AnswerID}, bson.M{"userid": authorInfo.UserID}}}
	opUpdate := bson.M{"$inc": bson.M{"emoxcount": 1}}
	err := collection.Update(opMatchQuery, opUpdate)

	if err != nil {
		fmt.Println("ERROR During update answeremoxecount for qa" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_QA503, "message": "Error", "result": "Error in updating answeremoxes to qa"}
		errResp, _ := json.Marshal(errRsponse)

		return string(errResp)
	}
	//QA Emox added successfully , respond success json to the client
	qaAddEmoxResult := map[string]interface{}{"code": util.CODE_QA503, "message": "Success", "result": "QA Answer Emox Count update Successfully"}
	resp, err := json.Marshal(qaAddEmoxResult)
	if err != nil {
		fmt.Println("ERROR creating JSON data for answeremoxecount for qa" + err.Error())
	}
			return string(resp)

}

func CheckAnswerEmoxCountExist(valueEmox QAEmox,c *iris.Context, mongoSession *mgo.Session) (CheckAnswerEmoxCountExist bool){

	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	fmt.Println("answer id",valueEmox.AnswerID)
	answeremoxcount := UserAnswerEmoxCount{}
	// Get access to storyemoxs collection
	collection := sessionCopy.DB(util.DATABASENAME).C("answeremoxcount")
	authorInfo := getUserJWT(c, "")
	err := collection.Find(bson.M{"$and": []interface{}{bson.M{"answerid": valueEmox.AnswerID}, bson.M{"userid": authorInfo.UserID}}}).One(&answeremoxcount)

	if err != nil {
	//	fmt.Println("ERROR During update answeremoxecount for qa" + err.Error())
		return false
	}
	
			return true

}

func GetAnswerEmoxCount(c *iris.Context, mongoSession *mgo.Session) (GetAnswerEmoxCountResult string){
	fmt.Println("answer Emox Get")
//	var answerEmoxCount []UserAnswerEmoxCount
	answerEmoxCount := UserAnswerEmoxCount{}

	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	// Get access to storyemoxs collection
	collection := sessionCopy.DB(util.DATABASENAME).C("answeremoxcount")

opMatchQuery := bson.M{}
	if bson.IsObjectIdHex(c.Param("uid")) == true{
		uid := bson.ObjectIdHex(c.Param("uid"))
		fmt.Println("uid ",uid)
	opMatchQuery = bson.M{"$match":bson.M{"userid": uid}}
	fmt.Println(opMatchQuery)
	}else{
		authorInfo := getUserJWT(c, "")
		uid := authorInfo.UserID 
		fmt.Println("uid auth",uid)
	opMatchQuery = bson.M{"$match":bson.M{"userid": authorInfo.UserID}}
	}
	//opMatchQuery := bson.M{"$match":bson.M{"userid": authorInfo.UserID}}
	opProject1 := bson.M{ "$group": bson.M{"_id": "","emoxcount": bson.M{ "$sum": "$emoxcount"}}}
	opProject2 := bson.M{ "$project": bson.M{"_id": 1,"emoxcount": 1}}

	operations := []bson.M{opMatchQuery,opProject1,opProject2}
	pipe := collection.Pipe(operations)
	//err := pipe.All(&answerEmoxCount)
	err := pipe.One(&answerEmoxCount)

	//err := collection.Update(opMatchQuery, opUpdate)

	if err != nil {
		fmt.Println("ERROR During Get answeremoxecount for qa" + err.Error())

	}
	//QA Emox added successfully , respond success json to the client
	qaAddEmoxResult := map[string]interface{}{"code": util.CODE_QA503, "message": "Success", "result":answerEmoxCount.EmoxCount}
	resp, err := json.Marshal(qaAddEmoxResult)
	if err != nil {
		fmt.Println("ERROR creating JSON data for answeremoxecount for qa" + err.Error())
	}
			return string(resp)

}
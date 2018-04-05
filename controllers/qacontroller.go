package controllers

import (
//	"fmt"
	"log"
	"time"
	//"github.com/CHH/eventemitter"
	//"github.com/dgrijalva/jwt-go"
	"github.com/joybynature/jbnserverapp/models"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
)

const (
	MongoDBHosts = "mongodb1318.aws-us-east-1-portal.19.dblayer.com:11318"
	
//	MongoDBHosts = "ec2-52-66-145-155.ap-south-1.compute.amazonaws.com:27017"
	//Database = "jbnproddb"
	//Database = "jbndemodb"
	Database = "jbndevdb"
	UserName = "jbn"
	Password = "jbn12345$"


)

var (
	mongoSession *mgo.Session
	err          error
)

type Response struct {
	Text string `json:"text"`
}

func init() {
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: Database,
		Username: UserName,
		Password: Password,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)
}

func QAFeedHandler(c *iris.Context) {
	//Use User Token for personalization of QA feed
	//userToken := c.Get("jwt").(*jwt.Token)
	//	fmt.Println("USER TOKEN-", userToken)

	println("===START - GET QA FEED")

	qaFeedResult := models.GetQAFeed(c, mongoSession)
  

	//Send qa feed list
	c.Write(qaFeedResult)

	 
	println("===END - GET QA FEED")
}

func QAAanswerHandler(c *iris.Context) {
	println("===START - GET QA UNANSWERED QUESTIONS")

	qaAnswerResult := models.GetQAAnswer(c, mongoSession)
	
	//Send list of unanswered question
	c.Write(qaAnswerResult)
	
	println("===END - GET QA UNANSWERED QUESTIONS")
}

func QATrendingHandler(c *iris.Context) {
	println("===START - GET QA TRENDS QUESTIONS")

	qaTrendingResult := models.GetQATrends(c, mongoSession)

	//Send trending list
	c.Write(qaTrendingResult)

	println("===END - GET QA TRENDS QUESTIONS")
}

func QATagsHandler(c *iris.Context) {
	println("Start QA Tags")
	models.GenerateQATags(c.Param("text"), c, mongoSession)
	println("End Qa Tags")
}

func AddQuestionHandler(c *iris.Context) {

	println("===START - POST QUESTIONS")
	addQuestionResponse := models.AddQuestion(c, mongoSession)

	c.Write(addQuestionResponse)
	println("===END - POST QUESTIONS")
}

func QAListHandler(c *iris.Context) {
	println("===START - GET QA LIST")

	qaList := models.GetQuestionsList(c, mongoSession)
	//Send QA list
	c.Write(qaList)

	println("===END - GET QA LIST")
}

func QAListWebHandler(c *iris.Context) {
	println("===START - GET QA LIST")

	qaList := models.GetQuestionsListWeb(c, mongoSession)
	//Send QA list
	c.Write(qaList)

	println("===END - GET QA LIST")
}

func QADetailHandler(c *iris.Context) {
	println("===START - GET QA DETAILS")

	qaDetail := models.GetQADetail(c, mongoSession)
	//Send QA details
	c.Write(qaDetail)
	println("===END - GET QA DETAILS")
}

func RelatedQuestionHandler(c *iris.Context) {
	println("===START - GET RELATED QA DETAILS")
	qaDetail := models.GetRelatedQuestions(c, mongoSession)

	//Send QA related question
	c.Write(qaDetail)
	println("===END - GET RELATED QA DETAILS")

}

func DeleteQuestionHandler(c *iris.Context) {
	println("===Start - Delete Question from QA")

	deleteQuestionResult := models.DeleteQuestion(c, mongoSession)
	c.Write(deleteQuestionResult)
	println("===End - Delete Question from QA")

}

func SearchQAHandler(c *iris.Context) {
	println("===START - GET SEARCH QA ")
	searchTextResult := models.SearchQA(c, mongoSession)
	c.Write(searchTextResult)
	println("===End - GET SEARCH QA ")
}

func AddAnswerHandler(c *iris.Context) {
	println("===Start - POST ADD ANSWER QA ")

	AddAnswerToQuestionResult := models.AddAnswerToQuestion(c, mongoSession)
	c.Write(AddAnswerToQuestionResult)
	println("===End - POST ADD ANSWER QA ")
}

func GetAllAnswerCommentsHandler(c *iris.Context) {
	println("===Start - GET All Answer Comments QA ")

	getAllAnswerCommentsResult := models.GetAllAnswerComments(c, mongoSession)
	c.Write(getAllAnswerCommentsResult)
	println("===End - GET All Answer Comments QA ")

}

func MyQuestionsHandler(c *iris.Context) {
	println("===Start - GET My Questions QA ")

	getMyQuestion := models.GetMyQuestions(c, mongoSession)
	c.Write(getMyQuestion)
	println("===End - GET My Questions QA ")

}

func MyAnswersHandler(c *iris.Context) {

	println("===Start - GET My Answer QA ")
	getMyAnswerResult := models.GetMyAnswers(c, mongoSession)
	c.Write(getMyAnswerResult)
	println("===Start - GET My Answer QA ")

}

func AddCommentHandler(c *iris.Context) {
	println("===START - ADD QA COMMENT")

	addQAComment := models.AddComment(c, mongoSession)
	//Send related story list
	c.Write(addQAComment)

	println("===END - ADD QA COMMENT")
}

func ReviseQuestionHandler(c *iris.Context) {
	println("===START - Revise QA ")

	reviseQuestionResult := models.ReviseQuestion(c, mongoSession)
	c.Write(reviseQuestionResult)
	println("===End Revise QA ")

}

func DeleteAnswerHandler(c *iris.Context) {
	println("===START - Delete Answer QA ")

	deleteAnswerResult := models.DeleteAnswer(c, mongoSession)

	c.Write(deleteAnswerResult)
	println("===End Delete Answer QA ")

}

func AddEmoxHandler(c *iris.Context) {
	println("===START - POST EMOX FOR QA")
	addEmoxQA := models.AddEmoxQA(c, mongoSession)

	//Send post emox for qa
	c.Write(addEmoxQA)

	println("===END - POST EMOX FOR QA")
}

func QAListEmoxHandler(c *iris.Context) {

	println("===START - GET EMOX FOR QA")
	getEmoxQA := models.GetQAEmox(c, mongoSession)

	//Send post emox for qa
	c.Write(getEmoxQA)

	println("===END - GET EMOX FOR QA")
}

//TAGS
func tagHandler(c *iris.Context) {

}

func categoryHandler(c *iris.Context) {
	// Retrieve the parameter name
	tagParentId := c.Param("tagparentcode")
	tagLevelId := c.Param("taglevel")

	c.Write("TBD" + tagLevelId + tagParentId)
}

//QA WEB

func AddQuestionWebHandler(c *iris.Context) {

	println("===START - POST Add Question Web FOR QA")
	//postAddQuestion := models.AddQuestionWeb(c, mongoSession)
	models.AddQuestionWeb(c, mongoSession)
	//Send Add Question Web
	//	c.Write(postAddQuestion)

	println("===END - GET Add Question Web")

}

func AddAnswerWebHandler(c *iris.Context) {

	println("===START - POST Add Answer Web FOR QA")
	//postAddQuestion := models.AddQuestionWeb(c, mongoSession)
	addAnswerToQuestionWebResult := models.AddAnswerToQuestionWeb(c, mongoSession)
	//Send Add Question Web
	c.Write(addAnswerToQuestionWebResult)

	println("===END - GET Add Answer Web")

}


func GetAnswerHandler(c *iris.Context){
	println("===Start Get answer By Id")
	getanswerResult := models.GetAnswer(c,mongoSession)
	c.Write(getanswerResult)
		println("===Start Get answer By Id")
}
func GetTimeFilterQAHandler(c *iris.Context){
		println("===Start get Time Filter QA Result")
	getTimeFilterQAResult := models.GetTimeFilterQA(c,mongoSession)
	c.Write(getTimeFilterQAResult)
		println("===Start get Time Filter QA Result")
}

func GetCheckEmoxExistHandler(c *iris.Context){
	println("===Start Check Emox Exist")
	CheckEmoxExistResult := models.GetCheckEmoxExist(c,mongoSession)
	c.Write(CheckEmoxExistResult)
	println("===Start Check Emox Exist")
}
func GetAnswerEmoxCountHandler(c *iris.Context){
	println("===Start get totlat qa Emox count user")
	GetAnswerEmoxCountResult := models.GetAnswerEmoxCount(c,mongoSession)
	c.Write(GetAnswerEmoxCountResult)
	println("===Start get totlat qa Emox count user")
}

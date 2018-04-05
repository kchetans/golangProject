package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/joybynature/jbnserverapp/util"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Paragraph struct {
	ParagraphText  string `json:"text" bson:"text"`
	ParagraphImage string `json:"image" bson:"image"`
}

type StoryEmox struct {
	StoryID      bson.ObjectId `json:"storyid" bson:"storyid"`
	EmoxText     string        `json:"emoxtext" bson:"emoxtext"`
	ByUser       bson.ObjectId `json:"userid" bson:"userid"`
	EmoxDateTime time.Time     `json:"emoxdatetime" bson:"emoxdatetime"`
}

//for handling refreshs
type StoryEmoxScale struct {
	StoryId   bson.ObjectId  `json:"_id" bson:"_id"`
	EmoxScale EmoactionScale `json:"emoxscale" bson:"emoxscale"`
}

type StoryComment struct {
	CommentID    bson.ObjectId `json:"_id" bson:"_id"`
	StoryID      bson.ObjectId `json:"storyid" bson:"storyid"`
	PostText     string        `json:"posttext" bson:"posttext"`
	Author       Author        `json:"author" bson:"author"`
	PostDateTime time.Time     `json:"postdatetime" bson:"postdatetime"`
}

type Story struct {
	StoryId         bson.ObjectId  `json:"_id" bson:"_id"`
	Author          Author         `json:"author" bson:"author"`
	Title           string         `json:"title" bson:"title"`
	Paragraphs      []Paragraph    `json:"paragraph" bson:"paragraph"`
	PostDateTime    time.Time      `json:"postdatetime" bson:"postdatetime"`
	Readtime        string         `json:"readtime" bson:"readtime"`
	MainImageUrl    string         `json:"mainimageurl" bson:"mainimageurl"`
	RelatedImageUrl string         `json:"relatedimageurl" bson:"relatedimageurl"`
	VideoUrl        string         `json:"videourl" bson:"videourl"`
	Status          string         `json:"status" bson:"status"`
	Tags            []string       `json:"tags" bson:"tags"`
	EmoxScale       EmoactionScale `json:"emoxscale" bson:"emoxscale"`
	CommentCount    int            `json:"commentcount" bson:"commentcount"`
}

type RelatedStoryResponse struct {
	StoryId         bson.ObjectId `json:"_id" bson:"_id"`
	Title           string        `json:"title" bson:"title"`
	RelatedImageUrl string        `json:"relatedimageurl" bson:"relatedimageurl"`
}

func GetListOfStories(c *iris.Context, mongoSession *mgo.Session) (storylist string) {
	tagsParam := c.Param("tags")
	storyTags := strings.Split(tagsParam, ",")

	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var story []Story

	collection := sessionCopy.DB(util.DATABASENAME).C("stories")
	opMatch := bson.M{}

	if tagsParam == "" || tagsParam == "\"\"" {
		opMatch = bson.M{"$match": bson.M{"status": "Active"}}
	} else { //if category has a story tag
		//Set the regular expression to find related stories
		storyTagsQueryFmt := make([]bson.RegEx, len(storyTags))
		for tg := range storyTags {
			storyTagsQueryFmt[tg] = bson.RegEx{"^" + storyTags[tg] + "$", "i"}
		}
		opMatch = bson.M{"$match": bson.M{"status": "Active", "tags": bson.M{"$in": storyTagsQueryFmt}}}
	}

	opProject1 := bson.M{
		"$project": bson.M{"_id": 1, "title": 1, "paragraph": 1, "postdatetime": 1, "mainimageurl": 1, "videourl": 1, "tags": 1, "status": 1, "author": 1, "emoxscale": 1, "readtime": 1, "commentcount": 1},
	}

	opProject2 := bson.M{
		"$project": bson.M{"_id": 1, "title": 1, "paragraph": bson.M{"$slice": []interface{}{"$paragraph", 1}}, "postdatetime": 1, "mainimageurl": 1, "videourl": 1, "tags": 1, "status": 1, "author": 1, "emoxscale": 1, "readtime": 1, "commentcount": 1},
	}
	opSort := bson.M{"$sort": bson.M{"postdatetime": -1}}
	opLimit := bson.M{"$limit": 10} //can change later

	opMain := []bson.M{opMatch, opProject1, opProject2, opSort, opLimit}
	pipe := collection.Pipe(opMain)
	err := pipe.All(&story)

	if err != nil {
		fmt.Println("Error Getting Story List" + err.Error())
	}
	storyListResult := map[string]interface{}{"code": util.CODE_ST401, "message": "Success", "result": story}
	resp, err := json.Marshal(storyListResult)
	if err != nil {
		fmt.Println("Get Story List Result" + err.Error())
	}

	return string(resp)
}

func GetDetailStory(c *iris.Context, mongoSession *mgo.Session) (detailStory string) {
	fmt.Println("Inside Story Detail")
	storyId := c.Param("sid")
	fmt.Println(storyId)
	//c.Write(questionId)
	bsonStoryId := bson.ObjectIdHex(storyId)
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	var story []Story

	collection := sessionCopy.DB(util.DATABASENAME).C("stories")

	operation1 := bson.M{
		"$match": bson.M{
			"status": "Active",
			"_id":    bsonStoryId,
		},
	}

	operation2 := bson.M{
		"$project": bson.M{"_id": 1, "title": 1, "paragraph": 1, "postdatetime": 1, "mainimageurl": 1, "videourl": 1, "tags": 1, "status": 1, "author": 1, "emoxscale": 1,"commentcount":1, "readtime": 1},
	}

	operations := []bson.M{operation1, operation2}
	pipe := collection.Pipe(operations)
	err := pipe.All(&story)
	if err != nil {
		fmt.Println("Story Detail result err " + err.Error())
	}
	storyDetailResult := map[string]interface{}{"code": util.CODE_ST401, "message": "Success", "result": story}
	resp, err := json.Marshal(storyDetailResult)
	if err != nil {
		fmt.Println("STORY Detail REsult " + err.Error())
	}

	return string(resp)
}

func GetRelatedStories(c *iris.Context, mongoSession *mgo.Session) (relatedStories string) {
	//Get story search tag
	storyTags := strings.Split(c.Param("tags"), ",")

	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var relatedStory []RelatedStoryResponse

	collection := sessionCopy.DB(util.DATABASENAME).C("stories")

	//Set the regular expression to find related stories
	storyTagsQueryFmt := make([]bson.RegEx, len(storyTags))
	for tg := range storyTags {
		storyTagsQueryFmt[tg] = bson.RegEx{"^" + storyTags[tg] + "$", "i"}
	}
	fmt.Println("STORY REGEXP-", storyTagsQueryFmt)
	//Find the related stories based on related tags
	opMatch := bson.M{"$match": bson.M{"status": "Active", "tags": bson.M{"$in": storyTagsQueryFmt}}}
	opProject := bson.M{
		"$project": bson.M{"title": 1, "relatedimageurl": 1},
	}

	opMain := []bson.M{opMatch, opProject}
	pipe := collection.Pipe(opMain)
	err := pipe.All(&relatedStory)
	if err != nil {
		fmt.Println("Error in finding related stories" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_ST404, "message": "Error", "result": "Error in finding the related stories"}
		errResp, _ := json.Marshal(errRsponse)

		return string(errResp)
	}

	storyTagJsonResult := map[string]interface{}{"code": util.CODE_TAG001, "message": "Success", "result": relatedStory}
	resp, err := json.Marshal(storyTagJsonResult)
	if err != nil {
		fmt.Println("Error in creating JSON format for RELATED STORY" + err.Error())
	}

	return string(resp)
}

//Post comment on stories
func AddStoryComment(c *iris.Context, mongoSession *mgo.Session) (postCommentResult string) {
	//Retrieve all the form data
	//StoryID, PostText
	comment := StoryComment{}
	err := c.ReadForm(&comment)

	if err != nil {
		fmt.Println("Error when reading Add story comment form: " + err.Error())
	}

	//Validate Form data
	isValid, errMsg := validateAddStoryComment(comment)

	if !isValid {
		vStoryResult := map[string]interface{}{"code": util.CODE_ST410, "message": "Error", "result": errMsg}
		vResp, err := json.Marshal(vStoryResult)
		if err != nil {
			fmt.Println("Add story Comment Validation Error " + err.Error())
		}
		return string(vResp)
	}
	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(util.DATABASENAME).C("storycomments")

	//get the user id from user sessions ***REVISIT LATER
	authorInfo := getUserJWT(c, comment.Author.Name)

	//Add comment
	commentData := StoryComment{
		CommentID:    bson.NewObjectId(),
		StoryID:      comment.StoryID,
		PostText:     comment.PostText,
		Author:       authorInfo,
		PostDateTime: time.Now(),
	}
	err = collection.Insert(&commentData)

	if err != nil {
		fmt.Println("ERROR During Post Comment for stories" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_ST410, "message": "Error", "result": "Error in adding comment to stories"}
		errResp, _ := json.Marshal(errRsponse)

		return string(errResp)
	}
	//add comment count to main story
	collection = sessionCopy.DB(util.DATABASENAME).C("stories")
	opMatchQuery := bson.M{"_id": comment.StoryID}
	opUpdate := bson.M{"$inc": bson.M{"commentcount": 1}}
	err = collection.Update(opMatchQuery, opUpdate)

	if err != nil {
		fmt.Println("Add single comment to the story" + err.Error())
	}

	//Story Comment added successfully , respond success json to the client
	storyAddCommentResult := map[string]interface{}{"code": util.CODE_ST409, "message": "Success", "result": "Comment Added Successfully"}
	resp, err := json.Marshal(storyAddCommentResult)
	if err != nil {
		fmt.Println("ERROR creating JSON data for posting comments" + err.Error())
	}

	return string(resp)
}

//Validate Add comment form data
func validateAddStoryComment(comment StoryComment) (isValid bool, errorMessage string) {
	vErrors := []string{
		"Missing Story ID. Cannot be blank",
		"Comment cannot be blank",
		"Comment cannot be more than 1000 characters",
	}
	posttextstring := comment.PostText
	cCharacterCount := strings.Count(posttextstring, "")

	if comment.StoryID == "" { //Check if story is blank
		return false, vErrors[0]
	} else if comment.PostText == "" { //checkIsBlank
		return false, vErrors[1]
	} else if cCharacterCount >= 1000 { //checkWordCount
		return false, vErrors[2]
	}

	//Valid form data
	return true, ""
}

func GetStoryComments(c *iris.Context, mongoSession *mgo.Session) (storyComments string) {
	//Get all the comments for a story
	storyId := bson.ObjectIdHex(c.Param("sid"))

	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var storyCommentData []StoryComment

	collection := sessionCopy.DB(util.DATABASENAME).C("storycomments")

	//Find all the comments for a story id
	opMatch := bson.M{"$match": bson.M{"storyid": storyId}}
	opProject := bson.M{
		"$project": bson.M{"storyid": 1, "posttext": 1, "author": 1, "postdatetime": 1},
	}
	opSort := bson.M{"$sort": bson.M{"postdatetime": -1}}
	opLimit := bson.M{"$limit": 50} //can change later

	opMain := []bson.M{opMatch, opProject, opSort, opLimit}
	pipe := collection.Pipe(opMain)
	err := pipe.All(&storyCommentData)
	if err != nil {
		fmt.Println("Error in finding story comments" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_ST412, "message": "Error", "result": "Error in finding story comments"}
		errResp, _ := json.Marshal(errRsponse)

		return string(errResp)
	}

	storyCommentResult := map[string]interface{}{"code": util.CODE_ST411, "message": "Success", "result": storyCommentData}
	resp, err := json.Marshal(storyCommentResult)
	if err != nil {
		fmt.Println("Error in creating JSON format for story comments" + err.Error())
	}

	return string(resp)
}

//func AddEmoxStory(c *iris.Context, mongoSession *mgo.Session) (addEmox string) {
//	//Retrieve all the form data
//	emox := StoryEmox{}
//	err := c.ReadForm(&emox)

//	if err != nil {
//		fmt.Println("Error when reading Add story Emox form: " + err.Error())
//	}

//	//Validate Form data
//	isValid, errMsg := validateAddStoryEmox(emox)

//	if !isValid {
//		emoxValidationResult := map[string]interface{}{"code": util.CODE_ST413, "message": "Error", "result": errMsg}
//		vResp, err := json.Marshal(emoxValidationResult)
//		if err != nil {
//			fmt.Println("ERROR during Emox story validation" + err.Error())
//		}
//		return string(vResp)
//	}

//	//Get a mongo db connection
//	sessionCopy := mongoSession.Copy()
//	defer sessionCopy.Close()

//	// Get access to storyemoxs collection
//	collection := sessionCopy.DB(util.DATABASENAME).C("storyemoxs")
//	authorInfo := getUserJWT(c, "jbn")

//	//Add Emox
//	emoxData := StoryEmox{
//		StoryID:      emox.StoryID,
//		EmoxText:     emox.EmoxText,
//		ByUser:       authorInfo.UserID,
//		EmoxDateTime: time.Now(),
//	}
//	//Add emox for user/story
//	err = collection.Insert(&emoxData)

//	if err != nil {
//		fmt.Println("ERROR During add emoxes for stories" + err.Error())
//		errRsponse := map[string]interface{}{"code": util.CODE_ST414, "message": "Error", "result": "Error in adding emoxes to stories"}
//		errResp, _ := json.Marshal(errRsponse)

//		return string(errResp)
//	}
//	//Increment the appropriate emotions for the story
//	collection = sessionCopy.DB(util.DATABASENAME).C("stories")

//	opMatchQuery := bson.M{"_id": emox.StoryID}
//	opUpdate := bson.M{"$inc": bson.M{"emoxscale." + emox.EmoxText: 1}}
//	err = collection.Update(opMatchQuery, opUpdate)

//	if err != nil {
//		fmt.Println("Increase the emox count for the story" + err.Error())
//	}

//	//Story Emox added successfully , respond success json to the client
//	storyAddEmoxResult := map[string]interface{}{"code": util.CODE_ST413, "message": "Success", "result": "Story Emox Added Successfully"}
//	resp, err := json.Marshal(storyAddEmoxResult)
//	if err != nil {
//		fmt.Println("ERROR creating JSON data for story emox" + err.Error())
//	}

//	return string(resp)
//}

////Validate post emox data form data
//func validateAddStoryEmox(emox StoryEmox) (isValid bool, errorMessage string) {
//	vErrors := []string{
//		"Emox text not found",
//	}
//	//Check story emox text
//	if emox.EmoxText == "" { //checkIfBlank
//		return false, vErrors[0]
//	}
//	//Valid form data
//	return true, ""
//}

func GetStoryEmox(c *iris.Context, mongoSession *mgo.Session) (storyEmox string) {
	//Get all the comments for a story
	storyId := bson.ObjectIdHex(c.Param("sid"))
	fmt.Println("STORYID-", storyId)
	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var storyEmoxData []StoryEmoxScale

	collection := sessionCopy.DB(util.DATABASENAME).C("stories")

	//Find all the emox for a story id
	opMatch := bson.M{"$match": bson.M{"_id": storyId}}
	opProject := bson.M{
		"$project": bson.M{"storyid": "$_id", "emoxscale": 1},
	}

	opMain := []bson.M{opMatch, opProject}
	pipe := collection.Pipe(opMain)
	err := pipe.All(&storyEmoxData)
	fmt.Println("MAIN - ", opMain)
	if err != nil {
		fmt.Println("Error in finding story emoxs" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_ST416, "message": "Error", "result": "Error in finding emoxs for story"}
		errResp, _ := json.Marshal(errRsponse)

		return string(errResp)
	}

	storyEmoxResult := map[string]interface{}{"code": util.CODE_ST415, "message": "Success", "result": storyEmoxData}
	resp, err := json.Marshal(storyEmoxResult)
	if err != nil {
		fmt.Println("Error in creating JSON format for story Emox API" + err.Error())
	}

	return string(resp)
}

/**********************************************************************************************/

func AddEmoxStory(c *iris.Context, mongoSession *mgo.Session) (addEmox string) {
	//Retrieve all the form data
	emox := StoryEmox{}
	err := c.ReadForm(&emox)

	if err != nil {
		fmt.Println("Error when reading Add story Emox form: " + err.Error())
	}

	//Validate Form data
	isValid, errMsg := validateAddStoryEmox(emox)

	if !isValid {
		emoxValidationResult := map[string]interface{}{"code": util.CODE_ST413, "message": "Error", "result": errMsg}
		vResp, err := json.Marshal(emoxValidationResult)
		if err != nil {
			fmt.Println("ERROR during Emox story validation" + err.Error())
		}
		return string(vResp)
	}

	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	// Get access to storyemoxs collection
	collection := sessionCopy.DB(util.DATABASENAME).C("storyemoxs")
	authorInfo := getUserJWT(c, "jbn")
	checkResult ,preEmoxText	:= EmoxExistStory(emox,c, mongoSession)
	fmt.Println("check emox exist story",checkResult)
	fmt.Println("per emox text",preEmoxText)
	//Add Emox
	if checkResult != true{
		
	emoxData := StoryEmox{
		StoryID:      emox.StoryID,
		EmoxText:     emox.EmoxText,
		ByUser:       authorInfo.UserID,
		EmoxDateTime: time.Now(),
	}
	//Add emox for user/story
	err = collection.Insert(&emoxData)

	if err != nil {
		fmt.Println("ERROR During add emoxes for stories" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_ST414, "message": "Error", "result": "Error in adding emoxes to stories"}
		errResp, _ := json.Marshal(errRsponse)

		return string(errResp)
	}
	
	//Story Emox added successfully , respond success json to the client
	storyAddEmoxResult := map[string]interface{}{"code": util.CODE_ST413, "message": "Success", "result": "Story Emox Added Successfully"}
	resp, err := json.Marshal(storyAddEmoxResult)
	if err != nil {
		fmt.Println("ERROR creating JSON data for story emox" + err.Error())
	}
	//Increment emox count and emoxScale
		IncrementEmoxCountResult := IncrementEmoxCountStory(emox,c, mongoSession)
		//emox increment then result true else false
		if IncrementEmoxCountResult != true{
			qaIncEmoxErr:= map[string]interface{}{"code": util.CODE_QA503, "message": "fail", "result": "emox increment err"}
			resp, err := json.Marshal(qaIncEmoxErr)
			if err != nil {
				fmt.Println("ERROR creating JSON data for story  emox" + err.Error())
			}
					return string(resp)

		}


	return string(resp)
	}
	//if emox already present then update 
	if preEmoxText	!= emox.EmoxText{
			StoryEmoxUpdateResult := StoryEmoxUpdate(emox,c,mongoSession)
			IncrementDecrementEmoxCountResult  := DecrementIncrementEmoxCountStory(preEmoxText,emox,c, mongoSession)
			if IncrementDecrementEmoxCountResult != true{
			storyIncDecEmoxErr:= map[string]interface{}{"code": util.CODE_QA503, "message": "fail", "result": "emox increment err"}
			resp, err := json.Marshal(storyIncDecEmoxErr)
			if err != nil {
				fmt.Println("ERROR creating JSON data for story  emox" + err.Error())
			}
					return string(resp)

		}
			return StoryEmoxUpdateResult
		}
	emoxSameResult:= map[string]interface{}{"code": util.CODE_QA503, "message": "Success", "result": "Selected SameEMox"}
			resp, err := json.Marshal(emoxSameResult)
			if err != nil {
				fmt.Println("ERROR creating JSON data for story  emox" + err.Error())
			}
					return string(resp)
}

//Validate post emox data form data
func validateAddStoryEmox(emox StoryEmox) (isValid bool, errorMessage string) {
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

//func AddEmoxStory(c *iris.Context, mongoSession *mgo.Session) (addEmox string) {
//	//Retrieve all the form data
//	emox := QAEmox{}
//	err := c.ReadForm(&emox)
//	if err != nil {
//		fmt.Println("Error when reading Add QA Emox form: " + err.Error())
//	}

//	//Validate Form data
//	isValid, errMsg := validateAddQaEmox(emox)

//	if !isValid {
//		emoxValidationResult := map[string]interface{}{"code": util.CODE_QA503, "message": "Error", "result": errMsg}
//		vResp, err := json.Marshal(emoxValidationResult)
//		if err != nil {
//			fmt.Println("ERROR during Emox QA validation" + err.Error())
//		}
//		return string(vResp)
//	}

//	//Get a mongo db connection
//	sessionCopy := mongoSession.Copy()
//	defer sessionCopy.Close()

//	// Get access to storyemoxs collection
//	collection := sessionCopy.DB(util.DATABASENAME).C("qaemoxs")
//	authorInfo := getUserJWT(c, "")
//	checkResult ,preEmoxText	:= EmoxExist(emox,c, mongoSession)
//	fmt.Println("check exox exist",checkResult)
//	fmt.Println("check exox PreText",preEmoxText)
//	//check already emox or not if not then add
//	if checkResult != true{
//	//Add Emox
//	emoxData := QAEmox{
//		QuestionID:   emox.QuestionID,
//		AnswerID:     emox.AnswerID,
//		EmoxText:     emox.EmoxText,
//		ByUser:       authorInfo.UserID,
//		EmoxDateTime: time.Now(),
//	}
//	//Add emox for user/story
//	err = collection.Insert(&emoxData)

//	if err != nil {
//		fmt.Println("ERROR During add emoxes for qa" + err.Error())
//		errRsponse := map[string]interface{}{"code": util.CODE_QA503, "message": "Error", "result": "Error in adding emoxes to qa"}
//		errResp, _ := json.Marshal(errRsponse)

//		return string(errResp)
//	}
//	//QA Emox added successfully , respond success json to the client
//	qaAddEmoxResult := map[string]interface{}{"code": util.CODE_QA503, "message": "Success", "result": "QA Emox Added Successfully"}
//	resp, err := json.Marshal(qaAddEmoxResult)
//	if err != nil {
//		fmt.Println("ERROR creating JSON data for qa emox" + err.Error())
//	}
//	//Increment emox count and emoxScale
//		IncrementEmoxCountResult := IncrementEmoxCount(emox,c, mongoSession)
//		//emox increment then result true else false
//		if IncrementEmoxCountResult != true{
//			qaIncEmoxErr:= map[string]interface{}{"code": util.CODE_QA503, "message": "fail", "result": "emox increment err"}
//			resp, err := json.Marshal(qaIncEmoxErr)
//			if err != nil {
//				fmt.Println("ERROR creating JSON data for qa  emox" + err.Error())
//			}
//					return string(resp)

//		}
//		return string(resp)

//	}
//	//if emox already present then update 
//	if preEmoxText	!= emox.EmoxText{
//			QAEmoxUpdateResult := QAEmoxUpdate(emox,c,mongoSession)
//			IncrementDecrementEmoxCountResult  := DecrementIncrementEmoxCount(preEmoxText,emox,c, mongoSession)
//			if IncrementDecrementEmoxCountResult != true{
//			qaIncDecEmoxErr:= map[string]interface{}{"code": util.CODE_QA503, "message": "fail", "result": "emox increment err"}
//			resp, err := json.Marshal(qaIncDecEmoxErr)
//			if err != nil {
//				fmt.Println("ERROR creating JSON data for qa  emox" + err.Error())
//			}
//					return string(resp)

//		}
//			return QAEmoxUpdateResult
//		}
//	emoxSameResult:= map[string]interface{}{"code": util.CODE_QA503, "message": "Success", "result": "Selected SameEMox"}
//			resp, err := json.Marshal(emoxSameResult)
//			if err != nil {
//				fmt.Println("ERROR creating JSON data for qa  emox" + err.Error())
//			}
//					return string(resp)
//}

////Validate post emox data form data
//func validateAddQaEmox(emox QAEmox) (isValid bool, errorMessage string) {
//	vErrors := []string{
//		"Emox text not found",
//	}
//	//Check story emox text
//	if emox.EmoxText == "" { //checkIfBlank
//		return false, vErrors[0]
//	}
//	//Valid form data
//	return true, ""
//}

func EmoxExistStory(emox StoryEmox,c *iris.Context, mongoSession *mgo.Session) (checkResult bool,preEmoxText string) {
	fmt.Println("Check qa emox")

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	
	collection := sessionCopy.DB(util.DATABASENAME).C("storyemoxs")
	authorInfo := getUserJWT(c,"jbn")
	 StoryEmox	:=	StoryEmox{}
	fmt.Println("storyid",emox.StoryID)
	fmt.Println("emox",emox.EmoxText)
	fmt.Println("userid",authorInfo.UserID)

	//check emox 
	err := collection.Find(bson.M{"storyid":emox.StoryID,"userid":authorInfo.UserID}).One(&StoryEmox)

	if err != nil {
		fmt.Println("ERROR During check emoxes for story" + err.Error())
		return false,""
	}
	
	fmt.Println("pre emox text",StoryEmox.EmoxText)
	fmt.Println("current emox text",emox.EmoxText)
			
	return true,StoryEmox.EmoxText
	
}

func StoryEmoxUpdate(emox StoryEmox,c *iris.Context, mongoSession *mgo.Session)(QAEmoxUpdateResult string){
	fmt.Println("Update story emox")

	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	
	collection := sessionCopy.DB(util.DATABASENAME).C("storyemoxs")
	authorInfo := getUserJWT(c,"jbn")

	fmt.Println("storyid",emox.StoryID)
	fmt.Println("emox",emox.EmoxText)
	fmt.Println("userid",authorInfo.UserID)

	opMatch := bson.M{"storyid":emox.StoryID,"userid":authorInfo.UserID}
	opUpdate := bson.M{"$set":bson.M{"emoxtext":emox.EmoxText}}
	err := collection.Update(opMatch,opUpdate)
	if err != nil{
		fmt.Println("error updating emox in check emox",err.Error())
	}
	qaUpdateEmoxResult := map[string]interface{}{"code": util.CODE_QA503, "message": "Success", "result": "QA Emox Added Successfully"}
	resp, err := json.Marshal(qaUpdateEmoxResult)
	if err != nil {
		fmt.Println("ERROR creating JSON data for story update emox" + err.Error())
	}

	return string(resp)
}

func IncrementEmoxCountStory(emox StoryEmox,c *iris.Context, mongoSession *mgo.Session)(IncrementEmoxCountResult bool){
	fmt.Println("Increment story emox")
	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	authorInfo := getUserJWT(c,"jbn")
	fmt.Println("storyid",emox.StoryID)
	fmt.Println("emox",emox.EmoxText)
	fmt.Println("userid",authorInfo.UserID)
	collection := sessionCopy.DB(util.DATABASENAME).C("stories")
	opMatchQuery := bson.M{"_id": emox.StoryID}
	opUpdate := bson.M{"$inc": bson.M{"emoxscale." + emox.EmoxText: 1}}
	err := collection.Update(opMatchQuery, opUpdate)

	if err != nil {
		fmt.Println("Increase the emox count for the story" + err.Error())
		return false
	}
	return true
}

func DecrementIncrementEmoxCountStory(preEmoxText string,emox StoryEmox,c *iris.Context, mongoSession *mgo.Session)(IncrementEmoxCountResult bool){
	fmt.Println("Increment story emox")
		fmt.Println("preemoxtext",preEmoxText)
	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB(util.DATABASENAME).C("stories")
	opMatchQuery := bson.M{"_id": emox.StoryID}
	opUpdate := bson.M{"$inc": bson.M{"emoxscale." + preEmoxText: -1,"emoxscale." + emox.EmoxText: 1}}

	err := collection.Update(opMatchQuery, opUpdate)

	if err != nil {
		fmt.Println("Decrease the emox count for the story" + err.Error())
		return false
	}
	return true
}

func GetEmoxExistStory(c *iris.Context, mongoSession *mgo.Session) (GetEmoxExistStoryResult string) {
	fmt.Println("Check qa emox")
	StoryID := bson.ObjectIdHex(c.Param("sid")) 
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	
	collection := sessionCopy.DB(util.DATABASENAME).C("storyemoxs")
	authorInfo := getUserJWT(c,"jbn")
	 StoryEmox	:=	StoryEmox{}
	fmt.Println("storyid",StoryID)
	fmt.Println("userid",authorInfo.UserID)

	//check emox 
	OpMatch := bson.M{"$match": bson.M{"storyid":StoryID,"userid":authorInfo.UserID}}
	OpProject := bson.M{"$project": bson.M{"emoxtext": 1}}

	operations := []bson.M{OpMatch, OpProject}
	pipe := collection.Pipe(operations)
	err := pipe.One(&StoryEmox)
	//check emox 
	if err != nil {
		fmt.Println("ERROR During check emoxes for qa" + err.Error())
			
	}
	
	getStoryEmoxResult := map[string]interface{}{"code": util.CODE_QA503, "message": "Success", "result":StoryEmox}
	resp, err := json.Marshal(getStoryEmoxResult)
	if err != nil {
		fmt.Println("ERROR creating JSON data for story emox" + err.Error())
	}

	return string(resp)
	
}
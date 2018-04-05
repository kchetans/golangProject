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

type SpamAsset struct {
	AnswerID        bson.ObjectId `json:"answerid" bson:"answerid"`
	AnswerCommentID bson.ObjectId `json:"answercommentid" bson:"answercommentid"`
	StoryCommentID  bson.ObjectId `json:"storycommentid" bson:"storycommentid"`
	ReportText      string        `json:"reporttext" bson:"reporttext"`
	ByUser          bson.ObjectId `json:"byuser" bson:"byuser"`
	ByUserDateTime  time.Time     `json:"byuserdatetime" bson:"byuserdatetime"`
}

type SpamAction struct {
	SpamActionID   bson.ObjectId `json:"_id" bson:"_id"`
	SpamActionText string        `json:"spamactiontext" bson:"spamactiontext"`
	ByUser         bson.ObjectId `json:"byuser" bson:"byuser"`
	ByUserDateTime time.Time     `json:"byuserdatetime" bson:"byuserdatetime"`
}

type SpamReport struct {
	SpamID             bson.ObjectId `json:"_id" bson:"_id"`
	ReportedSpamAsset  []SpamAsset   `json:"spamasset" bson:"spamasset"`
	ReportedSpamAction []SpamAction  `json:"spamaction" bson:"spamaction"`
	Tags               []string      `json:"tags" bson:"tags"`
	Status             string        `json:"status" bson:"status"`
}

func GetSpam(c *iris.Context, mongoSession *mgo.Session) (spamReportList string) {
	tagsParam := c.Param("tags")
	spamTags := strings.Split(tagsParam, ",")

	//Get a mongo connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	var spamReport []SpamReport

	collection := sessionCopy.DB(util.DATABASENAME).C("spamreports")
	opMatch := bson.M{}

	if tagsParam == "" || tagsParam == "\"\"" {
		opMatch = bson.M{"$match": bson.M{"status": "Active"}}
	} else {
		//Set the regular expression to find related stories
		spamTagsQueryFmt := make([]bson.RegEx, len(spamTags))
		for stag := range spamTags {
			spamTagsQueryFmt[stag] = bson.RegEx{"^" + spamTags[stag] + "$", "i"}
		}
		opMatch = bson.M{"$match": bson.M{"status": "Active", "tags": bson.M{"$in": spamTagsQueryFmt}}}
	}

	opProject := bson.M{
		"$project": bson.M{"_id": 1, "spamasset": 1, "spamaction": 1, "tags": 1, "status": 1},
	}

	opMain := []bson.M{opMatch, opProject}
	pipe := collection.Pipe(opMain)
	err := pipe.All(&spamReport)

	if err != nil {
		fmt.Println("ERROR getting spam report list" + err.Error())
	}
	spamListResult := map[string]interface{}{"code": util.CODE_SP201, "message": "Success", "result": spamReport}
	resp, err := json.Marshal(spamListResult)
	if err != nil {
		fmt.Println("Get Spam report List" + err.Error())
	}

	return string(resp)
}

//Post comment on stories
func AddSpamReport(c *iris.Context, mongoSession *mgo.Session) (addSpamReportResult string) {

	//Retrieve all the form data
	//QuestionID, AnswerId, AnswerCommentID, StoryCommentID, ReportText
	spamAsset := SpamAsset{}
	err := c.ReadForm(&spamAsset)

	if err != nil {
		fmt.Println("Error when reading spam asset form: " + err.Error())
	}

	//Validate Form data
	isValid, errMsg := validateAddSpamReport(spamAsset)

	if !isValid {
		vSpamResult := map[string]interface{}{"code": util.CODE_SP101, "message": "Error", "result": errMsg}
		vResp, err := json.Marshal(vSpamResult)
		if err != nil {
			fmt.Println("Error Adding spam report validation" + err.Error())
		}
		return string(vResp)
	}
	//Get a mongo db connection
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(util.DATABASENAME).C("spamreport")

	//get the user id from user sessions ***REVISIT LATER
	authorInfo := getUserJWT(c, "jbn")

	//Add Spam Asset
	spamAssetData := []SpamAsset{
		{
			AnswerID:        spamAsset.AnswerID,
			AnswerCommentID: spamAsset.AnswerCommentID,
			StoryCommentID:  spamAsset.StoryCommentID,
			ReportText:      spamAsset.ReportText,
			ByUser:          authorInfo.UserID,
			ByUserDateTime:  time.Now(),
		},
	}

	//Add comment
	spamReportData := SpamReport{
		SpamID:             bson.NewObjectId(),
		ReportedSpamAsset:  spamAssetData,
		ReportedSpamAction: nil,
		Tags:               GetSpamTags(),
		Status:             "Active",
	}
	err = collection.Insert(&spamReportData)

	if err != nil {
		fmt.Println("ERROR During raising spam report" + err.Error())
		errRsponse := map[string]interface{}{"code": util.CODE_SP102, "message": "Error", "result": "Error raising spam report"}
		errResp, _ := json.Marshal(errRsponse)

		return string(errResp)
	}

	//Spam reportadded successfully , respond success json to the client
	spamReportResult := map[string]interface{}{"code": util.CODE_ST409, "message": "Success", "result": "Comment Added Successfully"}
	resp, err := json.Marshal(spamReportResult)
	if err != nil {
		fmt.Println("ERROR while creating JSON data for spam report" + err.Error())
	}

	return string(resp)
}

//Validate spam report form data
func validateAddSpamReport(spam SpamAsset) (isValid bool, errorMessage string) {
	vErrors := []string{
		"Spam report text cannot be blank",
	}
	spamText := spam.ReportText

	if spamText == "" {
		return false, vErrors[0]
	}

	//Valid form data
	return true, ""
}

func GetSpamTags() (tags []string) {
	spamTags := []string{"story"}

	return spamTags
}

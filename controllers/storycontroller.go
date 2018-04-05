package controllers

import (
	"github.com/joybynature/jbnserverapp/models"
	"github.com/kataras/iris"
)

func StoryListHandler(c *iris.Context) {
	println("===START - GET STORY LIST")

	listStories := models.GetListOfStories(c, mongoSession)
	//Send story list w/o tags
	c.Write(listStories)

	println("===END - GET STORY LIST")
}

func StoryDetailHandler(c *iris.Context) {
	println("===START - GET STORY DETAIL")

	detailStory := models.GetDetailStory(c, mongoSession)
	//Send story detail
	c.Write(detailStory)

	println("===END - GET STORY DETAIL")
}

func StoryRelatedHandler(c *iris.Context) {
	println("===START - GET RELATED STORIES")

	relatedStories := models.GetRelatedStories(c, mongoSession)
	//Send related story list
	c.Write(relatedStories)

	println("===END - GET RELATED STORIES")
}

func StoryPostCommentHandler(c *iris.Context) {
	println("===START - POST COMMENT STORIES")

	addCommentStories := models.AddStoryComment(c, mongoSession)
	//Send able to post comment for story
	c.Write(addCommentStories)

	println("===END - POST COMMENT STORIES")
}

func StoryListCommentHandler(c *iris.Context) {
	println("===START - GET Story List Comment Handler")

	listCommentStories := models.GetStoryComments(c, mongoSession)
	//Send able to get comments for a story
	c.Write(listCommentStories)

	println("===END - GET Story List Comment Handler")
}

func StoryAddEmoxHandler(c *iris.Context) {
	println("===START - POST Story Add Emox Handler")

	addEmoxStories := models.AddEmoxStory(c, mongoSession)
	//Send able to post comment for story
	c.Write(addEmoxStories)

	println("===END - POST Story Add Emox Handler")
}

func StoryListEmoxHandler(c *iris.Context) {
	println("===START - GET Story List Emox Handler")

	listEmoxStories := models.GetStoryEmox(c, mongoSession)
	//Send emoxs for a story
	c.Write(listEmoxStories)

	println("===END - Story List Emox Handler")
}
func GetEmoxExistStoryHandler(c *iris.Context) {
	println("===START - Get Emox Exist Story ")

	GetEmoxExistStoryResult := models.GetEmoxExistStory(c,mongoSession)
	//Send emoxs for a story
	c.Write(GetEmoxExistStoryResult)

	println("===END - Get Emox Exist Story Handler")
}

package controllers

import (
	"github.com/joybynature/jbnserverapp/models"
	"github.com/kataras/iris"
)

func QATagCategoryHandler(c *iris.Context) {
	println("===START - QA TAG CATEGORY")
	models.GetTagCategory(c, mongoSession)
	println("===END - QA TAG CATEGORY")
}

func ShopCategoryTagHandler(c *iris.Context) {
	println("===START - SHOP CATEGORY HANDLER")
	models.GetShopTag(c, mongoSession)
	println("===END - SHOP CATAGORY HANDLER")
}

func TagPoolHandler(c *iris.Context) {
	println("===START - TAG POOL HANDLER")
	tags := models.GetPoolTags(c, mongoSession)

	//Send the response to the client
	c.Write(tags)
	println("===END - TAG POOL HANDLER")
}

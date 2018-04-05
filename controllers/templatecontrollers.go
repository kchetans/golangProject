package controllers

import (
	_ "log"
	_ "time"

	"github.com/joybynature/jbnserverapp/models"
	"github.com/kataras/iris"
	_ "gopkg.in/mgo.v2"
)

func TermOfServicesHandler(c *iris.Context) {
	models.TermOfServices(c)
}

func PrivacyAndPolicyHandler(c *iris.Context) {
	models.PrivacyAndPolicy(c)
}

func SupportHandler(c *iris.Context) {
	models.Support(c)
}
func ReleaseHandler(c *iris.Context) {
	models.Release(c)
}
func TemplateAddQuestionHandler(c *iris.Context){
	models.AddQuestionTemplate(c)
}
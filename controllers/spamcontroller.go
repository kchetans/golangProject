package controllers

import (
	"github.com/joybynature/jbnserverapp/models"
	"github.com/kataras/iris"
)

func SpamListHandler(c *iris.Context) {
	println("===START - GET SPAM REPORTS")

	spamReportList := models.GetSpam(c, mongoSession)
	//Send spam report list
	c.Write(spamReportList)

	println("===END - GET SPAM REPORTS")
}

func SpamReportHandler(c *iris.Context) {
	println("===START - REPORT SPAM")

	addSpamReport := models.AddSpamReport(c, mongoSession)
	//Send spam report added successfull message
	c.Write(addSpamReport)

	println("===END - REPORT SPAM")
}

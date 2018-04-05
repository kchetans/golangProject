package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/iris-contrib/middleware/logger"
	"github.com/joybynature/jbnserverapp/config"
	"github.com/joybynature/jbnserverapp/routes"
	"github.com/joybynature/jbnserverapp/util"
	"github.com/kataras/iris"
		//"github.com/olebedev/emitter"

)

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Set Environment
	setJBNEnv()
}

func setJBNEnv() {
	//Set environment LOCAL, DEV, DEMO, STAGING, PRODUCTION
	JBNENV := os.Getenv("HOSTENV")
	fmt.Println("Environment : " + JBNENV)

	if JBNENV == "DEV" {
		util.JBNSetup = config.DevJBNSetup()
	} else if JBNENV == "DEMO" {
		util.JBNSetup = config.DemoJBNSetup()
	} else if JBNENV == "QA" {
		util.JBNSetup = config.QAJBNSetup()
	} else if JBNENV == "PROD" {
		util.JBNSetup = config.ProductionJBNSetup()
	} else { //JBNENV == "LOCAL"
		util.JBNSetup = config.LocalJBNSetup()
	}
}

func main() {
	//iris.Config.IsDevelopment = true // this will reload the templates on each request, defaults to false

	// set the global middlewares
	iris.Use(logger.New())

	//Set Routes
    iris.Set(iris.OptionMaxRequestBodySize(5 << 20))
	
	routes.RegisterQARoutes()
	routes.RegisterTagRoutes()
	routes.RegisterAdminRoutes()
	routes.RegisterStoryRoutes()
	routes.RegisterShopifyShopRoutes()
	routes.RegisterUserRoutes()
	routes.RegisterTemplatesRoutes()

	fmt.Println(":" + util.PORT)
	//iris.Listen("127.0.0.1:" + util.PORT)
	iris.Listen("192.168.0.111:" + util.PORT)
	//iris.Listen("10.0.0.9:" + util.PORT)
}

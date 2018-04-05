package routes

import (
	_ "github.com/joybynature/jbnserverapp/controllers"
	_ "github.com/kataras/iris"
	// "github.com/kataras/iris-contrib/middleware/basicauth"
)

func RegisterAdminRoutes() {
	//iris.Get("/api/v1/qa/postqaweb", controllers.massUploadQAHandler)
}

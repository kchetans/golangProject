package routes

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/joybynature/jbnserverapp/controllers"
	"github.com/kataras/iris"
)

func RegisterTemplatesRoutes() {
	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("JoybynatureWellbeing"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	iris.Config.IsDevelopment = true // this will reload the templates on each request, defaults to false

	iris.Get("/api/v1/template/termofservices",  controllers.TermOfServicesHandler)
	iris.Get("/api/v1/template/privacyandpolicy", controllers.PrivacyAndPolicyHandler)
	iris.Get("/api/v1/template/support",  controllers.SupportHandler)
	iris.Get("/api/v1/template/release",  controllers.ReleaseHandler)
	iris.Get("/api/v1/template/AddQuestion", myJwtMiddleware.Serve, controllers.TemplateAddQuestionHandler)

}

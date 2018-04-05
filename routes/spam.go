package routes

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/joybynature/jbnserverapp/controllers"
	"github.com/kataras/iris"
)

func RegisterSpamRoutes() {
	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("JoybynatureWellbeing"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	iris.Get("/api/v1/spam/list", myJwtMiddleware.Serve, controllers.SpamListHandler)
	iris.Post("/api/v1/spam/reportspam", myJwtMiddleware.Serve, controllers.SpamReportHandler)
}

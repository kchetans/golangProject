package routes

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/joybynature/jbnserverapp/controllers"
	"github.com/kataras/iris"
)

func RegisterTagRoutes() {
	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("JoybynatureWellbeing"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	//iris.Get("/api/v1/qa/tag", controllers.tagHandler)
	iris.Get("/api/v1/qa/tag/:tagparentcode/:taglevel", myJwtMiddleware.Serve, controllers.QATagCategoryHandler)
	iris.Get("/api/v1/shop/tag/:tagparentname/:taglevel", myJwtMiddleware.Serve, controllers.ShopCategoryTagHandler)
	iris.Get("/api/v1/qa/tagpool/:tagname", myJwtMiddleware.Serve, controllers.TagPoolHandler)
}

package routes

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/joybynature/jbnserverapp/controllers"
	"github.com/kataras/iris"
)

func RegisterStoryRoutes() {
	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("JoybynatureWellbeing"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	iris.Get("/api/v1/story/list/:tags", myJwtMiddleware.Serve, controllers.StoryListHandler)
	iris.Get("/api/v1/story/detail/:sid", myJwtMiddleware.Serve, controllers.StoryDetailHandler)
	iris.Get("/api/v1/story/relatedlist/:tags", myJwtMiddleware.Serve, controllers.StoryRelatedHandler)
	iris.Post("/api/v1/story/addcomment", myJwtMiddleware.Serve, controllers.StoryPostCommentHandler)
	iris.Get("/api/v1/story/storycomments/:sid", myJwtMiddleware.Serve, controllers.StoryListCommentHandler)
	iris.Post("/api/v1/story/addstoryemox", myJwtMiddleware.Serve, controllers.StoryAddEmoxHandler)
	iris.Get("/api/v1/story/listemox/:sid", myJwtMiddleware.Serve, controllers.StoryListEmoxHandler)
	iris.Get("/api/v1/story/checkstoryemox/:sid",myJwtMiddleware.Serve,controllers.GetEmoxExistStoryHandler)
}

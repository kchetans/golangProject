package routes

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/joybynature/jbnserverapp/controllers"
	"github.com/kataras/iris"
)

func RegisterQARoutes() {
	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("JoybynatureWellbeing"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	//QA Landing-Feed, Unanswered, Trending
	iris.Get("/api/v1/qa/feed/:tags/:qid", myJwtMiddleware.Serve, controllers.QAFeedHandler)
	iris.Get("/api/v1/qa/answer/:tags/:qid", myJwtMiddleware.Serve, controllers.QAAanswerHandler)
	iris.Get("/api/v1/qa/trending/:tags/:pageno", myJwtMiddleware.Serve, controllers.QATrendingHandler)
	iris.Get("/api/v1/qa/filter/:tags/:time", myJwtMiddleware.Serve,controllers.GetTimeFilterQAHandler)
	//QA
	iris.Get("/api/v1/qa/qalist/:tags/:qid", myJwtMiddleware.Serve, controllers.QAListHandler)
	iris.Get("/api/v1/qa/qadetail/:qid", myJwtMiddleware.Serve, controllers.QADetailHandler)
	iris.Get("/api/v1/qa/search/:text", myJwtMiddleware.Serve, controllers.SearchQAHandler)

	//Question
	iris.Post("/api/v1/qa/addquestion", myJwtMiddleware.Serve, controllers.AddQuestionHandler)
	iris.Get("/api/v1/qa/relatedquestion/:tags", myJwtMiddleware.Serve, controllers.RelatedQuestionHandler)
	iris.Get("/api/v1/qa/deletequestion/:qid", myJwtMiddleware.Serve, controllers.DeleteQuestionHandler)
	iris.Post("/api/v1/qa/addrevision", myJwtMiddleware.Serve, controllers.ReviseQuestionHandler)

	//Answer
	iris.Post("/api/v1/qa/addanswer", myJwtMiddleware.Serve, controllers.AddAnswerHandler)
	iris.Get("/api/v1/qa/deleteanswer/:qid/:aid", myJwtMiddleware.Serve, controllers.DeleteAnswerHandler)
	iris.Get("/api/v1/qa/qaanswer/:qid/:aid",myJwtMiddleware.Serve,controllers.GetAnswerHandler)
	//Comments
	iris.Post("/api/v1/qa/addcomment", myJwtMiddleware.Serve, controllers.AddCommentHandler)
	iris.Get("/api/v1/qa/qacomment/:qid/:aid", myJwtMiddleware.Serve, controllers.GetAllAnswerCommentsHandler)

	//Emox
	iris.Post("/api/v1/qa/addqaemox", myJwtMiddleware.Serve, controllers.AddEmoxHandler)
	iris.Get("/api/v1/qa/listemox/:qid/:aid", myJwtMiddleware.Serve, controllers.QAListEmoxHandler)
	iris.Get("/api/v1/qa/checkemoxexist/:qid/:aid", myJwtMiddleware.Serve,controllers.GetCheckEmoxExistHandler)
	iris.Get("/api/v1/qa/totalemoxcountuser/:uid", myJwtMiddleware.Serve,controllers.GetAnswerEmoxCountHandler)
	
	//Tags
	iris.Get("/api/v1/qa/qatags/:text", myJwtMiddleware.Serve, controllers.QATagsHandler)

	//My-Profile
	iris.Get("/api/v1/qa/myquestions", myJwtMiddleware.Serve, controllers.MyQuestionsHandler)
	iris.Get("/api/v1/qa/myanswers", myJwtMiddleware.Serve, controllers.MyAnswersHandler)

	//QA Web ADMIN ONLY ACCESS

	iris.Post("/api/v1/qa/web/addquestion", controllers.AddQuestionWebHandler)
	iris.Post("/api/v1/qa/web/addanswer",controllers.AddAnswerWebHandler)
	iris.Get("/api/v1/qa/webqalist/:tags/:qid",  controllers.QAListWebHandler)
	//Image Upload

	//	iris.Post("/api/v1/qa/addimage", myJwtMiddleware.Serve, controllers.ImageUploadHandler)

}

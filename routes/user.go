package routes

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/joybynature/jbnserverapp/controllers"
	"github.com/kataras/iris"
)

func RegisterUserRoutes() {
	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("JoybynatureWellbeing"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	//USER
	iris.Post("/api/v1/user/login", controllers.LoginUserHandler)
	iris.Post("/api/v1/user/adduser", controllers.AddUserHandler)
//	iris.Get("/api/v1/user/getuserprofilepublic/:userid", myJwtMiddleware.Serve, controllers.GetUserProfilePublicHandler)
	iris.Get("/api/v1/user/getuserprofile/:userid", myJwtMiddleware.Serve, controllers.GetUserProfileHandler)
	iris.Get("/api/v1/user/verifyuser",myJwtMiddleware.Serve,controllers.VerifyUserHandler)	
	//FORGOT
	iris.Get("/api/v1/user/forgotpassword", myJwtMiddleware.Serve, controllers.ForgotUserPasswordHandler)
	iris.Post("/api/v1/user/editprofile",myJwtMiddleware.Serve,controllers.EditProfileHandler)

	//OTP
	iris.Post("/api/v1/user/verifyotp", myJwtMiddleware.Serve, controllers.VerifyUserOtpHandler)
	//User Web ADMIN ONLY ACCESS

	iris.Get("/api/v1/user/users",  controllers.ListAllUsersDataWeb)
	iris.Post("/api/v1/user/addremovefollower", myJwtMiddleware.Serve,controllers.AddRemoveFollowersHandler)
	iris.Post("/api/v1/user/resendotp", myJwtMiddleware.Serve,controllers.ReSendOtpHandler)
	iris.Get("/api/v1/user/ifollowers/:uid",myJwtMiddleware.Serve,controllers.GetWhoIFollowersHandler)
	iris.Get("/api/v1/user/myfollowers/:uid",myJwtMiddleware.Serve,controllers.GetWhoFollowersMeHandler)
	iris.Get("/api/v1/user/getfollowerexist/:fid",myJwtMiddleware.Serve,controllers.GetFollowerExistHandler)
	iris.Post("/api/v1/user/checkemail",myJwtMiddleware.Serve,controllers.CheckEmailGetMobileNumber)
	iris.Post("/api/v1/user/forgotpassword",myJwtMiddleware.Serve,controllers.ForgotPasswordUserHandler)
	iris.Get("/api/v1/user/subscriberid/:uid",myJwtMiddleware.Serve,controllers.FindSubscriberIDHandler)
	iris.Post("/api/v1/user/addlocation",myJwtMiddleware.Serve,controllers.AddLocationHandler)
	iris.Get("/api/v1/user/mysubscriberlist/:uid",myJwtMiddleware.Serve,controllers.GetSubscriberListMYFollowersHandler)

//	iris.Get("/api/v1/user/allsubscriberlist/",controllers.GetAllSubscriberListHandler)
	iris.Post("/api/v1/user/addallsubscriberlist/",myJwtMiddleware.Serve,controllers.InsertAllSubscriberListHandler)
	iris.Post("/api/v1/user/updatesubscriberid/",myJwtMiddleware.Serve,controllers.UpdateUserSubscriberIdHandler)
	iris.Post("/api/v1/user/sendnotification/",controllers.SendNotificationToAllHandler)
}
	

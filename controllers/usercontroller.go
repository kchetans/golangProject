package controllers

import (
	_ "log"
	_ "time"

	"github.com/joybynature/jbnserverapp/models"
	"github.com/kataras/iris"
	_ "gopkg.in/mgo.v2"
)

func LoginUserHandler(c *iris.Context) {
	println("===START - LOGIN USER")

	loginUserResult := models.LoginUser(c, mongoSession)

	//Send USER LOGGED IN MESSAGE
	c.Write(loginUserResult)

	println("===END - LOGIN USER")
}

func AddUserHandler(c *iris.Context) {
	println("===START - ADD USER")

	addUserResult := models.AddUser(c, mongoSession)

	//Send add user is successfull message
	c.Write(addUserResult)

	println("===END - ADD USER")
}
/*
func VerifyUserOtpHandler(c *iris.Context){
	models.VerifyUserOtp(c, mongoSession)
	
}*/

func ListAllUsersDataWeb(c *iris.Context) {
	println("Get List Of users")

	models.GetAllUsersDataWeb(c, mongoSession)
	println("Get List Of users")
	
}

func GetUserProfileHandler(c *iris.Context) {
	println("===START - GET USER  PROFILE")

	UserResult := models.GetUserProfile(c, mongoSession)

	//Send user profile message
	c.Write(UserResult)

	println("===END - GET PUBLIC PROFILE")
}

//func GetUserProfileHandler(c *iris.Context) {
//	println("===START - GET USER PROFILE")

//	privUserResult := models.GetUserProfilePrivate(c, mongoSession)

//	//Send user profile message
//	c.Write(privUserResult)

//	println("===END - GET USER PROFILE")
//}

func ForgotUserPasswordHandler(c *iris.Context) {
	println("===START - FORGOT USER PASSWORD")

	forgotPassUserResult := models.ForgotPasswordUser(c, mongoSession)

	//Send FORGOT PASSWORD
	c.Write(forgotPassUserResult)

	println("===END - FORGOT USER PASSWORD")
}

func VerifyUserOtpHandler(c *iris.Context) {
	println("===START - VERIFY USER OTP")

	verifyOTPResult := models.VerifyOTP(c, mongoSession)

	//Send verify otp successful result
	c.Write(verifyOTPResult)

	println("===END - VERIFY USER OTP")
}

//func AddUserFollowerHandler(c *iris.Context) {
//	println("===START -Add User Follower Handler")

//AddUserFollowerResult := models.AddUserFollower(c,mongoSession)
//	//Send verify otp successful result
//	c.Write(AddUserFollowerResult)

//	println("===END - Add User Follower Handler")
//}


func ReSendOtpHandler(c *iris.Context){
	println("===START -RESEND OTP Handler")
	ResultReSendOtp :=	models.ReSendOtp(c,mongoSession)
	c.Write(ResultReSendOtp)
	println("===END-RESEND OTP Handler")

}

func VerifyUserHandler(c *iris.Context){
	VerifyUserResult := models.VerifyUser(c,mongoSession)
	c.Write(VerifyUserResult)
}

func AddRemoveFollowersHandler(c *iris.Context){
AddFollowersResult := models.AddRemoveFollowers(c,mongoSession)
	c.Write(AddFollowersResult)
	}
	
func GetWhoIFollowersHandler(c *iris.Context){
	println("inside Get Who Followers Me")
	GetWhoIFollowersResult := models.GetWhoIFollowers(c,mongoSession)
	c.Write(GetWhoIFollowersResult)
	println("inside Get Who Followers Me")

}

func GetWhoFollowersMeHandler(c *iris.Context){
	println("inside Get Who Followers Me Handler")
	GetWhoFollowersMeResult := models.GetWhoFollowersMe(c,mongoSession)
	c.Write(GetWhoFollowersMeResult)
	println("inside Get Who Followers Me Handler")

}	

func GetFollowerExistHandler(c *iris.Context){
	println("inside Get Follower Exist Handler")
	GetFollowerExistResult := models.GetFollowerExist(c,mongoSession)
	c.Write(GetFollowerExistResult)
	println("inside Get Follower Exist Handler")

}		

func EditProfileHandler(c *iris.Context){
	println("inside EditProfileHandler")
	EditProfileResult := models.EditProfile(c,mongoSession)
	c.Write(EditProfileResult)
	println("inside Edit Profile Handler")	
}

func CheckEmailGetMobileNumber(c *iris.Context){
	phone ,CheckEmailGetMobileNumberResult:= models.CheckEmailGetMobileNumber(c,mongoSession)
	println("phone",phone)
	c.Write(CheckEmailGetMobileNumberResult)
}
func ForgotPasswordUserHandler(c *iris.Context){
	println("Inside Forgot Password")
	ForgotPasswordUserResult := models.ForgotPasswordUser(c,mongoSession)
	c.Write(ForgotPasswordUserResult)
	println("Inside Forgot Password")
}

func FindSubscriberIDHandler(c *iris.Context){
	println("Inside Find Subscriber ID Handler")
	FindSubscriberIDResult := models.FindSubscriberID(c,mongoSession)
	c.Write(FindSubscriberIDResult)
	println("Inside Find Subscriber ID Handler")
	
}

func AddLocationHandler(c *iris.Context){
	println("Inside Add Location Handler")
	AddLocationResult := models.AddLocation(c,mongoSession)
	c.Write(AddLocationResult)
	println("Inside Add Location Handler")
}
//GetUserListMYFollowers(c *iris.Context,mongoSession *mgo.Session)(GetUserListMYFollowers string)
func GetSubscriberListMYFollowersHandler(c *iris.Context){
	println("Inside Add Location Handler")
	GetSubscriberListMYFollowersResult := models.GetSubscriberListMYFollowers(c,mongoSession)
	c.Write(GetSubscriberListMYFollowersResult)
	println("Inside Add Location Handler")
}

//func GetAllSubscriberListHandler(c *iris.Context){
//	println("Inside Get All Subscriber List Handler")
//	GetAllSubscriberListResult := models.GetAllSubscriberList(c,mongoSession)
//	c.Write(GetAllSubscriberListResult)
//	println("Inside Get All Subscriber List Handler")
//}

func InsertAllSubscriberListHandler(c *iris.Context){
		println("InsertAllSubscriberListHandler")
	InsertAllSubscriberListResult := models.InsertAllSubscriberList(c,mongoSession)
		c.Write(InsertAllSubscriberListResult)

		println("InsertAllSubscriberListHandler")

}

func UpdateUserSubscriberIdHandler(c *iris.Context){
		println("UpdateUserSubscriberIdHandler")
	UpdateUserSubscriberIdResult := models.UpdateUserSubscriberId(c,mongoSession)
		c.Write(UpdateUserSubscriberIdResult)

		println("UpdateUserSubscriberIdHandler")
}

func SendNotificationToAllHandler(c *iris.Context){
	println("SendNotificationToAllHandler")
	models.SendNotificationToAll(c,mongoSession)
	

		println("SendNotificationToAllHandler")
}
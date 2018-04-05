package models

import (
	"encoding/json"
	"fmt"
	//	"strings"
	"strconv"
	"time"
	"strings"
	"github.com/joybynature/jbnserverapp/util"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ShopifyCollections struct {
	Id             bson.ObjectId `json:"_id" bson:"_id"`
	Name           string        `json:"name" bson:"name"`
	CollectionType string        `json:"type" bson:"type"`
	Component      []component   `json:"component" bson:"component"`
	ShortOrder     int           `json:"sorder" bson:"sorder"`
	Status         string        `json:"status" bson:"status"`
}

type component struct {
	Cid        int       `json:"cid" bson:"cid"`
	Product    []Product `json:"product" bson:"product"`
	Name       string    `json:"name" bson:"name"`
	ShortOrder int       `json:"sorder" bson:"sorder"`
	ImageUrl   string    `json:"imageurl" bson:"imageurl"`
	Status     string    `json:"status" bson:"status"`
}

type Product struct {
	Pid      int    `json:"pid" bson:"pid"`
	Name     string `json:"name" bson:"name"`
	ImageUrl string `json:"imageurl" bson:"imageurl"`
	Price    int    `json:"price" bson:"price"`
	OldPrice int    `json:"oldprice" bson:"oldprice"`
	Status   string `json:"status" bson:"status"`
}
type RatingReview struct {
	ReviewId      bson.ObjectId 	`json:"_id" bson:"_id"`
	ProductId     int64         `json:"productid" bson:"productid"`
	Sku			  string		`json:"sku" bson:"sku"`	
	ProductHandle string        `json:"producthandle" bson:"producthandle"`
	State         string        `json:"state" bson:"state"`
	Rating        int           `json:"rating" bson:"rating"`
	Title         string        `json:"title" bson:"title"`
	AuthorName    string        `json:"authorname" bson:"authorname"`
	//Email         string      `json:"email" bson:"email"`
	//Location      string      `json:"location" bson:"location"`
	Body          string        `json:"body" bson:"body"`
	Reply         string        `json:"reply" bson:"reply"`
	CreatedAt     time.Time     `json:"createdat" bson:"createdat"`
	RepliedAt     time.Time     `json:"repliedat" bson:"repliedat"`
	UserType      string        `json:"usertype" bson:"usertype"`
	AuthorID      bson.ObjectId		 `json:"authorid" bson:"authorid"`
}

type ListProductRating struct{
	//Id     int64        `json:"_id" bson:"_id"`
	ProductId     int64        `json:"productid" bson:"productid"`
	Rating        float64           `json:"rating" bson:"rating"`
}

//type ListProductRatings struct{
//	ProductId     int         `json:"productid" bson:"productid"`
//	Rating		  string	 	`json:"rating" bson:"rating"`
//}
type ListProductForm struct{
	ListProductID[]		string
}

type RatingReviewsForm struct{
ProductId     int64  
ProductHandle string     
Rating        int         
Title         string      
Body          string      
//Reply         string      
//CreatedAt     time.Time   
//RepliedAt     time.Time     
//UserType      string        
//AuthorID        string        


}

type Grocerypincodes struct {
	Pincodes []pincode `json:"pincodes" bson:"pincodes"`
}
type pincode struct {
	Pincode int `json:"pincode" bson:"pincode"`
}

type PincodeResult struct {
	Availability string `json:"availability" bson:"availability"`
}

func GetShopifyShopCollections(c *iris.Context, mongoSession *mgo.Session) {
	fmt.Println("GetShopifyShopCollections")
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
//	pageno, _ := strconv.Atoi(c.Param("pageno"))
//	limit := 10
//	skip := 10 * (pageno - 1)
//	fmt.Println("skip", skip)
//	fmt.Println("PAGE NO-", pageno)
	var shopifyshop []ShopifyCollections

	collection := sessionCopy.DB(util.DATABASENAME).C("shopifycollections")
	operation1 := bson.M{"$match": bson.M{"status": "Active"}}
	operation2 := bson.M{"$match": bson.M{"component.status": "Active"}}
	
	operation4 := bson.M{"$sort": bson.M{"component.sorder": 1}}
	operation3 := bson.M{"$sort": bson.M{"sorder": 1}}
//	operation5 := bson.M{"$skip":skip}
//	operation6 := bson.M{"$limit":limit}
	//operation4 := bson.M{"$sort":bson.M{"component.sorder":1,},}
	operations := []bson.M{operation1, operation2, operation4, operation3}
	pipe := collection.Pipe(operations)
	err := pipe.All(&shopifyshop)
	if err != nil {
		fmt.Println("shopify shop collection  err " + err.Error())
	}
	shopifyShopCollectionResult := map[string]interface{}{"code": util.CODE_SH601, "message": "Success", "result": shopifyshop}
	resp, err := json.Marshal(shopifyShopCollectionResult)
	if err != nil {
		fmt.Println("shopify shop collection REsult " + err.Error())
	}

	//c.JSON(iris.StatusOK, resp)
	c.Write(string(resp))

}

func GetShopGroceryPincode(c *iris.Context, mongoSession *mgo.Session) {
	fmt.Println("i am in get shop grocery pincode in model")
	Pincode := c.Param("pincode")
	Grocerypincode, _ := strconv.Atoi(Pincode)
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	var grocerypincode []Grocerypincodes
	collection := sessionCopy.DB(util.DATABASENAME).C("pincodes")
	operation1 := bson.M{"$unwind": "$pincodes"}
	operation2 := bson.M{"$match": bson.M{"pincodes.pincode": Grocerypincode}}
	operation3 := bson.M{"$project": bson.M{"pincodes.pincode": 1}}
	operation := []bson.M{operation1, operation2, operation3}
	pipe := collection.Pipe(operation)
	err := pipe.All(&grocerypincode)
	if err != nil {
		fmt.Println("Grocery shop pincode err" + err.Error())
	}
	if grocerypincode != nil {
		//pincodecheck := []byte(`[{"availability":"true"}]`)
		pincoderesult := PincodeResult{
			Availability: "true",
		}
		ShopGroceryPincodeResult := map[string]interface{}{"code": util.CODE_SH601, "message": "Success", "result": pincoderesult}
		resp, err := json.Marshal(ShopGroceryPincodeResult)
		if err != nil {
			fmt.Println("Grocery shop pincode REsult " + err.Error())
		}

		//c.JSON(iris.StatusOK, resp)
		c.Write(string(resp))
	} else {
		pincoderesult := PincodeResult{
			Availability: "false",
		}
		ShopGroceryPincodeResult := map[string]interface{}{"code": util.CODE_SH601, "message": "Success", "result": pincoderesult}
		resp, err := json.Marshal(ShopGroceryPincodeResult)
		if err != nil {
			fmt.Println("Grocery shop pincode REsult " + err.Error())
		}

		//c.JSON(iris.StatusOK, resp)
		c.Write(string(resp))
	}

}
func GetShopReviewRating(c *iris.Context, mongoSession *mgo.Session) (GetShopReviewRatingResult string) {

	ProductId := c.Param("productid")
	Productid, _ := strconv.Atoi(ProductId)
	fmt.Println("pid",Productid)	
	sessionCopy := mongoSession.Copy()

	defer sessionCopy.Close()
	var	ratingreview []RatingReview
	
	collection := sessionCopy.DB(util.DATABASENAME).C("ratingreviews")
	err := collection.Find(bson.M{"productid": Productid}).All(&ratingreview)

	if err != nil {
		fmt.Println("shop review err" + err.Error())
	}
	ShopReviewRatingResult := map[string]interface{}{"code": util.CODE_SH601, "message": "Success", "result": ratingreview}
	resp, err := json.Marshal(ShopReviewRatingResult)
	if err != nil {
		fmt.Println("Shop review REsult " + err.Error())
	}
		
	return string(resp)

}
//func CheckUserRatingReviewsExits(c *iris.Context,mongoSession *mgo.Session)(CheckUserRatingReviewsExitsResult string){
//	sessionCopy := mongoSession.Copy()
//	defer sessionCopy.Clone()
//	ProductId := c.ParamInt64("productid")
//	var ratingReview []RatingReview
//	authorInfo := getUserJWT(c,"")
//	fmt.Println("name",authorInfo.UserID)
//	//userid := bson.ObjectIdHex(authorInfo.UserID)
//	fmt.Println(util.DATABASENAME)
//	collections := sessionCopy.DB(util.DATABASENAME).C("ratingreviews")
//	err := collections.Find(bson.M{"authorid":authorInfo.UserID,"productid":ProductId}).All(&ratingReview)
//	if err != nil{
//		fmt.Println("error in check user rating review",err.Error())
//	}
//	fmt.Println(user)
//	if user != nil{
//	UserResult := map[string]interface{}{"code": util.CODE_OTP1101, "message": "success", "result": true}
//			resp, err := json.Marshal(UserResult)
//			if err != nil {
//				fmt.Println("error in check user rating review " + err.Error())
//			}
//			fmt.Println(resp)
//	return string(resp)
//	}
//	UserResult := map[string]interface{}{"code": util.CODE_OTP1101, "message": "success", "result": false}
//			resp, err := json.Marshal(UserResult)
//			if err != nil {
//				fmt.Println(" error in check user rating review " + err.Error())
//			}
//			fmt.Println(resp)
//	return string(resp)
//}
func PostRatingReviews(c *iris.Context,mongoSession *mgo.Session) (postRatingReviewsResult string){
	var addRatingReviews = RatingReviewsForm{}
	authorInfo := getUserJWT(c,"")
	fmt.Println("name",authorInfo.UserID)
	err := c.ReadForm(&addRatingReviews)
		if err != nil {
		fmt.Println("Error when reading Rating and reviews:" + err.Error())
	}
		//Validate Form data
	isValid, errMsg := validatePostRatingReviews(addRatingReviews)

	if !isValid {
		qaValidationResult := map[string]interface{}{"code": util.CODE_QA502, "message": "Error", "result": errMsg}
		validationResp, err := json.Marshal(qaValidationResult)
		if err != nil {
			fmt.Println("Add Review and Rating Validation Error " + err.Error())
		}
		//c.JSON(iris.StatusOK, validationResp)
		return string(validationResp)
	} else { //Valida Form data
		//Get a mongo db connection
		sessionCopy := mongoSession.Copy()
		defer sessionCopy.Close()

		// Get a collection to execute the query against.
		collection := sessionCopy.DB(util.DATABASENAME).C("ratingreviews")
		sDataInfo := &RatingReview{ReviewId: bson.NewObjectId(),ProductId: addRatingReviews.ProductId,ProductHandle:addRatingReviews.ProductHandle,State:"published",Rating:addRatingReviews.Rating,Title: addRatingReviews.Title,AuthorName: authorInfo.Name,Body: addRatingReviews.Body,CreatedAt:time.Now(),AuthorID: authorInfo.UserID}
		data, _ := json.MarshalIndent(sDataInfo, "", "    ")
	fmt.Println("hi1")	
		fmt.Println("DATA -" + string(data))
		err = collection.Insert(sDataInfo)
		if err != nil {
			fmt.Println("Adding ratingreview err " + err.Error())
			addReviewRatingShopResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Fail", "result": err.Error()}
			resp, err := json.Marshal(addReviewRatingShopResult)
			if err != nil {
				fmt.Println("Add ratingreview REsult err " + err.Error())
			}
			return string(resp)
		}
	addReviewRatingShopResult := map[string]interface{}{"code": util.CODE_QA501, "message": "Success", "result": "Rating Review Added Successfully"}
		resp, err := json.Marshal(addReviewRatingShopResult)
		if err != nil {
			fmt.Println("Add ratingreview  REsult " + err.Error())
		}

		//c.JSON(iris.StatusOK, resp)
		return string(resp)
	
}
}


 func validatePostRatingReviews(addRatingReviews RatingReviewsForm) (isValid bool, errorMessage string){
	
	vErrors := []string{

		"Review Title cannot be more than 150 characters",
		"Review Body cannot be more than 500 characters",
	}
	titleText := addRatingReviews.Title
	bodyText := addRatingReviews.Body
	rCharacterCount1 := strings.Count(titleText, "")
	fmt.Println("Review Title character", rCharacterCount1)
	 if rCharacterCount1 >= 150 { //checkWordCount
		return false, vErrors[0]
	}
	rCharacterCount2 := strings.Count(bodyText, "")
	fmt.Println("Review Body character", rCharacterCount2)
	 if rCharacterCount2 >= 500 { //checkWordCount
		return false, vErrors[1]
	}
	//Valid form data
	return true, ""
}



func ProductRating(c *iris.Context,mongoSession *mgo.Session)(ProductRatingResult string){

	ProductId := c.Param("productid")
	Productid, _ := strconv.Atoi(ProductId)
	fmt.Println("pid",Productid)	
	var	ratingreview []ListProductRating


	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()	
	collection := sessionCopy.DB(util.DATABASENAME).C("ratingreviews")

	opMatch := bson.M{
	"$match":bson.M{"productid":Productid},
	}

	opProject1 := bson.M{
		"$group":bson.M{"_id": "$productid","rating":bson.M{"$avg":"$rating"}},
	}
	opProject2 := bson.M{
		"$project":bson.M{"productid":"$_id","rating":1},
	}
//db.ratingreviews.aggregate({$match:{productid:{$in:[ 135960907, 7079938183, 3678638983]}}},{"$group":{"_id": "$productid","avgQuantity":{"$avg":"$rating"}}}).pretty()


	opMain := []bson.M{opMatch, opProject1,opProject2}

	pipe := collection.Pipe(opMain)
	err := pipe.All(&ratingreview)


	if err != nil {
		fmt.Println("shop review err" + err.Error())
	}
	ShopReviewRatingResult := map[string]interface{}{"code": util.CODE_SH601, "message": "Success", "result": ratingreview}
	resp, err := json.Marshal(ShopReviewRatingResult)
	if err != nil {
		fmt.Println("Shop review REsult " + err.Error())
	}
	return string(resp)

}

func ListProductsRating(c *iris.Context,mongoSession *mgo.Session){
	ListProductData := ListProductForm{}
	err := c.ReadForm(&ListProductData)
	if err != nil{
		fmt.Println("List Product err",err.Error())
	}
//	  result := strings.Split(ListProductData, ",")
//	fmt.Println("length result",len(result))
//	fmt.Println("length result1",result)

	fmt.Println("length pid",len(ListProductData.ListProductID))
	fmt.Println("list pid",ListProductData.ListProductID)
    ProductIdList := []int{}

	//CONVERT PRODUCT ID TO INTEGER
	for	list := range ListProductData.ListProductID {
			pid := ListProductData.ListProductID[list]
			pnumber, _ := strconv.Atoi(pid)
			ProductIdList = append(ProductIdList, pnumber)

	}
	
fmt.Println("plist",ProductIdList)
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()
	var	ratingreview []ListProductRating
	collection := sessionCopy.DB(util.DATABASENAME).C("ratingreviews")

	opMatch := bson.M{
	"$match":bson.M{"productid":bson.M{"$in":ProductIdList}},
	}
//	opMatch := bson.M{
//	"$match":bson.M{"productid":bson.M{"$in":ListProductData.ListProductID}},
//	}
	opProject1 := bson.M{
		"$group":bson.M{"_id": "$productid","rating":bson.M{"$avg":"$rating"}},
	}
	opProject2 := bson.M{
		"$project":bson.M{"productid":"$_id","rating":1},
	}


	opMain := []bson.M{opMatch, opProject1,opProject2}
	pipe := collection.Pipe(opMain)
	err = pipe.All(&ratingreview)

				if err != nil {
					fmt.Println("shop review err" + err.Error())
				}
				
	listProductRatingResult := map[string]interface{}{"code": util.CODE_SH601, "message": "Success", "result": ratingreview}
		resp, err := json.Marshal(listProductRatingResult)
		if err != nil {
			fmt.Println("list Product Rating Result REsult " + err.Error())
		}
		c.Write( string(resp))			

}


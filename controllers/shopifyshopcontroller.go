package controllers

import(
	"github.com/joybynature/jbnserverapp/models"
	"github.com/kataras/iris"
)

func ShopifyShopCollectionHandler(c *iris.Context){
	println("inside shopify collection Handler")
	models.GetShopifyShopCollections(c, mongoSession)
	c.Write("shopify collection Handler")
}

func ShopGroceryPincodeHandler(c *iris.Context){
	println("inside shop pincode controller")
	models.GetShopGroceryPincode(c,mongoSession)
}

func ShopReviewRatingHandler(c *iris.Context){
	println("inside shop review rating")
	GetShopReviewRatingResult := models.GetShopReviewRating(c,mongoSession)
	c.Write(GetShopReviewRatingResult)
}

func PostRatingReviewsHandler(c *iris.Context){
	println("inside shop review rating")
	 postRatingReviewsResult := models.PostRatingReviews(c,mongoSession)
	c.Write(postRatingReviewsResult)
}

//func ShopListProductIDHandler(c *iris.Context){
//	println("inside shop Shop List ProductID")
//	models.ShopListProductID(c,mongoSession)
//	/*[ ShopListProductIDResult := models.ShopListProductID(c,mongoSession)
//	c.Write(ShopListProductIDResult[])*/
//}

func ListProductsRatingHandler(c *iris.Context){
	println("inside List Products Review")
	models.ListProductsRating(c,mongoSession)
}
//func ProductRating(c *iris.Context,mongoSession *mgo.Session)(ProductRatingResult string){
func ProductRatingHandler(c *iris.Context){
	println("inside  Products Review")
	ProductRatingResult := models.ProductRating(c,mongoSession)
	c.Write(string(ProductRatingResult))
}
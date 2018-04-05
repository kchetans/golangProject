package routes

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/joybynature/jbnserverapp/controllers"
	"github.com/kataras/iris"
)

func RegisterShopifyShopRoutes() {

	myJwtMiddleware := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("JoybynatureWellbeing"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	iris.Get("/api/v1/shop/collection", myJwtMiddleware.Serve, controllers.ShopifyShopCollectionHandler)
	iris.Get("/api/v1/shop/grocery/:pincode", myJwtMiddleware.Serve, controllers.ShopGroceryPincodeHandler)
	iris.Get("/api/v1/shop/ratingreviews/:productid",myJwtMiddleware.Serve,controllers.ShopReviewRatingHandler)
	iris.Post("/api/v1/shop/addratingreviews",myJwtMiddleware.Serve,controllers.PostRatingReviewsHandler)
	iris.Post("/api/v1/shop/listproductrating",myJwtMiddleware.Serve,controllers.ListProductsRatingHandler)
	iris.Get("/api/v1/shop/productrating/:productid",myJwtMiddleware.Serve,controllers.ProductRatingHandler)

}

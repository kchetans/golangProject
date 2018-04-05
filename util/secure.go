package util

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var (
	SALT = []byte("@#$%")
)

type JBNClaims struct {
	UserID      string `json:"userid"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Description string `json:"description"`
	Mobile      string `json:"mobile"`
	AvatarURL   string `json:"avatarurl"`
	EmailID     string `json:"emailid"`
	Scope       string `json:"scope"`
	jwt.StandardClaims
}

func GenerateToken(claims JBNClaims) (ss string) {
	mySigningKey := []byte("JoybynatureWellbeing")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jbntok, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)

	return jbntok
}

func ValidateToken(token string) (valid bool) {
	token1, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("JoybynatureWellbeing"), nil
	})

	if token1.Valid {
		fmt.Println("You look nice today")
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

	return true
}

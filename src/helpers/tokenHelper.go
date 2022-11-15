package helper

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

type SignedDetails struct {
	Username string
	Uid      string
	jwt.StandardClaims
}

var SecretKey = os.Getenv("TOKEN_SECRET_KEY")

func GenerateToken(username string, userId string) (signedToken string, signedRefresh string, err error) {
	claims := &SignedDetails{
		Username: username,
		Uid:      userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SecretKey))

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SecretKey))

	if err != nil {
		log.Fatal(err)
	}
	return token, refreshToken, err
}
func UpdateToken(Token string, RefreshToken string, userId string) {
	//var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	//var updateObj primitive.D
	return
}
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken, &SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
	)

	if err != nil {
		msg = "token invalid"
		//msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "token not match"
		//msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		//msg = err.Error()
		return
	}
	return claims, msg
}

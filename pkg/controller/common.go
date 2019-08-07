package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/restaurant/pkg/models"
	"time"
)

func generateToken (user *models.User) (string,error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["rank"] = user.Rank
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	t,err := token.SignedString([]byte("SecretKey"))

	return t,err
}

func GetTokenClaims (t string) (jwt.MapClaims, error) {


	//	var claims jwt.MapClaims
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error checking the signing method")
		}
		return []byte("SecretKey"), nil
	})
	return token.Claims.(jwt.MapClaims),err
}


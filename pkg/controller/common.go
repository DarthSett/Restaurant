package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/restaurant/pkg/models"
	"time"
)

//todo: replace panics with errors

func GenerateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["rank"] = user.Rank
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	t, err := token.SignedString([]byte("SecretKey"))

	return t, err
}

func GetTokenClaims(t string) (jwt.MapClaims, error) {

	//	var claims jwt.MapClaims
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error checking the signing method")
		}
		return []byte("SecretKey"), nil
	})
	return token.Claims.(jwt.MapClaims), err
}

func (u *UserController) Logout(c *gin.Context) {

	err := u.LogoutUser(c.GetHeader("token"))
	if err != nil {
		panic("error logging out" + err.Error())
	}
	c.Writer.Write([]byte("User Logged Out"))
}

package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/restaurant/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

func (u *UserController) Login (c *gin.Context) {
	//email := c.PostForm("email")
	//pass := c.PostForm("pass")
	var user *models.User
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting inputs: "+err.Error()))
	}

	if input["email"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no email sent"))
	}
	if input["pass"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no pass sent"))
	}
	if input["rank"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no rank sent"))
	}
	if input["rank"] == "0" {
		user, err = u.GetUser(input["email"], 0)
	} else if input["rank"] == "1" {
		user, err = u.GetAdmin(input["email"], 0)
	} else if input["rank"] == "2" {
		user, err = u.GetSuperAdmin(input["email"], 0)
	} else {
		err = fmt.Errorf("Invalid value for rank")
	}
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting user from db: %v", err))
	}
	println(user.Pass)
	println(input["pass"])
	err = bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(input["pass"]))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error matching passwords: "+err.Error()))
	} else {
		t, err := GenerateToken(user)
		println(t)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error creating tokens: "+err.Error()))
		}
		c.Writer.Header().Set("token", t)
		println(t)
		c.Writer.Write([]byte("User logged in. Token generated"))

	}
}

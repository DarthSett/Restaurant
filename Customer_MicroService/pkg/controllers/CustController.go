package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/restaurant/Customer_MicroService/pkg/database"
	"github.com/restaurant/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type CustController struct {
	database.Database
}

func NewCustController(db database.Database) *CustController {
	return &CustController{db}
}

func (u *CustController) CustMake(c *gin.Context) {
	cust := models.Customer{}
	err := c.BindJSON(&cust)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting inputs: "+err.Error()))
		return
	}
	pass,err := bcrypt.GenerateFromPassword([]byte(cust.Pass),4)
	if err!= nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("couldn't encrypt provided password"))
		return
	}
	cust.Pass = string(pass)
	println(cust.Pass)
	err = u.CreateCust(&cust)
	if err!= nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("couldn't save customer to db"))
		return
	}
	c.Writer.Write([]byte("Customer saved"))
}

func (u *CustController) CustLogin(c *gin.Context) {
	login := models.Credentials{}
	err := c.BindJSON(&login)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

		return
	}
	cust, err := u.GetCust(0,login.Email)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		//c.AbortWithStatusJSON(500,err)
		return
	}
	println("Email:" + cust.Email)
	err = bcrypt.CompareHashAndPassword([]byte(cust.Pass),[]byte(login.Pass))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	token,err := GenerateToken(cust)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Header("token",token)
	c.Writer.Write([]byte("Customer logged in. Token generated"))
}


func (u *CustController) CustGetResList(c *gin.Context) {
	id,name,err := u.RestList()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	o := make([]struct {
		Id   int
		Name string
	}, len(id))
	for i := range id {
		o[i].Id = id[i]
		o[i].Name = name[i]
	}
	c.JSON(200,o)
}


func (u *CustController) Logout(c *gin.Context) {

	t := c.GetHeader("token")
	err:=u.LogoutUser(t)
	if err != nil {
		panic("error logging out" + err.Error())
	}
	c.Writer.Write([]byte("User Logged Out"))
}

func GenerateToken(user *models.Customer) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	t, err := token.SignedString([]byte("SecretKey"))

	return t, err
}

func (u *CustController) TokenValidator(c *gin.Context) {
	t := c.GetHeader("token")
	if t == "" {
		panic("no token sent")
	}
	println(t)
	flag,err := u.Checktoken(t)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	if flag == false {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("token already deleted"))
		return
	}
	claims, err := GetTokenClaims(t)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	c.Set("claims", claims)

	//	claims := token.Claims.(jwt.MapClaims)

	if claims.Valid() == nil {
		c.Next()
	}
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
package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	database2 "github.com/restaurant/Res_Man_MicroService/pkg/database"
	"github.com/restaurant/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type SuperAdminController struct {
	database2.Database
}

func NewSuperAdminController(db database2.Database) *SuperAdminController {
	return &SuperAdminController{db}
}

func (u *SuperAdminController) SuperAdminmake(c *gin.Context) {
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting inputs: "+err.Error()))
	}
	//
	//name := c.PostForm("name")
	//pass := c.PostForm("pass")
	//email := c.PostForm("email")

	if input["email"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no email sent"))
	}
	if input["pass"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no pass sent"))
	}
	if input["name"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no name sent"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input["pass"]), 4)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error encrypting password: "+err.Error()))
	}
	SuperAdmin := models.NewUser(input["name"], input["email"], string(hash), 2, 0, 0, 0)
	println(SuperAdmin.Name, SuperAdmin.Pass)
	err = u.CreateSuperAdmin(SuperAdmin)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error while saving SuperAdmin to db: "+err.Error()))
	}
	c.Writer.Write([]byte("SuperAdmin Saved"))

}

func (u *SuperAdminController) SuperAdminget(c *gin.Context) {
	//email := c.PostForm("email")
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting inputs: "+err.Error()))
	}

	if input["id"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no id sent"))
	}

	id, err := strconv.Atoi(fmt.Sprintf("%v", input["id"]))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting id from input: "+err.Error()))
	}
	SuperAdmin, err := u.GetSuperAdmin("", id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting SuperAdmin from db: "+err.Error()))
	}
	SuperAdmin.Pass = ""
	c.JSON(200, SuperAdmin)
	//c.Writer.Write([]byte("Name: " + SuperAdmin.Name + "\nPass: " + SuperAdmin.Pass + "\nEmail: " + SuperAdmin.Email))
}

func (u *SuperAdminController) SuperAdminDel(c *gin.Context) {
	clms, _ := c.Get("claims")
	claims := clms.(jwt.MapClaims)
	id, err := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting id from input: "+err.Error()))
	}
	err = u.DeleteSuperAdmin(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting SuperAdmin from db: "+err.Error()))
	}
	i := fmt.Sprintf("%v",id)
	c.Writer.Write([]byte(i + " Deleted from db"))
}

//func (u *SuperAdminController) SuperAdminLogin(c *gin.Context) {
//	input := make(map[string]string)
//	err := c.BindJSON(&input)
//	if err != nil {
//		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting inputs: "+err.Error()))
//	}
//
//	if input["email"] == "" {
//		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no email sent"))
//	}
//	if input["pass"] == "" {
//		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no pass sent"))
//	}
//
//	sa, err := u.GetSuperAdmin(input["email"], 0)
//	if err != nil {
//		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting user from db: "+err.Error()))
//	}
//	println("id: ", sa.Id)
//	println(input["pass"])
//
//	err = bcrypt.CompareHashAndPassword([]byte(sa.Pass), []byte(input["pass"]))
//	if err != nil {
//		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error matching passwords: "+err.Error()))
//	} else {
//		t, err := GenerateToken(sa)
//		println(t)
//		if err != nil {
//			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error creating tokens: "+err.Error()))
//		}
//		c.Writer.Header().Set("token", t)
//		println(t)
//		c.Writer.Write([]byte("User logged in. Token generated"))
//	}
//}

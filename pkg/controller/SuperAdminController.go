package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/restaurant/pkg/database"
	"github.com/restaurant/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type SuperAdminController struct{
	database.Database

}

func NewSuperAdminController(db database.Database) *SuperAdminController {
	return &SuperAdminController{db}
}



func (u *SuperAdminController) SuperAdminmake(c *gin.Context){
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	//
	//name := c.PostForm("name")
	//pass := c.PostForm("pass")
	//email := c.PostForm("email")
	hash,err := bcrypt.GenerateFromPassword([]byte(input["pass"]),4)
	if err != nil {
		panic("Error encrypting password: " + err.Error())
	}
	SuperAdmin := models.NewUser(input["name"],input["email"],string(hash),2,"")
	println(SuperAdmin.Name, SuperAdmin.Pass)
	err = u.CreateSuperAdmin(SuperAdmin)
	if err != nil {
		panic("Error while saving SuperAdmin to db: " + err.Error())
	}
	c.Writer.Write([]byte("SuperAdmin Saved"))

}

func (u *SuperAdminController) SuperAdminget(c *gin.Context){
	//email := c.PostForm("email")
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	SuperAdmin,err := u.GetSuperAdmin(input["email"])
	if err != nil {
		panic("Error getting SuperAdmin from db: "+ err.Error())
	}
	SuperAdmin.Pass = ""
	c.JSON(200,SuperAdmin)
	//c.Writer.Write([]byte("Name: " + SuperAdmin.Name + "\nPass: " + SuperAdmin.Pass + "\nEmail: " + SuperAdmin.Email))
}

func (u *SuperAdminController) SuperAdminDel(c *gin.Context){
	claims,err := GetTokenClaims(c.GetHeader("token"))
	if err != nil {
		panic("Error getting claims from token: " + err.Error())
	}
	email := fmt.Sprintf("%v",claims["email"])
	err = u.DeleteSuperAdmin(email)
	if err != nil {
		panic("Error getting SuperAdmin from db: "+ err.Error())
	}
	c.Writer.Write([]byte(email + " Deleted from db"))
}

func (u *SuperAdminController) SuperAdminLogin(c *gin.Context)  {
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	sa,err := u.GetSuperAdmin(input["email"])
	if err != nil {
		panic("Error getting user from db: " + err.Error())
	}
	println(sa.Pass)
	println(input["pass"])

	err = bcrypt.CompareHashAndPassword([]byte(sa.Pass),[]byte(input["pass"]))
	if err != nil {
		panic("Error matching passwords: " + err.Error())
	} else {
		t ,err :=generateToken(sa)
		println(t)
		if err != nil {
			panic("Error creating tokens: " + err.Error())
		}
		c.Writer.Header().Set("token" , t)
		println(t)
		c.Writer.Write([]byte("User logged in. Token generated"))
	}
}

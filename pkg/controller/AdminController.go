package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/restaurant/pkg/database"
	"github.com/restaurant/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type AdminController struct{
	database.Database

}

func NewAdminController(db database.Database) *AdminController {
	return &AdminController{db}
}



func (u *AdminController) Adminmake(c *gin.Context){

	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}

	claims,err := GetTokenClaims(c.GetHeader("token"))
	if err != nil {
		panic("Error getting claims from token: " + err.Error())
	}
	adder := fmt.Sprintf("%v",claims["email"])


	hash,err := bcrypt.GenerateFromPassword([]byte(input["pass"]),4)
	if err != nil {
		panic("Error encrypting password: " + err.Error())
	}
	Admin := models.NewUser(input["name"],input["email"],string(hash),1,adder)
	err = u.CreateAdmin(Admin)
	if err != nil {
		panic("Error while saving Admin to db: " + err.Error())
	}




	c.Writer.Write([]byte("Admin Saved"))

}


func (u *AdminController) Adminget(c *gin.Context){
	//email := c.PostForm("email")
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	println(input["email"])
	Admin,err := u.GetAdmin(input["email"])
	if err != nil {
		panic("Error getting Admin from db: "+ err.Error())
	}
	Admin.Pass = ""
	c.JSON(200,Admin)

	//c.Writer.Write([]byte("Name: " + Admin.Name + "\nPass: " + Admin.Pass + "\nEmail: " + Admin.Email + "\n Rank: " + strconv.Itoa(Admin.Rank)))
}

func (u *AdminController) AdminDel(c *gin.Context){
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	err = u.DeleteAdmin(input["email"])
	if err != nil {
		panic("Error getting Admin from db: "+ err.Error())
	}
	c.Writer.Write([]byte(input["email"] + " Deleted from db"))
}

func (u *AdminController) AdminLogin(c *gin.Context)  {
	//email := c.PostForm("email")
	//pass := c.PostForm("pass")
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	admin,err := u.GetAdmin(input["email"])
	if err != nil {
		panic("Error getting user from db: " + err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Pass),[]byte(input["pass"]))
	if err != nil {
		panic("Error matching passwords: " + err.Error())
	} else {
		t ,err :=generateToken(admin)
		println(t)
		if err != nil {
			panic("Error creating tokens: " + err.Error())
		}
		c.Writer.Header().Set("token" , t)
		println(t)
		c.Writer.Write([]byte("User logged in. Token generated"))
	}

}

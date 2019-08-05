package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"

	//"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/restaurant/pkg/database"
	"github.com/restaurant/pkg/models"
	"golang.org/x/crypto/bcrypt"
	//"time"
)


type UserController struct{
	 database.Database

}

func NewUserController(db database.Database) *UserController {
	return &UserController{db}
}



func (u *UserController) Usermake(c *gin.Context){
	//name := c.PostForm("name")
	//pass := c.PostForm("pass")
	//email := c.PostForm("email")
	//println(email)
	input := decodeJson(c)
	claims,err := GetTokenClaims(c.GetHeader("token"))
	if err != nil {
		panic("Error getting claims from token: " + err.Error())
	}
	println(input["email"])
	adder := fmt.Sprintf("%v",claims["email"])


	hash,err := bcrypt.GenerateFromPassword([]byte(input["pass"]),4)
	if err != nil {
		panic("Error encrypting password: " + err.Error())
	}
	user := models.NewUser(input["name"],input["email"],string(hash),0,adder)
	err = u.CreateUser(user)
	if err != nil {
		panic("Error while saving user to db: " + err.Error())
	}
	c.Writer.Write([]byte("User Saved"))

}

func (u *UserController) Userget(c *gin.Context){
	//email := c.PostForm("email")
	input := decodeJson(c)
	user,err := u.GetUser(input["email"])
	if err != nil {
		panic("Error getting user from db: " + err.Error())
	}
	user.Pass = ""
	//c.Writer.Write([]byte("Name: " + user.Name + "\nPass: " + user.Pass + "\nEmail: " + user.Email))
	encodeJson(c,user)
}

func (u *UserController) UserDel(c *gin.Context){
	//Email := c.PostForm("email")
	input := decodeJson(c)
	Email := input["email"]
	claims,err := GetTokenClaims(c.GetHeader("token"))
	if err != nil {
		panic("Error getting claims from token: " + err.Error())
	}
	rank,err := strconv.Atoi(fmt.Sprintf("%v",claims["rank"]))
	if err != nil {
		panic("Error getting rank of admin: "+ err.Error())
	}
	if rank == 2 {
		err = u.DeleteUser(Email,"")
		if err != nil {
			panic("Error getting user from db: "+ err.Error())
		}
	}
	if rank == 1 {
		adder := fmt.Sprintf("%v",claims["email"])
		err = u.DeleteUser(Email,adder)
		if err != nil {
			panic("Error getting user from db: "+ err.Error())
		}
	}
	c.Writer.Write([]byte(Email + " Deleted from db"))
}

func (u *UserController) UserLogin(c *gin.Context)  {
	//email := c.PostForm("email")
	//pass := c.PostForm("pass")
	input := decodeJson(c)
	user,err := u.GetUser(input["email"])
	if err != nil {
		panic("Error getting user from db: " + err.Error())
	}
	println(user.Pass)
	println(input["pass"])

	err = bcrypt.CompareHashAndPassword([]byte(user.Pass),[]byte(input["pass"]))
	if err != nil {
		panic("Error matching passwords: " + err.Error())
	} else {
		t ,err :=generateToken(user)
		println(t)
		if err != nil {
			panic("Error creating tokens: " + err.Error())
		}
		c.Writer.Header().Set("token" , t)
		println(t)
		c.Writer.Write([]byte("User logged in. Token generated"))
	}

}




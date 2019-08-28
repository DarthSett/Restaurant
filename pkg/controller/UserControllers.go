package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	//"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/restaurant/pkg/database"
	"github.com/restaurant/pkg/models"
	"golang.org/x/crypto/bcrypt"
	//"time"
)

type UserController struct {
	database.Database
}

func NewUserController(db database.Database) *UserController {
	return &UserController{db}
}

func (u *UserController) Usermake(c *gin.Context) {
	//name := c.PostForm("name")
	//pass := c.PostForm("pass")
	//email := c.PostForm("email")
	//println(email)
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	clms, _ := c.Get("claims")
	claims := clms.(jwt.MapClaims)
	println(input["email"])
	adder, err := strconv.Atoi(fmt.Sprintf("%v", claims["id"]))
	if err != nil {
		panic("Error getting id of admin: " + err.Error())
	}
	role, err := strconv.Atoi(fmt.Sprintf("%v", claims["rank"]))
	if err != nil {
		panic("Error getting rank of admin: " + err.Error())
	}
	println(adder)

	if input["name"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("No name sent"))
	}
	if input["email"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("No email sent"))
	}
	if input["pass"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("No pass sent"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input["pass"]), 4)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error encrypting password: %v", err))
	}
	user := models.NewUser(input["name"], input["email"], string(hash), 0, adder, role, 0)
	err = u.CreateUser(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error saving user to db: %v", err))
	}
	c.Writer.Write([]byte("User Saved"))

}

func (u *UserController) Userget(c *gin.Context) {
	//email := c.PostForm("email")
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	if input["id"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("No name sent"))
	}
	id, err := strconv.Atoi(input["id"])
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting id from input: %v", err))
	}
	user, err := u.GetUser("", id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting user from db: "+err.Error()))
	}
	user.Pass = ""
	rids, names, err := u.GetUserRests(user.Email)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error getting the restaurants owned by user: "+err.Error()))
	}
	o := make([]struct {
		RID      int
		RestName string
	}, len(rids))
	for i := range rids {
		o[i].RID = rids[i]
		o[i].RestName = names[i]
	}
	//c.Writer.Write([]byte("Name: " + user.Name + "\nPass: " + user.Pass + "\nEmail: " + user.Email))
	c.JSON(200, gin.H{
		"user":              user,
		"Restaurants owned": o,
	})
	//c.JSON(200,user)
}

//func (u *UserController) UserRest (c *gin.Context) {
//	input := make(map[string]string)
//	c.BindJSON(&input)
//	rids,names,err := u.GetUserRests(input["Email"])
//	if err != nil {
//		panic("error getting the restaurants owned by user: " + err.Error())
//	}
//	o := make([]struct {
//		RID	int
//		RestName string
//	},len(rids))
//	n,_ := u.GetUser("",1)
//	for i,_ := range rids {
//		o[i].RID = rids[i]
//		o[i].RestName = names[i]
//	}
//
//	//c.Writer.Write([]byte("Name: " + user.Name + "\nPass: " + user.Pass + "\nEmail: " + user.Email))
//	c.JSON(200,o)
//}

func (u *UserController) UserDel(c *gin.Context) {
	//Email := c.PostForm("email")
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting inputs: "+err.Error()))
	}

	if input["id"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no id sent"))
	}

	id, err := strconv.Atoi(input["id"])
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting id from input: "+err.Error()))
	}

	err = u.DeleteUser(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting user from db: "+err.Error()))
	}
	c.Writer.Write([]byte(string(id) + " Deleted from db"))
}

//func (u *UserController) UserLogin(c *gin.Context) {
//	//email := c.PostForm("email")
//	//pass := c.PostForm("pass")
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
//	user, err := u.GetUser(input["email"], 0)
//	if err != nil {
//		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting user from db: %v", err))
//	}
//	println(user.Pass)
//	println(input["pass"])
//	err = bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(input["pass"]))
//	if err != nil {
//		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error matching passwords: "+err.Error()))
//	} else {
//		t, err := GenerateToken(user)
//		println(t)
//		if err != nil {
//			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error creating tokens: "+err.Error()))
//		}
//		c.Writer.Header().Set("token", t)
//		println(t)
//		c.Writer.Write([]byte("User logged in. Token generated"))
//
//	}
//
//}

func (u *UserController) UserUpdate(c *gin.Context) {
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		panic("Error getting inputs: " + err.Error())
	}
	if input["id"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no id sent"))
	}
	if input["flag"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no flag sent"))
	}
	if input["update"] == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no update sent"))
	}
	id, err := strconv.Atoi(input["id"])
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting id from input: "+err.Error()))
	}
	flag, err := strconv.Atoi(input["flag"])
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error getting id from input: "+err.Error()))
	}
	update := input["update"]
	user, err := u.GetUser("", id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error getting user from db"+err.Error()))
	}
	clms, _ := c.Get("claims")
	claims := clms.(jwt.MapClaims)
	rank, err := strconv.Atoi(fmt.Sprintf("%v", claims["rank"]))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error getting rank from claims: "+err.Error()))
	}
	adder, err := strconv.Atoi(fmt.Sprintf("%v", claims["rank"]))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error getting id from claims: "+err.Error()))
	}
	if flag == 1 {
		upd, err := bcrypt.GenerateFromPassword([]byte(update), 5)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error generating hash from pass"+err.Error()))
		}
		update = string(upd)
	}
	if rank == 2 {
		err = u.UpdateUser(id, update, flag)
	} else {
		if user.Adder == adder && user.AdderRole == rank {
			err = u.UpdateUser(id, update, flag)
		} else {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("admin is not the adder of this user"))
		}
		err = u.UpdateUser(id, update, flag)
	}
	c.Writer.Write([]byte("User Updated"))
}

func (u *UserController) ListUser(c *gin.Context) {
	name, email, id, err := u.UserList()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("There was an error getting the list from db: "+err.Error()))
	}
	t := make([]struct {
		Name  string
		Email string
		Id    int
	}, len(name))
	//o := make(map[string]string)
	for i := range name {
		t[i].Name = name[i]
		t[i].Email = email[i]
		t[i].Id = id[i]
	}
	c.JSON(200, t)
}

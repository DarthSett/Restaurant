package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/restaurant/pkg/database"
	"github.com/restaurant/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
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
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error getting inputs: " + err.Error()))
	}


	if input["email"] == "" {c.AbortWithError(http.StatusBadRequest,fmt.Errorf("no email sent"))}
	if input["pass"] == "" {c.AbortWithError(http.StatusBadRequest,fmt.Errorf("no pass sent"))}
	if input["name"] == "" {c.AbortWithError(http.StatusBadRequest,fmt.Errorf("no name sent"))}


	clms,_ := c.Get("claims")
	claims := clms.(jwt.MapClaims)
	adder, err := strconv.Atoi(fmt.Sprintf("%v",claims["id"]))


	hash,err := bcrypt.GenerateFromPassword([]byte(input["pass"]),4)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error encrypting password: " + err.Error()))
	}
	Admin := models.NewUser(input["name"],input["email"],string(hash),1,adder,2,0)
	err = u.CreateAdmin(Admin)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error while saving Admin to db: " + err.Error()))
	}




	c.Writer.Write([]byte("Admin Saved"))

}


func (u *AdminController) Adminget(c *gin.Context){
	//email := c.PostForm("email")
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error getting inputs: " + err.Error()))
	}


	if input["id"] == "" {c.AbortWithError(http.StatusBadRequest,fmt.Errorf("no id sent"))}


	id,err := strconv.Atoi(input["id"])
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error getting id from input: " + err.Error()))
	}
	Admin,err := u.GetAdmin("",id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error getting Admin from db: "+ err.Error()))
	}
	Admin.Pass = ""
	c.JSON(200,Admin)

	//c.Writer.Write([]byte("Name: " + Admin.Name + "\nPass: " + Admin.Pass + "\nEmail: " + Admin.Email + "\n Rank: " + strconv.Itoa(Admin.Rank)))
}

func (u *AdminController) AdminDel(c *gin.Context){
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error getting inputs: " + err.Error()))
	}

	if input["id"] == "" {c.AbortWithError(http.StatusBadRequest,fmt.Errorf("no id sent"))}


	id,err := strconv.Atoi(fmt.Sprintf("%v",input["id"]))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error getting id from input: "+ err.Error()))
	}
	err = u.DeleteAdmin(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error deleting Admin from db: "+ err.Error()))
	}
	c.Writer.Write([]byte(input["id"] + " Deleted from db"))
}


func (u *AdminController) AdminLogin(c *gin.Context)  {
	//email := c.PostForm("email")
	//pass := c.PostForm("pass")
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error getting inputs: " + err.Error()))
	}


	if input["email"] == "" {c.AbortWithError(http.StatusBadRequest,fmt.Errorf("no email sent"))}
	if input["pass"] == "" {c.AbortWithError(http.StatusBadRequest,fmt.Errorf("no pass sent"))}


	admin,err := u.GetAdmin(input["email"],0)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error getting user from db: " + err.Error()))
	}
	println(input[admin.Pass])
	err = bcrypt.CompareHashAndPassword([]byte(admin.Pass),[]byte(input["pass"]))

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error matching passwords: " + err.Error()))
	} else {
		t ,err :=generateToken(admin)
		println(t)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error creating tokens: " + err.Error()))
		}
		c.Writer.Header().Set("token" , t)
		println(t)
		c.Writer.Write([]byte("User logged in. Token generated"))
	}

}

func (u* AdminController) AdminUpdate(c *gin.Context) {
	input := make(map[string]string)
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error getting inputs: " + err.Error()))
	}
	if input["id"] == "" {c.AbortWithError(http.StatusBadRequest,fmt.Errorf("no id sent"))}
	if input["flag"] == "" {c.AbortWithError(http.StatusBadRequest,fmt.Errorf("no flag sent"))}
	if input["update"] == "" {c.AbortWithError(http.StatusBadRequest,fmt.Errorf("no update sent"))}
	id,err := strconv.Atoi(input["id"])
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error getting id from input: " + err.Error()))
	}
	flag,err := strconv.Atoi(input["flag"])
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("Error getting id from input: " + err.Error()))
	}
	update:=input["update"]

	if flag == 1 {
		upd,err := bcrypt.GenerateFromPassword([]byte(update),5)
		if err != nil {c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("error generating hash from pass" + err.Error()))}
		update = string(upd)
	}
	err = u.UpdateAdmin(id,update,flag)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError,fmt.Errorf("error while updating admin in db: "+err.Error()))
	}
	c.Writer.Write([]byte("admin Updated"))
}
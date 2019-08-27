package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/restaurant/pkg/controller"
	"github.com/restaurant/pkg/database"
	"github.com/restaurant/pkg/middleware"
	"github.com/restaurant/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Router struct {
	database.Database
}

func NewRouter(db database.Database) *Router {
	return &Router{
		db,
	}
}


func (r *Router) Router() *gin.Engine {

	defaultRouter := gin.Default()

	defaultRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello World !!",
		})
	})

	defaultRouter.POST("/login", func(c *gin.Context) {
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
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no pass sent"))
			}
			if input["rank"] == "0" {
				user, err = r.GetUser(input["email"], 0)
			} else if input["rank"] == "1" {
				user, err = r.GetAdmin(input["email"], 0)
			} else {
				user, err = r.GetSuperAdmin(input["email"], 0)
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
				t, err := controller.GenerateToken(user)
				println(t)
				if err != nil {
					c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Error creating tokens: "+err.Error()))
				}
				c.Writer.Header().Set("token", t)
				println(t)
				c.Writer.Write([]byte("User logged in. Token generated"))

			}
	})

	userController := controller.NewUserController(r)
	adminController := controller.NewAdminController(r)
	superadminController := controller.NewSuperAdminController(r)
	restcontroller := controller.NewRestController(r)
	midController := middleware.NewMidController(r)

	userGroup := defaultRouter.Group("user")
	userGroup.POST("/create", midController.TokenValidator, middleware.AdminRankAuthenticator, userController.Usermake)
	userGroup.GET("/get", midController.TokenValidator, middleware.AdminRankAuthenticator, userController.Userget)
	userGroup.DELETE("/del", midController.TokenValidator, middleware.SuperAdminRankAuthenticator, userController.UserDel)
	//userGroup.POST("/login", userController.UserLogin)
	userGroup.GET("/list", midController.TokenValidator, middleware.AdminRankAuthenticator, userController.ListUser)
	userGroup.PUT("/update", midController.TokenValidator, middleware.AdminRankAuthenticator, userController.UserUpdate)
	//userGroup.GET("/rest",midController.TokenValidator,middleware.AdminRankAuthenticator,userController.UserRest)

	adminGroup := defaultRouter.Group("admin")
	adminGroup.POST("/create", midController.TokenValidator, middleware.SuperAdminRankAuthenticator, adminController.Adminmake)
	adminGroup.GET("/get", midController.TokenValidator, middleware.AdminRankAuthenticator, adminController.Adminget)
	adminGroup.DELETE("/del", midController.TokenValidator, middleware.SuperAdminRankAuthenticator, adminController.AdminDel)
	//adminGroup.POST("/login", adminController.AdminLogin)
	adminGroup.PUT("/update", midController.TokenValidator, middleware.SuperAdminRankAuthenticator, adminController.AdminUpdate)

	superadminGroup := defaultRouter.Group("superadmin")
	superadminGroup.POST("/create", midController.TokenValidator, middleware.SuperAdminRankAuthenticator, superadminController.SuperAdminmake)
	superadminGroup.GET("/get", midController.TokenValidator, middleware.SuperAdminRankAuthenticator, superadminController.SuperAdminget)
	superadminGroup.DELETE("/del", midController.TokenValidator, middleware.SuperAdminRankAuthenticator, superadminController.SuperAdminDel)
	//superadminGroup.POST("/login", superadminController.SuperAdminLogin)

	restGroup := defaultRouter.Group("rest")
	restGroup.POST("/create", midController.TokenValidator, middleware.AdminRankAuthenticator, restcontroller.RestMake)
	restGroup.GET("/get", midController.TokenValidator, restcontroller.Restget)
	restGroup.DELETE("/del", midController.TokenValidator, middleware.AdminRankAuthenticator, restcontroller.RestDel)
	restGroup.PUT("/update", midController.TokenValidator, middleware.AdminRankAuthenticator, restcontroller.RestUpdate)
	restGroup.GET("/menu", midController.TokenValidator, restcontroller.MenuGet)
	restGroup.GET("/dist", restcontroller.Getbydist)
	restGroup.GET("/list", midController.TokenValidator, restcontroller.ListRest)

	dishGroup := restGroup.Group("/dish")
	dishGroup.POST("/create", midController.TokenValidator, restcontroller.AddDish)
	dishGroup.DELETE("/del", midController.TokenValidator, restcontroller.DelDish)
	dishGroup.PUT("/update", midController.TokenValidator, restcontroller.UpdDish)

	defaultRouter.GET("/logout", midController.TokenValidator, userController.Logout)

	return defaultRouter
}

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurant/pkg/controller"
	"github.com/restaurant/pkg/database"
	"github.com/restaurant/pkg/middleware"
	"net/http"
)

type Router struct {
	database.Database
}

func NewRouter (db database.Database) *Router {
	return &Router{
		db,
	}
}


func (r *Router) Router() *gin.Engine {

	defaultRouter := gin.Default()

	defaultRouter.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"msg" : "Hello World !!",
		})
	})

	userController := controller.NewUserController(r)
	adminController := controller.NewAdminController(r)
	superadminController := controller.NewSuperAdminController(r)
	restcontroller := controller.NewRestController(r)
	midController := middleware.NewMidController(r)



	userGroup := defaultRouter.Group("user")
	userGroup.POST("/create",midController.TokenValidator,middleware.AdminRankAuthenticator, userController.Usermake)
	userGroup.GET ("/get",midController.TokenValidator,middleware.AdminRankAuthenticator,userController.Userget)
	userGroup.DELETE("/del",midController.TokenValidator,middleware.SuperAdminRankAuthenticator, userController.UserDel)
	userGroup.POST("/login",userController.UserLogin)
	userGroup.GET("/list",midController.TokenValidator,middleware.AdminRankAuthenticator,userController.ListUser)
	userGroup.PUT("/update",midController.TokenValidator,middleware.AdminRankAuthenticator,userController.UserUpdate)
	//userGroup.GET("/rest",midController.TokenValidator,middleware.AdminRankAuthenticator,userController.UserRest)





	adminGroup := defaultRouter.Group("admin")
	adminGroup.POST("/create",midController.TokenValidator,middleware.SuperAdminRankAuthenticator, adminController.Adminmake)
	adminGroup.GET ("/get",midController.TokenValidator,middleware.AdminRankAuthenticator,adminController.Adminget)
	adminGroup.DELETE("/del",midController.TokenValidator,middleware.SuperAdminRankAuthenticator, adminController.AdminDel)
	adminGroup.POST("/login",adminController.AdminLogin)
	adminGroup.PUT("/update",midController.TokenValidator,middleware.SuperAdminRankAuthenticator,adminController.AdminUpdate)

	superadminGroup := defaultRouter.Group("superadmin")
	superadminGroup.POST("/create",midController.TokenValidator,middleware.SuperAdminRankAuthenticator, superadminController.SuperAdminmake)
	superadminGroup.GET ("/get",midController.TokenValidator,middleware.SuperAdminRankAuthenticator, superadminController.SuperAdminget)
	superadminGroup.DELETE("/del",midController.TokenValidator,middleware.SuperAdminRankAuthenticator, superadminController.SuperAdminDel)
	superadminGroup.POST("/login",superadminController.SuperAdminLogin)


	restGroup := defaultRouter.Group("rest")
	restGroup.POST("/create",midController.TokenValidator,middleware.AdminRankAuthenticator,restcontroller.RestMake)
	restGroup.GET("/get",midController.TokenValidator,restcontroller.Restget)
	restGroup.DELETE("/del",midController.TokenValidator,middleware.AdminRankAuthenticator,restcontroller.RestDel)
	restGroup.PUT("/update",midController.TokenValidator,middleware.AdminRankAuthenticator,restcontroller.RestUpdate)
	restGroup.GET("/menu",midController.TokenValidator,restcontroller.MenuGet)
	restGroup.GET("/dist",restcontroller.Getbydist)
	restGroup.GET("/list",midController.TokenValidator,restcontroller.ListRest)

	dishGroup := restGroup.Group("/dish")
	dishGroup.POST("/create",midController.TokenValidator,restcontroller.AddDish)
	dishGroup.DELETE("/del",midController.TokenValidator,restcontroller.DelDish)
	dishGroup.PUT("/update",midController.TokenValidator,restcontroller.UpdDish)


	defaultRouter.GET("/logout",midController.TokenValidator,userController.Logout)


	return defaultRouter
}
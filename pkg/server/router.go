package server

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurant/pkg/controller"
	"github.com/restaurant/pkg/database"
	"github.com/restaurant/pkg/middleware"
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

	userController := controller.NewUserController(r)
	adminController := controller.NewAdminController(r)
	superadminController := controller.NewSuperAdminController(r)
	restcontroller := controller.NewRestController(r)



	userGroup := defaultRouter.Group("user")
	userGroup.POST("/create",middleware.TokenValidator,middleware.SuperAdminRankAuthenticator, userController.Usermake)
	userGroup.GET ("/get",middleware.TokenValidator,userController.Userget)
	userGroup.DELETE("/del",middleware.TokenValidator,middleware.SuperAdminRankAuthenticator, userController.UserDel)
	userGroup.POST("/login",userController.UserLogin)
	userGroup.GET("/list",middleware.TokenValidator,middleware.AdminRankAuthenticator,userController.ListUser)






	adminGroup := defaultRouter.Group("admin")
	adminGroup.POST("/create",middleware.TokenValidator,middleware.SuperAdminRankAuthenticator, adminController.Adminmake)
	adminGroup.GET ("/get",middleware.TokenValidator,adminController.Adminget)
	adminGroup.DELETE("/del",middleware.TokenValidator,middleware.SuperAdminRankAuthenticator, adminController.AdminDel)
	adminGroup.POST("/login",adminController.AdminLogin)


	superadminGroup := defaultRouter.Group("superadmin")
	superadminGroup.POST("/create",middleware.TokenValidator,middleware.SuperAdminRankAuthenticator, superadminController.SuperAdminmake)
	superadminGroup.GET ("/get",middleware.TokenValidator, superadminController.SuperAdminget)
	superadminGroup.DELETE("/del",middleware.TokenValidator,middleware.SuperAdminRankAuthenticator, superadminController.SuperAdminDel)
	superadminGroup.POST("/login",superadminController.SuperAdminLogin)


	restGroup := defaultRouter.Group("rest")
	restGroup.POST("/create",middleware.TokenValidator,middleware.AdminRankAuthenticator,restcontroller.RestMake)
	restGroup.GET("/get",middleware.TokenValidator,restcontroller.Restget)
	restGroup.DELETE("/del",middleware.TokenValidator,middleware.AdminRankAuthenticator,restcontroller.RestDel)
	restGroup.PUT("/update",middleware.TokenValidator,middleware.AdminRankAuthenticator,restcontroller.RestUpdate)
	restGroup.GET("/menu",middleware.TokenValidator,restcontroller.MenuGet)
	restGroup.GET("/dist",restcontroller.Getbydist)

	dishGroup := restGroup.Group("/dish")
	dishGroup.POST("/create",middleware.TokenValidator,restcontroller.AddDish)
	dishGroup.DELETE("/create",middleware.TokenValidator,restcontroller.DelDish)
	dishGroup.PUT("/update",middleware.TokenValidator,restcontroller.UpdDish)


	return defaultRouter
}
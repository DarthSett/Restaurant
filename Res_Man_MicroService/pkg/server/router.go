package server

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurant/Res_Man_MicroService/pkg/controller"
	"github.com/restaurant/Res_Man_MicroService/pkg/database"
	middleware2 "github.com/restaurant/Res_Man_MicroService/pkg/middleware"
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


	userController := controller.NewUserController(r)
	adminController := controller.NewAdminController(r)
	superadminController := controller.NewSuperAdminController(r)
	restcontroller := controller.NewRestController(r)
	midController := middleware2.NewMidController(r)

	userGroup := defaultRouter.Group("user")
	userGroup.POST("/create", midController.TokenValidator, middleware2.AdminRankAuthenticator, userController.Usermake)
	userGroup.GET("/get", midController.TokenValidator, middleware2.AdminRankAuthenticator, userController.Userget)
	userGroup.DELETE("/del", midController.TokenValidator, middleware2.SuperAdminRankAuthenticator, userController.UserDel)
	//userGroup.POST("/login", userController.UserLogin)
	userGroup.GET("/list", midController.TokenValidator, middleware2.AdminRankAuthenticator, userController.ListUser)
	userGroup.PUT("/update", midController.TokenValidator, middleware2.AdminRankAuthenticator, userController.UserUpdate)
	//userGroup.GET("/rest",midController.TokenValidator,middleware.AdminRankAuthenticator,userController.UserRest)

	adminGroup := defaultRouter.Group("admin")
	adminGroup.POST("/create", midController.TokenValidator, middleware2.SuperAdminRankAuthenticator, adminController.Adminmake)
	adminGroup.GET("/get", midController.TokenValidator, middleware2.AdminRankAuthenticator, adminController.Adminget)
	adminGroup.DELETE("/del", midController.TokenValidator, middleware2.SuperAdminRankAuthenticator, adminController.AdminDel)
	//adminGroup.POST("/login", adminController.AdminLogin)
	adminGroup.PUT("/update", midController.TokenValidator, middleware2.SuperAdminRankAuthenticator, adminController.AdminUpdate)

	superadminGroup := defaultRouter.Group("superadmin")
	superadminGroup.POST("/create", midController.TokenValidator, middleware2.SuperAdminRankAuthenticator, superadminController.SuperAdminmake)
	superadminGroup.GET("/get", midController.TokenValidator, middleware2.SuperAdminRankAuthenticator, superadminController.SuperAdminget)
	superadminGroup.DELETE("/del", midController.TokenValidator, middleware2.SuperAdminRankAuthenticator, superadminController.SuperAdminDel)
	//superadminGroup.POST("/login", superadminController.SuperAdminLogin)

	restGroup := defaultRouter.Group("rest")
	restGroup.POST("/create", midController.TokenValidator, middleware2.AdminRankAuthenticator, restcontroller.RestMake)
	restGroup.GET("/get", midController.TokenValidator, restcontroller.Restget)
	restGroup.DELETE("/del", midController.TokenValidator, middleware2.AdminRankAuthenticator, restcontroller.RestDel)
	restGroup.PUT("/update", midController.TokenValidator, middleware2.AdminRankAuthenticator, restcontroller.RestUpdate)
	restGroup.GET("/menu", midController.TokenValidator, restcontroller.MenuGet)
	restGroup.GET("/dist", restcontroller.Getbydist)
	restGroup.GET("/list", midController.TokenValidator, restcontroller.ListRest)

	dishGroup := restGroup.Group("/dish")
	dishGroup.POST("/create", midController.TokenValidator, restcontroller.AddDish)
	dishGroup.DELETE("/del", midController.TokenValidator, restcontroller.DelDish)
	dishGroup.PUT("/update", midController.TokenValidator, restcontroller.UpdDish)

	defaultRouter.POST("/login", userController.Login)
	defaultRouter.GET("/logout", midController.TokenValidator, userController.Logout)

	return defaultRouter
}

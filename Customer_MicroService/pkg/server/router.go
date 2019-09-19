package server

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurant/Customer_MicroService/pkg/controllers"
	"github.com/restaurant/Customer_MicroService/pkg/database"
	"net/http"
	//"time"
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

	c := controllers.NewCustController(r.Database)

	defaultRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Hello World !!",
		})
	})

	defaultRouter.POST("/create", c.CustMake)
	defaultRouter.POST("/login",c.CustLogin)
	defaultRouter.GET("/list",c.TokenValidator,c.CustGetResList)
	defaultRouter.GET("/logout",c.TokenValidator,c.Logout)



	return defaultRouter
}



package server

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurant/Reservation_MicroService/pkg/database"
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

	return defaultRouter
}
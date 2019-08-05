package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/restaurant/pkg/controller"
	"net/http"
	"strconv"
)

func AdminRankAuthenticator(c *gin.Context) {
	t := c.GetHeader("token")
	if t == "" {
		panic("no token sent")
	}

	claims,err := controller.GetTokenClaims(t)
	if err != nil {
		panic("token not valid: " + err.Error())
		c.AbortWithError(http.StatusUnauthorized,err)
	}
	rank,_ := strconv.Atoi(fmt.Sprintf("%v",claims["rank"]))
	//	claims := token.Claims.(jwt.MapClaims)

	if rank == 1 || rank == 2{
		c.Next()
	} else {
		c.AbortWithError(http.StatusUnauthorized,fmt.Errorf("user not an admin"))
	}


}

func SuperAdminRankAuthenticator(c *gin.Context) {
	t := c.GetHeader("token")
	if t == "" {
		panic("no token sent")
	}

	claims,err := controller.GetTokenClaims(t)
	if err != nil {
		panic("token not valid: " + err.Error())
		c.AbortWithError(http.StatusUnauthorized,err)
	}
	rank,_ := strconv.Atoi(fmt.Sprintf("%v",claims["rank"]))
	if rank == 2 {
		c.Next()
	} else {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("user not a Super Admin"))
	}
}

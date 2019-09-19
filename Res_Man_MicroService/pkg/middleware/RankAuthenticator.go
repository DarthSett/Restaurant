package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AdminRankAuthenticator(c *gin.Context) {
	t := c.GetHeader("token")
	if t == "" {
		panic("no token sent")
	}

	clms, _ := c.Get("claims")
	claims := clms.(jwt.MapClaims)
	rank, _ := strconv.Atoi(fmt.Sprintf("%v", claims["rank"]))
	//	claims := token.Claims.(jwt.MapClaims)

	if rank == 1 || rank == 2 {
		c.Next()
	} else {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("user not an admin"))
	}

}

func SuperAdminRankAuthenticator(c *gin.Context) {
	t := c.GetHeader("token")
	if t == "" {
		panic("no token sent")
	}

	clms, _ := c.Get("claims")
	claims := clms.(jwt.MapClaims)
	rank, _ := strconv.Atoi(fmt.Sprintf("%v", claims["rank"]))

	if rank == 2 {
		c.Next()
	} else {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("user not a Super Admin"))
	}
}

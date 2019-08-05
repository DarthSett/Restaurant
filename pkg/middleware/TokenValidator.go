package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurant/pkg/controller"
	"net/http"
)

func TokenValidator(c *gin.Context) {
	t := c.GetHeader("token")
	if t == "" {
		panic("no token sent")
	}

	claims,err := controller.GetTokenClaims(t)
	if err != nil {
		panic("token not valid: " + err.Error())
		c.AbortWithError(http.StatusUnauthorized,err)
	}

	//	claims := token.Claims.(jwt.MapClaims)

	if claims.Valid() == nil{
		c.Next()
	}


}
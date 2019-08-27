package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurant/pkg/controller"
	"github.com/restaurant/pkg/database"
	"net/http"
)

type MidController struct {
	database.Database
}

func NewMidController(db database.Database) *MidController {
	return &MidController{db}
}

func (u *MidController) TokenValidator(c *gin.Context) {
	t := c.GetHeader("token")
	if t == "" {
		panic("no token sent")
	}
	flag, err := u.Checktoken(t)
	if err != nil {
		panic("error checking with deletedtokens: " + err.Error())
		c.AbortWithError(http.StatusUnauthorized, err)
	}
	if flag == false {
		panic("token already logged out")
	}
	claims, err := controller.GetTokenClaims(t)
	if err != nil {
		panic("token not valid: " + err.Error())
		c.AbortWithError(http.StatusUnauthorized, err)
	}
	println(claims["id"])
	c.Set("claims", claims)

	//	claims := token.Claims.(jwt.MapClaims)

	if claims.Valid() == nil {
		c.Next()
	}
}

package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouter(egin *gin.Engine) *gin.RouterGroup {
	user := egin.Group("/user")
	{
		user.GET("/info", func(ctx *gin.Context) {
			var info map[string]interface{}
			info["name"] = "JsonSeaver"
			info["age"] = 26
			ctx.JSON(http.StatusOK, info)
		})
	}

	return user
}
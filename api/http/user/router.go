package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouter(egine *gin.Engine) {
	user := egine.Group("/user")
	{
		user.GET("/info", func(ctx *gin.Context) {
			info := make(map[string]interface{})
			info["age"] = 26
			info["name"] = "JsonSeaver"
			ctx.JSON(http.StatusOK, info)
		})

		user.GET("/height", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "187")
		})
	}
}
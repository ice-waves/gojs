package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ice-waves/gojs/api/http/user"
	"net/http"
)

func main() {
	fmt.Println("Welcome to GoJs!")
	egine := gin.Default()
	egine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,"Welcome to GoJs!")
	})
	user.RegisterRouter(egine)


	egine.Run(":8081")
}

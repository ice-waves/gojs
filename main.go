package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ice-waves/gojs/api/http/user"
	"net/http"
)

type jsgin struct {
	egine *gin.Engine
}

func main() {
	fmt.Println("Welcome to GoJs!")
	jsgin := new(jsgin)
	jsgin.egine = gin.Default()

	user.RegisterRouter(jsgin.egine)

	jsgin.egine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,"Welcome to GoJs!")
	})

	jsgin.Run()
}

func (jsgin jsgin) Run()  {
	jsgin.egine.Run(":8081")
}
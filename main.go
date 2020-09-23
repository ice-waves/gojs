package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ice-waves/gojs/api/http/user"
	"github.com/ice-waves/gojs/configs/conf"
	"net/http"
)

func main() {
	fmt.Println("Welcome to GoJs!")
	egine := gin.Default()
	egine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,"Welcome to GoJs!")
	})
	user.RegisterRouter(egine)

	m := new(conf.Database)
	databaseConf := m.GetConf()
	fmt.Println(databaseConf.Mysql.DSN)
	egine.Run(":8081")
}

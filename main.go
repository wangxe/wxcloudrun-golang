package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	engine := gin.Default()

	//engine.GET("/", service.IndexHandler)
	//engine.GET("/api/count", service.CounterHandler)
	//engine.POST("/api/count", service.CounterHandler)
	engine.GET("/auth", service.AuthorizeHandler)
	engine.GET("/code", service.GetCode)
	engine.GET("/token/:appId", service.SaveToken)
	engine.GET("/getToken", service.GetToken)

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}

}

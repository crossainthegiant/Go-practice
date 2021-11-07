package main

import (
	"gin-example/controllers"
	"gin-example/middlewares"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := gin.Default()
	server.Use(middlewares.MyAuth())

	server.GET("/ping", func(context *gin.Context) {
		context.String(200,"%s","this is pong")
	})


	server.Static("/resources","./resources")//serve静态文件
	server.StaticFile("myfavorpic","./resources/图片.jpg")

	picControllers := controllers.NewPicController()
	picGroup := server.Group("/pics")
	picGroup.Use(middlewares.MyLogger())
	picGroup.GET("/", picControllers.GETALL)
	picGroup.POST("/", picControllers.CREATE)
	picGroup.PUT("/:id", picControllers.UPDATE)
	picGroup.DELETE("/:id", picControllers.DELETE)

	log.Fatal(server.Run("localhost:8090"))
}

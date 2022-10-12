package main

import (
	"golang-rest-api-template/controllers"
	"golang-rest-api-template/models"

	"github.com/gin-gonic/gin"
	docs "github.com/golang-rest-api-template/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	models.ConnectDatabase()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", controllers.Helloworld)
		v1.GET("/books", controllers.FindBooks)
		v1.POST("/books", controllers.CreateBook)
		v1.GET("/books/:id", controllers.FindBook)
		v1.PUT("/books/:id", controllers.UpdateBook)
		v1.DELETE("/books/:id", controllers.DeleteBook)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8001")
}
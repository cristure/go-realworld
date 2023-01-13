package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-realworld/controllers"
	"github.com/go-realworld/models"
)

func main() {

	models.ConnectDataBase()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)

	router.Run(":8080")

}

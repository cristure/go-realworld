package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-realworld/controllers"
	"github.com/go-realworld/middlewares"
	"github.com/go-realworld/models"
)

func main() {

	models.ConnectDataBase()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	public.POST("/admin", controllers.LoginAdmin)

	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	protected.POST("/place-bet", controllers.PlaceBet)

	router.Run(":8080")

}

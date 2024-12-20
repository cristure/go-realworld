package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/go-realworld/controllers"
	"github.com/go-realworld/middlewares"
	"github.com/go-realworld/models"
)

func main() {
	models.ConnectDataBase()
	godotenv.Load(".env")

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	public.POST("/login-admin", controllers.LoginAdmin)
	public.GET("/articles/:slug", controllers.GetArticle)
	public.GET("/articles", controllers.ListArticles)

	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	protected.PUT("/user", controllers.UpdateUser)
	protected.GET("/profiles/:username", controllers.GetProfile)
	protected.POST("/profiles/:username/follow", controllers.FollowUser)
	protected.POST("/profiles/:username/unfollow", controllers.UnfollowUser)
	protected.POST("/articles", controllers.CreateArticle)
	protected.PUT("/articles/:slug", controllers.UpdateArticle)
	protected.DELETE("/articles/:slug", controllers.DeleteArticle)

	protected.GET("/articles/feed", controllers.FeedArticles)
	protected.POST("/articles/:slug/favorite", controllers.FavoriteArticle)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")

}

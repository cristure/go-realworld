package main

import (
	"log"
	"net/http"

	apphttp "github.com/go-realworld/internal/adapter/handler/http"
	"github.com/go-realworld/internal/adapter/storage"
	"github.com/go-realworld/internal/adapter/storage/repository"
)

func main() {
	db, err := storage.New()
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	userRepository := repository.NewUser(db)

	userHandler := apphttp.NewUserHandler(userRepository)
	router := apphttp.NewRouter(userHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

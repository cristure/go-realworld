package main

import (
	"log"

	"github.com/go-realworld/internal/adapter/storage"
	"github.com/go-realworld/internal/adapter/storage/repository"
	"github.com/go-realworld/internal/core/domain"
)

func main() {
	db, err := storage.New()
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	articleRepo := repository.NewArticle(db)
	err = articleRepo.Create(&domain.Article{
		Slug:          "whatever",
		Title:         "whatever",
		Description:   "whatever",
		Body:          "whatever",
		Tags:          nil,
		FavoriteCount: 0,
	})
	if err != nil {
		panic(err)
	}
}

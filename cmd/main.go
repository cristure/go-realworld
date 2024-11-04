package main

import (
	"fmt"
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

	//userRepo := repository.NewUser(db)
	//err = userRepo.Create(&domain.User{
	//	Username: "iceblast14",
	//	Password: "iceblast13",
	//	Email:    "iceblast13@gmail.com",
	//	Bio:      "someBio",
	//})
	//if err != nil {
	//	panic(err)
	//}

	articleRepo := repository.NewArticle(db)
	err = articleRepo.Create(&domain.Article{
		Slug:          "whatever",
		Title:         "whatever",
		Description:   "whatever",
		Body:          "whatever",
		Tags:          nil,
		FavoriteCount: 0,
		UserID:        1,
	})
	if err != nil {
		panic(err)
	}

	feed, err := articleRepo.Feed(0, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(feed[])
}

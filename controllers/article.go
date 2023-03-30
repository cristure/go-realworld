package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-realworld/models"
	"github.com/go-realworld/token"
	"net/http"
)

type ArticleInput struct {
	Title       string
	Description string
	Body        string
	TagList     []string
}

func CreateArticle(c *gin.Context) {
	var input ArticleInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a := models.Article{}
	a.Title = input.Title
	a.Description = input.Description
	a.Body = input.Body

	a.Tags = make([]*models.Tag, 0)
	for _, t := range input.TagList {
		a.Tags = append(a.Tags, &models.Tag{Name: t})
	}

	userId, err := token.ExtractTokenID(c)
	if err != nil {
		return
	}
	a.UserID = userId

	_, err = a.SaveArticle()
	if err != nil {
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "validated!"})
}

func ListArticles(c *gin.Context) {
	//TODO: Add query params

	articles, err := models.ListArticles()
	if err != nil {
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": articles})
}

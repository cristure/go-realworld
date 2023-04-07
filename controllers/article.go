package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-realworld/models"
	"github.com/go-realworld/token"
	"net/http"
	"strconv"
	"strings"
)

type ArticleInput struct {
	Title       string
	Description string
	Body        string
	TagList     []string
}

type ArticleResponse struct {
	Article models.Article `json:"article"`
	Author  models.Profile `json:"author"`
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
	a.Slug = strings.ToLower(strings.Replace(input.Title, " ", "-", -1))

	a.Tags = make([]*models.Tag, 0)
	for _, t := range input.TagList {
		tt, err := models.FindTagByName(t)
		if err != nil {
			a.Tags = append(a.Tags, &models.Tag{Name: t})
		} else {
			a.Tags = append(a.Tags, tt)
		}
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
	tag := c.Query("tag")
	author := c.Query("author")
	favorited := c.Query("favorited")
	limit := c.Query("limit")
	offset := c.Query("offset")

	var result []*models.Article
	articles, err := models.ListArticles()
	if err != nil {
		return
	}

	if tag != "" || author != "" || favorited != "" || limit != "" || offset != "" {
		var limitInt int
		var offsetInt int

		if limit == "" {
			limitInt = 20
		} else {
			if limitInt, err = strconv.Atoi(limit); err != nil {
				return
			}
		}

		if offset == "" {
			offsetInt = 0
		} else {
			if offsetInt, err = strconv.Atoi(offset); err != nil {
				return
			}
		}

		result = models.FilterArticles(articles, models.FilterArticle{
			Tag:       tag,
			Author:    author,
			Favorited: favorited,
			Limit:     limitInt,
			Offset:    offsetInt,
		})
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": result})
}

func FeedArticles(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	limitInt := 20
	offsetInt := 0

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articles, err := u.FeedArticles()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if limit != "" {
		limitInt, _ = strconv.Atoi(limit)
	}

	if offset != "" {
		offsetInt, _ = strconv.Atoi(offset)
	}

	if len(articles) > limitInt && len(articles) > offsetInt {
		articles = articles[offsetInt:limitInt]
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": articles})
}

func FavoriteArticle(c *gin.Context) {
	slug := c.Param("slug")

	article, err := models.FindArticleBySlug(slug)
	if err != nil {
		return
	}

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = u.FavoriteArticle(article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newArticle := article
	newArticle.FavoriteCount = article.FavoriteCount + 1

	updateArticle, err := article.UpdateArticle(newArticle)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": updateArticle})
}

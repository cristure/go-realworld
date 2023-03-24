package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-realworld/models"
	"github.com/go-realworld/token"
	"net/http"
)

type ProfileResponse struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

func GetProfile(c *gin.Context) {
	username := c.Param("username")

	user, err := models.GetUserByName(username)
	if err != nil {
		c.JSON(http.StatusNotFound, errors.New("Invalid username"))
		return
	}

	p, err := models.GetUserByName(user.Username)

	user_id, err := token.ExtractTokenID(c)

	currentUser, err := models.GetUserByID(user_id)

	following, err := currentUser.IsFollowing(p.ID)
	if err != nil {
		return
	}

	pr := ProfileResponse{
		p.Username,
		p.Bio,
		p.Image,
		following,
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": pr})
}

func FollowUser(c *gin.Context) {
	username := c.Param("username")

	user, err := models.GetUserByName(username)
	if err != nil {
		c.JSON(http.StatusNotFound, errors.New("Invalid username"))
		return
	}

	user_id, err := token.ExtractTokenID(c)

	currentUser, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusNotFound, errors.New("Invalid username"))
		return
	}

	followUser, err := currentUser.FollowUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.New("Something wrong happened!"))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": followUser})
}

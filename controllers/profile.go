package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-realworld/models"
	"net/http"
)

func GetProfile(c *gin.Context) {
	username := c.Param("username")

	user, err := models.GetUserByName(username)
	if err != nil {
		c.JSON(http.StatusNotFound, errors.New("Invalid username"))
		return
	}

	p, err := models.GetProfileByUserId(user.ID)

	fmt.Println(p.Following)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

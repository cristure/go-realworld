package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-realworld/models"
	"github.com/go-realworld/token"
	"net/http"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := models.LoginCheck(input.Username, input.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username or password" +
			" is incorrect"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "validated!"})
}

func LoginAdmin(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := models.LoginAdmin(input.Username, input.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "username or password" +
			" is incorrect"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func CurrentUser(c *gin.Context) {

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

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func UpdateUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := token.ExtractTokenID(c)
	if err != nil {
		return
	}

	u, err := models.UpdateUser(uid, input)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

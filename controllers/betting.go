package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

type BettingChoice struct {
	Choice string `json:"choice"`
}

func PlaceBet(c *gin.Context) {
	var input BettingChoice

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Choice != "odd" && input.Choice != "even" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("no valid choice")})
		return
	}

	var userParity int
	if input.Choice == "even" {
		userParity = 0
	} else {
		userParity = 1
	}

	parity := rand.Intn(100) % 2

	if parity == userParity {
		c.JSON(http.StatusOK, gin.H{"result": "You win!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "You lose!"})
	}
}

package auth

import (
	"net/http"

	"github.com/faizallmaullana/libraryManagement/models"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// POST /login
func Login(c *gin.Context) {
	var input Auth
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find the user by username
	var user models.Users
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials (username)"})
		return
	}

	// check the password
	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials (password)"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login Sukses", "data": user})
}

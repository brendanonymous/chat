package auth_handlers

import (
	"chat/pkg/rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler AuthHandler) Login(c *gin.Context) {
	var user_login UserLogin
	if err := c.ShouldBindJSON(&user_login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform user authentication logic (check credentials, generate JWT, etc.)

	// TODO: get user from DB by username
	// create dummy user for now
	u := &models.User{
		ID:           1,
		Username:     "bob69",
		Email:        "bob@email.com",
		PasswordHash: "$2a$10$JRUAdVSVXlJHRK5BO3nja.GvbB09byuznEjcDs8IJhOFBj5oIjCXS",
	}

	// check password against stored hash
	err := checkPasswordAgainstHash(user_login.Password, u.PasswordHash)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": err.Error()})
		return
	}

	// generate a JWT token
	token := generateJwtToken(u.ID, u.Username)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

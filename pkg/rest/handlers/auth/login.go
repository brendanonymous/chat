package auth_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler AuthHandler) Login(c *gin.Context) {
	var user UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform user authentication logic (check credentials, generate JWT, etc.)

	// TODO: get user by id
	// dummy vv
	id := 1234

	// password_hash
	_, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Dummy logic: generate a JWT token
	token := generateJwtToken(id, user.Username)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

package auth_handlers

import (
	"chat/pkg/rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler AuthHandler) Register(c *gin.Context) {
	var user_registration UserRegistration
	if err := c.ShouldBindJSON(&user_registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the user's password before storing it
	password_hash, err := hashPassword(user_registration.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	/*registered_user := */
	_ = models.User{
		Username:     user_registration.Username,
		Email:        user_registration.Email,
		PasswordHash: password_hash,
	}

	// TODO: Store registered user in database

	c.Status(http.StatusCreated)
}

package auth_handlers

import (
	"chat/pkg/auth"
	"chat/pkg/rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler AuthHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()

	var user_registration auth.UserRegistration
	if err := c.ShouldBindJSON(&user_registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password_hash, err := auth.HashPassword(user_registration.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	registered_user := &models.User{
		Username:     user_registration.Username,
		Email:        user_registration.Email,
		PasswordHash: password_hash,
	}

	err = handler.DBClient.AddNewUser(ctx, registered_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

package auth_handlers

import (
	"chat/pkg/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler AuthHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var user_login auth.UserLogin
	if err := c.ShouldBindJSON(&user_login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform user authentication logic (check credentials, validate password, generate JWT)

	// get user by username
	user, err := handler.DBClient.GetUserByUsername(ctx, user_login.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// check password against stored hash
	err = auth.CheckPasswordAgainstHash(user_login.Password, user.PasswordHash)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": err.Error()})
		return
	}

	// generate a JWT token
	token := auth.GenerateJwtToken(user.ID, user.Username)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

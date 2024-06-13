package auth_handlers

import (
	"chat/pkg/auth"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// authMiddleware is a middleware to verify JWT tokens
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return auth.SecretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token parsing error"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*auth.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Claims extraction error"})
			c.Abort()
			return
		}

		// Check if the token has expired
		if claims.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
			c.Abort()
			return
		}

		// Set user information in the context for later use
		c.Set("user", auth.AuthenticatedUser{ID: claims.UserID, Username: claims.Username})
	}
}

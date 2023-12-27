package auth_handlers

import (
	"github.com/gin-gonic/gin"
)

var secretKey = generateRandomKey(10)

func main() {
	r := gin.Default()

	AuthHandler := NewAuthHandler()

	// Endpoint for user registration
	r.POST("/register", AuthHandler.Register)

	// Endpoint for user login
	r.POST("/login", AuthHandler.Login)

	// TODO: Protected endpoint EXAMPLE DO NOT DELETE
	// r.GET("/protected", authMiddleware(), func(c *gin.Context) {
	// 	// Access the user information from the context
	// 	user := c.MustGet("user").(User)
	// 	c.JSON(http.StatusOK, gin.H{"message": "Hello, " + user.Username + "!"})
	// })

	r.Run(":8080")
}

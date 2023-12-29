package main

import (
	"net/http"

	"chat/pkg/config"
	"chat/pkg/data"
	auth_handlers "chat/pkg/rest/handlers/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	r := SetupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080
}

func SetupRouter() *gin.Engine {

	dbClient := data.NewDbClient()

	serverConfig := &config.ServerConfig{
		DBClient: dbClient,
	}

	authHandlers := auth_handlers.NewAuthHandler(serverConfig)

	router := gin.Default()

	// create routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "chat app home page!",
		})
	})

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authHandlers.Login)
	}

	// TODO:
	// messageRoutes := router.Group("/message")
	// {
	// }

	// TODO:
	// chatroomRoutes := router.Group("/message")
	// {
	// }

	return router
}

package main

import (
	"net/http"

	"chat/pkg/config"
	"chat/pkg/data"
	auth_handlers "chat/pkg/rest/handlers/auth"
	chatroom_handlers "chat/pkg/rest/handlers/chatroom"
	user_chatroom_handlers "chat/pkg/rest/handlers/user_chatroom"

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
	chatroomHandlers := chatroom_handlers.NewChatroomHandler(serverConfig)
	userChatroomHandlers := user_chatroom_handlers.NewUserChatroomHandler(serverConfig)

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
		authRoutes.POST("/register", authHandlers.Register)
	}

	// TODO:
	// messageRoutes := router.Group("/messages")
	// {
	// }

	// TODO:
	chatroomRoutes := router.Group("/chatrooms")
	{
		chatroomRoutes.GET("/:id", chatroomHandlers.GetByID)
	}

	// TODO:
	userChatroomRoutes := router.Group("/user_chatrooms")
	{
		userChatroomRoutes.GET("user/:id", userChatroomHandlers.ListByUserId)
		userChatroomRoutes.GET("chatroom/:id", userChatroomHandlers.ListByChatroomId)

	}

	return router
}

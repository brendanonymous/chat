package user_chatroom_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler UserChatroomHandler) JoinChatroom(c *gin.Context) {
	ctx := c.Request.Context()

	user := ctx.Value("user")
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	//authedUser := user.(auth.AuthenticatedUser)
}

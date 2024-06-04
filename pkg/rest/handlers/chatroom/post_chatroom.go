package chatroom_handlers

import (
	"chat/pkg/rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler ChatroomHandler) PostChatroom(c *gin.Context) {
	ctx := c.Request.Context()

	var chatroom models.Chatroom
	if err := c.ShouldBindJSON(&chatroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := handler.DBClient.AddChatroom(ctx, &chatroom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

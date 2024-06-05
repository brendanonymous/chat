package chatroom_handlers

import (
	"chat/pkg/rest/models"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (handler ChatroomHandler) PostChatroom(c *gin.Context) {
	ctx := c.Request.Context()

	var chatroom models.Chatroom
	if err := c.ShouldBindJSON(&chatroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsed_name := strings.Trim(chatroom.Name, " ")
	if parsed_name == "" || parsed_name == " " {
		log.Printf("invalid chatroom name: %s", parsed_name)
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid chatroom name"))
		return
	}

	err := handler.DBClient.AddChatroom(ctx, &chatroom)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

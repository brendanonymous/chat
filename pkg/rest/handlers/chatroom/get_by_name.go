package chatroom_handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (handler ChatroomHandler) GetByName(c *gin.Context) {
	ctx := c.Request.Context()

	parsed_name := strings.ReplaceAll(c.Param("name"), "+", " ")

	chatroom, err := handler.DBClient.GetChatroomByName(ctx, parsed_name)
	if err != nil {
		log.Printf("chatroom not found: %s", err.Error())
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, chatroom)
}

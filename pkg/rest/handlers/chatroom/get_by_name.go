package chatroom_handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (handler ChatroomHandler) GetByName(c *gin.Context) {
	ctx := c.Request.Context()

	parsed_name := strings.Trim(strings.ReplaceAll(c.Param("name"), "+", " "), " ")
	if parsed_name == "" || parsed_name == " " || parsed_name[0] == '+' {
		log.Printf("invalid chatroom name: %s", parsed_name)
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid chatroom name"))
		return
	}

	chatroom, err := handler.DBClient.GetChatroomByName(ctx, parsed_name)
	if err != nil {
		log.Printf("chatroom not found: %s", err.Error())
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, chatroom)
}

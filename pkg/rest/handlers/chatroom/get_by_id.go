package chatroom_handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (handler ChatroomHandler) GetByID(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		log.Printf("strconv.ParseInt failed: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	chatroom, err := handler.DBClient.GetChatroomById(ctx, int32(id))
	if err != nil {
		log.Printf("chatroom not found: %s", err.Error())
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, chatroom)
}

package chatroom_handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (handler ChatroomHandler) ListByUserId(c *gin.Context) {
	ctx := c.Request.Context()

	user_id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		log.Printf("strconv.ParseInt failed: %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	chatrooms, err := handler.DBClient.GetChatroomsByUserId(ctx, int32(user_id))
	if err != nil {
		log.Printf("Server error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if chatrooms == nil {
		err := fmt.Errorf("no chatrooms found for user ID %d", user_id)
		log.Printf("Not found: %s", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"Not found": err.Error()})
		return
	}

	c.JSON(http.StatusOK, chatrooms)
}

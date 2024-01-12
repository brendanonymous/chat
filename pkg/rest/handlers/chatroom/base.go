package chatroom_handlers

import (
	"chat/pkg/config"
	"chat/pkg/data"
)

type ChatroomHandler struct {
	DBClient data.DBClientInterface
}

func NewChatroomHandler(serverConfig *config.ServerConfig) ChatroomHandler {
	return ChatroomHandler{
		DBClient: serverConfig.DBClient,
	}
}

package user_chatroom_handlers

import (
	"chat/pkg/config"
	"chat/pkg/data"
)

type UserChatroomHandler struct {
	DBClient data.DBClientInterface
}

func NewUserChatroomHandler(serverConfig *config.ServerConfig) UserChatroomHandler {
	return UserChatroomHandler{
		DBClient: serverConfig.DBClient,
	}
}

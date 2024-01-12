package message_handlers

import (
	"chat/pkg/config"
	"chat/pkg/data"
)

type MessageHandler struct {
	DBClient data.DBClientInterface
}

func NewMessageHandler(serverConfig *config.ServerConfig) MessageHandler {
	return MessageHandler{
		DBClient: serverConfig.DBClient,
	}
}

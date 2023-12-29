package auth_handlers

import (
	"chat/pkg/config"
	"chat/pkg/data"
)

type AuthHandler struct {
	DBClient data.DBClientInterface
}

func NewAuthHandler(serverConfig *config.ServerConfig) AuthHandler {
	return AuthHandler{
		DBClient: serverConfig.DBClient,
	}
}

package data

import (
	"chat/pkg/rest/models"
	"context"
)

func (client DBClient) AddChatroom(ctx context.Context, chatroom *models.Chatroom) error {
	_, err := client.DB.NewInsert().Model(chatroom).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

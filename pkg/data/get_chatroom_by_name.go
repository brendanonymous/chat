package data

import (
	"chat/pkg/rest/models"
	"context"
)

func (client DBClient) GetChatroomByName(ctx context.Context, name string) (*models.Chatroom, error) {
	chatroom := new(models.Chatroom)

	err := client.DB.NewSelect().Model(chatroom).Where("name = ?", name).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

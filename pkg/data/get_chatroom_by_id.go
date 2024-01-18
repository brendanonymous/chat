package data

import (
	"chat/pkg/rest/models"
	"context"
)

func (client DBClient) GetChatroomById(ctx context.Context, id int32) (*models.Chatroom, error) {
	chatroom := new(models.Chatroom)

	err := client.DB.NewSelect().Model(chatroom).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return chatroom, nil
}

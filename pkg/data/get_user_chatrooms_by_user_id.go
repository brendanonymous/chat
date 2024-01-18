package data

import (
	"chat/pkg/rest/models"
	"context"
)

func (client DBClient) GetUserChatroomsByUserId(ctx context.Context, user_id int32) ([]*models.UserChatroom, error) {
	user_chatrooms := new([]*models.UserChatroom)

	err := client.DB.NewSelect().Model(user_chatrooms).Where("user_id = ?", user_id).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return *user_chatrooms, nil
}

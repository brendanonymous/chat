package data

import (
	"chat/pkg/rest/models"
	"context"
)

func (client DBClient) UserIsMemberOfChatroom(ctx context.Context, userId, chatroomId int) (bool, error) {
	var count int
	user_chatroom := new(models.UserChatroom)

	count, err := client.DB.NewSelect().Model(user_chatroom).Where("user_id = ? AND chatroom_id = ?", userId, chatroomId).ScanAndCount(ctx)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

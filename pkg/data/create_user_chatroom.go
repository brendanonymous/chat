package data

import (
	"chat/pkg/rest/models"
	"context"
)

func (client DBClient) AddUserChatroom(ctx context.Context, userChatroom *models.UserChatroom) error {
	_, err := client.DB.NewInsert().Model(userChatroom).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

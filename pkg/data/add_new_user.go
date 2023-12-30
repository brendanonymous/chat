package data

import (
	"chat/pkg/rest/models"
	"context"
)

func (client DBClient) AddNewUser(ctx context.Context, user *models.User) error {
	_, err := client.DB.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

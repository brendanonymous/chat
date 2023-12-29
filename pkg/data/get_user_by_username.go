package data

import (
	"chat/pkg/rest/models"
	"context"
)

func (client DBClient) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user := new(models.User)

	err := client.DB.NewSelect().Model(user).Where("username = ?", username).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

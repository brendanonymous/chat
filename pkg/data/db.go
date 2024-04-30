package data

import (
	"chat/pkg/rest/models"
	"context"
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	pg "github.com/uptrace/bun/driver/pgdriver"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . DBClientInterface
type DBClientInterface interface {
	// messages
	// AddNewMessage(context.Context, *models.Message) error
	// GetMessageById(context.Context, int32) (*models.Message, error)
	// GetMessagesByChatroomId(context.Context, int32) ([]*models.Message, error)
	// DeleteMessage(context.Context, int32) error

	// chatrooms
	GetChatroomById(context.Context, int32) (*models.Chatroom, error)
	// GetChatroomByMessageId(context.Context, int32) (*models.Chatroom, error)
	// GetChatrooms(context.Context) ([]*models.Chatroom, error)

	// user chatrooms
	GetUserChatroomsByUserId(context.Context, int32) ([]*models.UserChatroom, error)
	GetUserChatroomsByChatroomId(context.Context, int32) ([]*models.UserChatroom, error)

	// users
	AddNewUser(context.Context, *models.User) error
	// UpdateUser(context.Context, models.User) error
	GetUserByUsername(context.Context, string) (*models.User, error)
	// GetUserByEmail(context.Context, string) (*models.User, error)
	// DeleteUserById(context.Context, int32) error
}

type DBClient struct {
	DB *bun.DB
}

func NewDbClient() *DBClient {
	dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"

	sqlDb := sql.OpenDB(
		pg.NewConnector(
			pg.WithDSN(dsn),
			pg.WithUser("brendanlauck"),
			pg.WithDatabase("chat_dev")))

	return &DBClient{
		DB: bun.NewDB(sqlDb, pgdialect.New()),
	}
}

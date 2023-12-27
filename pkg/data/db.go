package data

import (
	"chat/pkg/rest/models"

	"github.com/go-pg/pg"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . DBClientInterface
type DBClientInterface interface {
	// messages
	AddNewMessage(*models.Message) error
	GetMessageById(int32) (*models.Message, error)
	GetAllMessagesByChatroomId(int32) ([]*models.Message, error)
	DeleteMessage(int32) error

	// chatrooms
	GetChatroomById(int32) (*models.Chatroom, error)
	GetChatroomByMessageId(int32) (*models.Chatroom, error)
	GetAllChatroomsByUserId(int32) ([]*models.Chatroom, error)
	GetAllChatrooms() ([]*models.Chatroom, error)

	// users
	AddNewUser(models.User) error
	UpdateUser(models.User) error
	GetUserByUsername(string) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	DeleteUserById(int32) error
}

type DBClient struct {
	DB *pg.DB
}

func NewDbClient() *DBClient {
	return &DBClient{
		DB: pg.Connect(
			&pg.Options{
				Addr:     ":5432",
				User:     "brendanlauck",
				Database: "chat_test",
			}),
	}
}

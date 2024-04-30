package user_chatroom_handlers_test

import (
	"chat/pkg/config"
	"chat/pkg/data/datafakes"
	user_chatroom_handlers "chat/pkg/rest/handlers/user_chatroom"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

var (
	fakeDbClient *datafakes.FakeDBClientInterface
)

type UserChatroomHandlersTestSuite struct {
	suite.Suite
	router              *gin.Engine
	userChatroomHandler user_chatroom_handlers.UserChatroomHandler
}

func (suite *UserChatroomHandlersTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	suite.router = gin.New()

	fakeDbClient = &datafakes.FakeDBClientInterface{}

	config := &config.ServerConfig{
		DBClient: fakeDbClient,
	}

	suite.userChatroomHandler = user_chatroom_handlers.NewUserChatroomHandler(config)

	// TODO: add new routes here for UT's
	suite.router.GET("/user_chatrooms/user/:id", suite.userChatroomHandler.ListByUserId)
	suite.router.GET("/user_chatrooms/chatroom/:id", suite.userChatroomHandler.ListByChatroomId)
}

func TestUserChatroomHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(UserChatroomHandlersTestSuite))
}

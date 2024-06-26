package chatroom_handlers_test

import (
	"chat/pkg/config"
	"chat/pkg/data/datafakes"
	chatroom_handlers "chat/pkg/rest/handlers/chatroom"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

var (
	fakeDbClient *datafakes.FakeDBClientInterface
)

type ChatroomHandlersTestSuite struct {
	suite.Suite
	router          *gin.Engine
	chatroomHandler chatroom_handlers.ChatroomHandler
}

func (suite *ChatroomHandlersTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	suite.router = gin.New()

	fakeDbClient = &datafakes.FakeDBClientInterface{}

	config := &config.ServerConfig{
		DBClient: fakeDbClient,
	}

	suite.chatroomHandler = chatroom_handlers.NewChatroomHandler(config)

	suite.router.GET("/chatrooms/:id", suite.chatroomHandler.GetByID)
	suite.router.GET("/chatrooms/name/:name", suite.chatroomHandler.GetByName)
	suite.router.POST("/chatrooms", suite.chatroomHandler.PostChatroom)
}

func TestChatroomHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(ChatroomHandlersTestSuite))
}

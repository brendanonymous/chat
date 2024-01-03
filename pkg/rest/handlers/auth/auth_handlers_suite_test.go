package auth_handlers_test

import (
	"chat/pkg/config"
	"chat/pkg/data/datafakes"
	auth_handlers "chat/pkg/rest/handlers/auth"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

var (
	fakeDbClient *datafakes.FakeDBClientInterface
)

type AuthHandlersTestSuite struct {
	suite.Suite
	router      *gin.Engine
	authHandler auth_handlers.AuthHandler
}

func (suite *AuthHandlersTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	suite.router = gin.New()

	fakeDbClient = &datafakes.FakeDBClientInterface{}

	config := &config.ServerConfig{
		DBClient: fakeDbClient,
	}

	suite.authHandler = auth_handlers.NewAuthHandler(config)

	// TODO: add other handlers here before writing their unit tests
	suite.router.POST("/auth/register", suite.authHandler.Register)
	suite.router.POST("/auth/login", suite.authHandler.Login)
}

func TestAuthHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(AuthHandlersTestSuite))
}

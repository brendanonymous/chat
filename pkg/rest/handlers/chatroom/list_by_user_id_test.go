package chatroom_handlers_test

import (
	"chat/pkg/rest/models"
	"net/http"
	"net/http/httptest"
)

func (suite *ChatroomHandlersTestSuite) TestListByUserIdHandler_GoodRequest_Succeeds() {
	req, err := http.NewRequest(http.MethodGet, "/chatrooms/1", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	testUserChatrooms := []*models.UserChatroom{
		{
			UserID:     1,
			ChatroomID: 1,
		},
		{
			UserID:     1,
			ChatroomID: 2,
		},
	}
	fakeDbClient.GetChatroomsByUserIdReturns(testUserChatrooms, nil)

	suite.router.ServeHTTP(res, req)

	suite.Equal(fakeDbClient.GetChatroomsByUserIdCallCount(), 1)

	suite.Equal(http.StatusOK, res.Code)
}

func (suite *ChatroomHandlersTestSuite) TestListByUserIdHandler_GoodRequest_ChatroomNotFound() {
	req, err := http.NewRequest(http.MethodGet, "/chatrooms/12", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	testUserChatrooms := []*models.UserChatroom{}

	fakeDbClient.GetChatroomsByUserIdReturns(testUserChatrooms, nil)

	suite.router.ServeHTTP(res, req)

	suite.Equal(fakeDbClient.GetChatroomsByUserIdCallCount(), 1)

	suite.Equal(http.StatusNotFound, res.Code)
}

func (suite *ChatroomHandlersTestSuite) TestListByUserIdHandler_BadRequest_BadPathParameter_Fails() {
	req, err := http.NewRequest(http.MethodGet, "/chatrooms/one", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	suite.router.ServeHTTP(res, req)

	suite.Equal(fakeDbClient.GetChatroomsByUserIdCallCount(), 0)

	suite.Equal(http.StatusBadRequest, res.Code)
}

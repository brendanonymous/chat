package user_chatroom_handlers_test

import (
	"chat/pkg/rest/models"
	"net/http"
	"net/http/httptest"
)

func (suite *UserChatroomHandlersTestSuite) TestListByUserIdHandler_GoodRequest_Succeeds() {
	req, err := http.NewRequest(http.MethodGet, "/user_chatrooms/1", nil)
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
	fakeDbClient.GetUserChatroomsByUserIdReturns(testUserChatrooms, nil)

	suite.router.ServeHTTP(res, req)

	suite.Equal(fakeDbClient.GetUserChatroomsByUserIdCallCount(), 1)

	suite.Equal(http.StatusOK, res.Code)
}

func (suite *UserChatroomHandlersTestSuite) TestListByUserIdHandler_GoodRequest_ChatroomNotFound() {
	req, err := http.NewRequest(http.MethodGet, "/user_chatrooms/12", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	testUserChatrooms := []*models.UserChatroom{}

	fakeDbClient.GetUserChatroomsByUserIdReturns(testUserChatrooms, nil)

	suite.router.ServeHTTP(res, req)

	suite.Equal(fakeDbClient.GetUserChatroomsByUserIdCallCount(), 1)

	suite.Equal(http.StatusNotFound, res.Code)
}

func (suite *UserChatroomHandlersTestSuite) TestListByUserIdHandler_BadRequest_BadPathParameter_Fails() {
	req, err := http.NewRequest(http.MethodGet, "/user_chatrooms/one", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	suite.router.ServeHTTP(res, req)

	suite.Equal(fakeDbClient.GetUserChatroomsByUserIdCallCount(), 0)

	suite.Equal(http.StatusBadRequest, res.Code)
}

package user_chatroom_handlers_test

import (
	"chat/pkg/rest/models"
	"net/http"
	"net/http/httptest"
)

func (suite *UserChatroomHandlersTestSuite) TestListByChatroomIdHandler_GoodRequest_Succeeds() {
	req, err := http.NewRequest(http.MethodGet, "/user_chatrooms/chatroom/1", nil)
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
	fakeDbClient.ListUserChatroomsByChatroomIdReturns(testUserChatrooms, nil)

	suite.router.ServeHTTP(res, req)

	suite.Equal(1, fakeDbClient.ListUserChatroomsByChatroomIdCallCount())

	suite.Equal(http.StatusOK, res.Code)
}

func (suite *UserChatroomHandlersTestSuite) TestListByChatroomIdHandler_GoodRequest_NotFound() {
	req, err := http.NewRequest(http.MethodGet, "/user_chatrooms/chatroom/12", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	testUserChatrooms := []*models.UserChatroom{}

	fakeDbClient.ListUserChatroomsByChatroomIdReturns(testUserChatrooms, nil)

	suite.router.ServeHTTP(res, req)

	suite.Equal(1, fakeDbClient.ListUserChatroomsByChatroomIdCallCount())

	suite.Equal(http.StatusNotFound, res.Code)
}

func (suite *UserChatroomHandlersTestSuite) TestListByChatroomIdHandler_BadRequest_BadPathParameter() {
	req, err := http.NewRequest(http.MethodGet, "/user_chatrooms/chatroom/one", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	suite.router.ServeHTTP(res, req)

	suite.Equal(0, fakeDbClient.ListUserChatroomsByChatroomIdCallCount(), 0)

	suite.Equal(http.StatusBadRequest, res.Code)
}

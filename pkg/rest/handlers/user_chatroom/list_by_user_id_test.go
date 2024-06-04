package user_chatroom_handlers_test

import (
	"chat/pkg/rest/models"
	"net/http"
	"net/http/httptest"
)

func (suite *UserChatroomHandlersTestSuite) TestListByUserIdHandler_GoodRequest_Succeeds() {
	req, err := http.NewRequest(http.MethodGet, "/user_chatrooms/user/1", nil)
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
	fakeDbClient.ListUserChatroomsByUserIdReturns(testUserChatrooms, nil)

	suite.router.ServeHTTP(res, req)

	suite.Equal(1, fakeDbClient.ListUserChatroomsByUserIdCallCount())

	suite.Equal(http.StatusOK, res.Code)
}

func (suite *UserChatroomHandlersTestSuite) TestListByUserIdHandler_GoodRequest_NotFound() {
	req, err := http.NewRequest(http.MethodGet, "/user_chatrooms/user/12", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	testUserChatrooms := []*models.UserChatroom{}

	fakeDbClient.ListUserChatroomsByUserIdReturns(testUserChatrooms, nil)

	suite.router.ServeHTTP(res, req)

	suite.Equal(1, fakeDbClient.ListUserChatroomsByUserIdCallCount())

	suite.Equal(http.StatusNotFound, res.Code)
}

func (suite *UserChatroomHandlersTestSuite) TestListByUserIdHandler_BadRequest_BadPathParameter() {
	req, err := http.NewRequest(http.MethodGet, "/user_chatrooms/user/one", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	suite.router.ServeHTTP(res, req)

	suite.Equal(0, fakeDbClient.ListUserChatroomsByUserIdCallCount(), 0)

	suite.Equal(http.StatusBadRequest, res.Code)
}

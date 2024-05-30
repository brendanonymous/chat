package chatroom_handlers_test

import (
	"chat/pkg/rest/models"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func (suite *ChatroomHandlersTestSuite) TestGetByIdHandler_GoodRequest_Succeeds() {
	req, err := http.NewRequest(http.MethodGet, "/chatrooms/1", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	testChatroom := &models.Chatroom{
		ID:   1,
		Name: "wsu alum",
	}
	fakeDbClient.GetChatroomByIdReturns(testChatroom, nil)

	suite.router.ServeHTTP(res, req)

	suite.Equal(fakeDbClient.GetChatroomByIdCallCount(), 1)

	suite.Equal(http.StatusOK, res.Code)
}

func (suite *ChatroomHandlersTestSuite) TestGetByIdHandler_GoodRequest_ChatroomNotFound() {
	req, err := http.NewRequest(http.MethodGet, "/chatrooms/12", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	fakeDbClient.GetChatroomByIdReturns(nil, fmt.Errorf("sql: no rows in result set"))

	suite.router.ServeHTTP(res, req)

	suite.Equal(1, fakeDbClient.GetChatroomByIdCallCount())

	suite.Equal(http.StatusNotFound, res.Code)
}

func (suite *ChatroomHandlersTestSuite) TestGetByIdHandler_BadRequest_BadPathParameter_Fails() {
	req, err := http.NewRequest(http.MethodGet, "/chatrooms/one", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	suite.router.ServeHTTP(res, req)

	suite.Equal(0, fakeDbClient.GetChatroomByIdCallCount())

	suite.Equal(http.StatusBadRequest, res.Code)
}

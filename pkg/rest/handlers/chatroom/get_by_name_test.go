package chatroom_handlers_test

import (
	"chat/pkg/rest/models"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func (suite *ChatroomHandlersTestSuite) TestGetByNameHandler_GoodRequest_Succeeds() {
	req, err := http.NewRequest(http.MethodGet, "/chatrooms/name/wsu+alum", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	testChatroom := &models.Chatroom{
		ID:   1,
		Name: "wsu alum",
	}
	fakeDbClient.GetChatroomByNameReturns(testChatroom, nil)

	suite.router.ServeHTTP(res, req)

	suite.Equal(1, fakeDbClient.GetChatroomByNameCallCount())

	suite.Equal(http.StatusOK, res.Code)
}

func (suite *ChatroomHandlersTestSuite) TestGetByNameHandler_GoodRequest_ChatroomNotFound() {
	req, err := http.NewRequest(http.MethodGet, "/chatrooms/name/uw+alum", nil)
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	fakeDbClient.GetChatroomByNameReturns(nil, fmt.Errorf("sql: no rows in result set"))

	suite.router.ServeHTTP(res, req)

	suite.Equal(1, fakeDbClient.GetChatroomByNameCallCount())

	suite.Equal(http.StatusNotFound, res.Code)
}

func (suite *ChatroomHandlersTestSuite) TestGetByNameHandler_BadRequest_InvalidChatroomName_Fails() {
	res := httptest.NewRecorder()

	for _, badChatroomName := range []string{
		" ",
		"         ",
		" + ",
		" +",
		"+ "} {

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/chatrooms/name/%s", badChatroomName), nil)
		suite.NoError(err)

		req.Header.Set("Content-Type", "application/json")

		suite.router.ServeHTTP(res, req)

		suite.Equal(0, fakeDbClient.GetChatroomByNameCallCount())

		suite.Equal(http.StatusBadRequest, res.Code)
	}
}

package chatroom_handlers_test

import (
	"chat/pkg/rest/models"
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

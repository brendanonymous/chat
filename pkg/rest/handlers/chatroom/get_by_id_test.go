package chatroom_handlers_test

import (
	"chat/pkg/rest/models"
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

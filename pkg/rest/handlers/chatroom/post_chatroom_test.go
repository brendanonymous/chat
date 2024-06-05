package chatroom_handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func (suite *ChatroomHandlersTestSuite) TestPostChatroomHandler_GoodRequest_Succeeds() {
	// prepare a sample POST chatroom request payload
	payload := map[string]interface{}{
		"name": "test-chatroom-name",
	}
	jsonPayload, _ := json.Marshal(payload)

	// create a request using the payload
	req, err := http.NewRequest("POST", "/chatrooms", bytes.NewBuffer(jsonPayload))
	suite.NoError(err)

	// set the request header to indicate JSON content
	req.Header.Set("Content-Type", "application/json")

	// create a response recorder to record the response
	res := httptest.NewRecorder()

	// perform the request
	suite.router.ServeHTTP(res, req)

	// assert db call count
	suite.Equal(1, fakeDbClient.AddChatroomCallCount())

	// assert the response status code
	suite.Equal(http.StatusCreated, res.Code)
}

func (suite *ChatroomHandlersTestSuite) TestPostChatroomHandler_BadRequest_Fails() {
	testNames := []string{
		"Subtest_MissingName",
		"Subtest_WhitespaceName",
	}
	payloads := []map[string]interface{}{
		{},
		{
			"name": "     ",
		},
	}
	for i, payload := range payloads {
		suite.Run(testNames[i], func() {
			jsonPayload, _ := json.Marshal(payload)

			req, err := http.NewRequest("POST", "/chatrooms", bytes.NewBuffer(jsonPayload))
			suite.NoError(err)

			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			suite.router.ServeHTTP(res, req)

			suite.Equal(0, fakeDbClient.AddChatroomCallCount())

			suite.Equal(http.StatusBadRequest, res.Code)
		})
	}
}

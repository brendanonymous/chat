package auth_handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func (suite *AuthHandlersTestSuite) TestAuthHandler_GoodRegistration_Succeeds() {
	// Prepare a sample registration request
	payload := map[string]interface{}{
		"username": "testuser",
		"email":    "test@example.com",
		"password": "testpassword",
	}
	jsonPayload, _ := json.Marshal(payload)

	// Create a request with the sample payload
	req, err := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonPayload))
	assert.NoError(suite.T(), err)

	// Set the request header to indicate JSON content
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	res := httptest.NewRecorder()

	// Perform the request
	suite.router.ServeHTTP(res, req)

	// Assert call count of db client function
	assert.Equal(suite.T(), fakeDbClient.AddNewUserCallCount(), 1)

	// Assert the response status code
	assert.Equal(suite.T(), http.StatusCreated, res.Code)
}

func (suite *AuthHandlersTestSuite) TestAuthHandler_MissingRegistrationField_Fails() {
	testNames := []string{
		"Subtest_MissingUsername",
		"Subtest_MissingEmail",
		"Subtest_MissingPassword",
	}
	payloads := []map[string]interface{}{
		{
			"username": "",
			"email":    "test@example.com",
			"password": "testpassword",
		},
		{
			"username": "testuser",
			"email":    "",
			"password": "testpassword",
		},
		{
			"username": "testuser",
			"email":    "test@example.com",
			"password": "",
		},
	}
	for i, payload := range payloads {
		suite.Run(testNames[i], func() {
			jsonPayload, _ := json.Marshal(payload)

			req, err := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonPayload))
			assert.NoError(suite.T(), err)

			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			suite.router.ServeHTTP(res, req)

			assert.Equal(suite.T(), fakeDbClient.AddNewUserCallCount(), 0)

			assert.Equal(suite.T(), http.StatusBadRequest, res.Code)
		})
	}
}

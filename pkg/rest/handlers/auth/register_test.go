package auth_handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func (suite *AuthHandlersTestSuite) TestSuccessfulRegistration() {
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

	// Assert the response status code
	assert.Equal(suite.T(), http.StatusCreated, res.Code)
}

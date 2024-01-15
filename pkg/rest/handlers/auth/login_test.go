package auth_handlers_test

import (
	"bytes"
	"chat/pkg/rest/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func (suite *AuthHandlersTestSuite) TestLoginHandler_GoodRequest_Succeeds() {
	payload := map[string]interface{}{
		"username": "testuser",
		"password": "testpassword",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonPayload))
	suite.NoError(err)

	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	testUser := &models.User{
		ID:           1,
		Username:     "testuser",
		PasswordHash: "$2a$10$YUDZBRarB8LQfkKIiLopgO.GUaGe8YMRV8vWetRBFQIC3WNMAmUbG",
	}
	fakeDbClient.GetUserByUsernameReturns(testUser, nil)

	suite.router.ServeHTTP(res, req)

	suite.Equal(fakeDbClient.GetUserByUsernameCallCount(), 1)

	suite.Equal(http.StatusOK, res.Code)
}

func (suite *AuthHandlersTestSuite) TestLoginHandler_BadRequest_MissingField_Fails() {
	testNames := []string{
		"Subtest_MissingUsername",
		"Subtest_MissingPassword",
	}
	payloads := []map[string]interface{}{
		{
			"username": "",
			"password": "testpassword",
		},
		{
			"username": "testuser",
			"password": "",
		},
	}
	for i, payload := range payloads {
		suite.Run(testNames[i], func() {
			jsonPayload, _ := json.Marshal(payload)

			req, err := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonPayload))
			suite.NoError(err)

			req.Header.Set("Content-Type", "application/json")

			res := httptest.NewRecorder()

			testUser := &models.User{
				ID:           1,
				Username:     "testuser",
				PasswordHash: "$2a$10$YUDZBRarB8LQfkKIiLopgO.GUaGe8YMRV8vWetRBFQIC3WNMAmUbG",
			}
			fakeDbClient.GetUserByUsernameReturns(testUser, nil)

			suite.router.ServeHTTP(res, req)

			suite.Equal(http.StatusBadRequest, res.Code)
		})
	}
}

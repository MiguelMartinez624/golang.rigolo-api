package integration_tests

import (
	"bytes"
	"encoding/json"
	"github.com/bmizerany/assert"
	"github.com/rigolo-api/common"
	"github.com/rigolo-api/services/authentication/domain/data"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	t.Run("should create account with valid data", func(t *testing.T) {
		// Given
		rr := httptest.NewRecorder()
		requestStruct := data.NewAccount{
			Email:     "tester@test.com",
			Password:  "test",
			LastName:  "tester",
			FirstName: "tesresr",
			Age:       12,
		}
		data, err := json.Marshal(requestStruct)
		assert.Equal(t, err, nil)
		requestBody := bytes.NewReader(data)
		request := httptest.NewRequest(http.MethodPost, "/auth/signup", requestBody)

		// When
		server.ServeHTTP(rr, request)
		expectedResponse := common.Result{Error: nil, Data: true}
		expectedBytes, err := json.Marshal(&expectedResponse)
		if err != nil {
			panic(err)
		}

		//expected
		assert.Equal(t, rr.Body.Bytes(), expectedBytes)
		assert.Equal(t, 200, rr.Code)
	})
}

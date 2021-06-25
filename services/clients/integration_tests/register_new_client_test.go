package integration_tests

import (
	"encoding/json"
	"github.com/bmizerany/assert"
	"github.com/rigolo-api/common"
	"github.com/rigolo-api/services/clients/domain"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterNewClient(t *testing.T) {

	t.Run("should register client with all valid data", func(t *testing.T) {
		// Given
		rr := httptest.NewRecorder()
		requestBody := strings.NewReader(`{"name":"test","phone":"04141234562", "email":"test@tester.com"}`)
		request := httptest.NewRequest(http.MethodPost, "/clients", requestBody)

		// When
		server.ServeHTTP(rr, request)
		expectedResponse := common.Result{Error: nil, Data: "client registed succefully"}
		expectedBytes, err := json.Marshal(&expectedResponse)
		if err != nil {
			panic(err)
		}

		//expected
		assert.Equal(t, rr.Body.Bytes(), expectedBytes)
		assert.Equal(t, 200, rr.Code)
	})

	t.Run("it should not register without a phone", func(t *testing.T) {
		// Given
		rr := httptest.NewRecorder()
		requestBody := strings.NewReader(`{"name":"test","phone":"", "email":"test@tester.com"}`)
		request := httptest.NewRequest(http.MethodPost, "/clients", requestBody)

		// When
		server.ServeHTTP(rr, request)
		expectedResponse := common.Result{Error: domain.NewVerificationLayerError("phone"), Data: nil}
		expectedBytes, err := json.Marshal(&expectedResponse)
		if err != nil {
			panic(err)
		}

		//expected
		assert.Equal(t, rr.Body.Bytes(), expectedBytes)
		assert.Equal(t, 400, rr.Code)
	})

}

package sms

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sample-api/internal/repository"
	"sample-api/internal/tests/mock"
	"testing"
)

func Test_AddData(t *testing.T) {
	validReq, _ := json.Marshal(&AddDataReq{
		Firstname: "kritchat",
	})
	tt := []struct {
		name           string
		userDetailRepo repository.UserDetail
		w              *httptest.ResponseRecorder
		r              *http.Request
		expectedR      int
	}{
		{
			name:           "Should return status 200 when add data successfully",
			userDetailRepo: mock.UserDetail("OK"),
			w:              httptest.NewRecorder(),
			r:              httptest.NewRequest(http.MethodPost, "http://127.0.0.1:9000/sms", bytes.NewReader(validReq)),
			expectedR:      http.StatusOK,
		},
		{
			name:           "Should return status 500 when err while create user detail",
			userDetailRepo: mock.UserDetail("!OK"),
			w:              httptest.NewRecorder(),
			r:              httptest.NewRequest(http.MethodPost, "http://127.0.0.1:9000/sms", bytes.NewReader(validReq)),
			expectedR:      http.StatusInternalServerError,
		},
		{
			name:      "Should return status 400 when request is invalid json format",
			w:         httptest.NewRecorder(),
			r:         httptest.NewRequest(http.MethodPost, "http://127.0.0.1:9000/sms", bytes.NewReader([]byte("invalid"))),
			expectedR: http.StatusBadRequest,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := NewService(tc.userDetailRepo)
			s.AddData(tc.w, tc.r)
			assert.Equal(t, tc.expectedR, tc.w.Result().StatusCode)
		})
	}
}

func Test_validateReq(t *testing.T) {
	validJson, _ := json.Marshal(&AddDataReq{
		Firstname: "AAA",
	})
	invalidJson, _ := json.Marshal(&AddDataReq{
		Firstname: "",
	})
	tt := []struct {
		name      string
		r         *http.Request
		expectedR *AddDataReq
		expectedE error
	}{
		{
			name: "should return valid result when request is valid",
			r:    httptest.NewRequest(http.MethodPost, "http://127.0.0.1:9000/sms/test", bytes.NewReader(validJson)),
			expectedR: &AddDataReq{
				Firstname: "AAA",
			},
			expectedE: nil,
		},
		{
			name:      "should return err when firstname is empty",
			r:         httptest.NewRequest(http.MethodPost, "http://127.0.0.1:9000/sms/test", bytes.NewReader(invalidJson)),
			expectedR: nil,
			expectedE: errors.New("mock err"),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r, e := validateReq(tc.r)
			if tc.expectedE != nil {
				assert.NotNil(t, e)
			}
			assert.Equal(t, tc.expectedR, r)
		})
	}
}

func Test_GetData(t *testing.T) {
	tt := []struct {
		name      string
		w         *httptest.ResponseRecorder
		expectedR []string
	}{
		{
			name:      "should return valid result when get data successfully",
			w:         httptest.NewRecorder(),
			expectedR: []string{"A"},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := service{
				Data: []string{"A"},
			}
			s.GetData(tc.w, nil)
			var r GetDataResp
			_ = json.NewDecoder(tc.w.Result().Body).Decode(&r)
			assert.Equal(t, tc.expectedR, r.Name)
		})
	}
}

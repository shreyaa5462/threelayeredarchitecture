package user

import (
	"ThreeLayeredArchitecture/models"
	"bytes"
	"encoding/json"
	"errors"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_HandleUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockUserServiceInterface(ctrl)
	handler := NewUserHandler(mockService)

	type testCase struct {
		id         int
		desc       string
		method     string
		body       interface{}
		mockFunc   func()
		wantStatus int
	}

	testCases := []testCase{
		{
			id:     1,
			desc:   "GET users - success",
			method: http.MethodGet,
			mockFunc: func() {
				mockService.EXPECT().GetAllUsers().Return([]models.User{
					{ID: 1, Name: "Ram"},
					{ID: 2, Name: "Shyam"},
				}, nil)
			},
			wantStatus: http.StatusOK,
		},
		{
			id:     2,
			desc:   "GET users - failure",
			method: http.MethodGet,
			mockFunc: func() {
				mockService.EXPECT().GetAllUsers().Return(nil, errors.New("db error"))
			},
			wantStatus: http.StatusInternalServerError,
		},
		{
			id:     3,
			desc:   "POST user - success",
			method: http.MethodPost,
			body:   models.User{Name: "Sita"},
			mockFunc: func() {
				mockService.EXPECT().CreateUser(models.User{Name: "Sita"}).
					Return(models.User{ID: 3, Name: "Sita"}, nil)
			},
			wantStatus: http.StatusCreated,
		},
		{
			id:         4,
			desc:       "POST user - invalid JSON",
			method:     http.MethodPost,
			body:       "invalid-json",
			mockFunc:   func() {},
			wantStatus: http.StatusBadRequest,
		},
		{
			id:     5,
			desc:   "POST user - creation failed",
			method: http.MethodPost,
			body:   models.User{Name: "Lakshman"},
			mockFunc: func() {
				mockService.EXPECT().CreateUser(models.User{Name: "Lakshman"}).
					Return(models.User{}, errors.New("insert fail"))
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		var reqBody *bytes.Buffer
		if tc.body != nil {
			if s, ok := tc.body.(string); ok {
				reqBody = bytes.NewBuffer([]byte(s))
			} else {
				b, _ := json.Marshal(tc.body)
				reqBody = bytes.NewBuffer(b)
			}
		} else {
			reqBody = &bytes.Buffer{}
		}

		req := httptest.NewRequest(tc.method, "/users", reqBody)
		w := httptest.NewRecorder()

		if tc.mockFunc != nil {
			tc.mockFunc()
		}

		handler.HandleUsers(w, req)

		if w.Code != tc.wantStatus {
			t.Errorf("Test ID %d - %s: expected status %d, got %d", tc.id, tc.desc, tc.wantStatus, w.Code)
		}
	}
}

func TestUserHandler_HandleUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := NewMockUserServiceInterface(ctrl)
	handler := NewUserHandler(mockService)

	type testCase struct {
		id         int
		desc       string
		url        string
		method     string
		mockFunc   func()
		wantStatus int
	}

	testCases := []testCase{
		{
			id:     1,
			desc:   "GET user by ID - success",
			url:    "/user/1",
			method: http.MethodGet,
			mockFunc: func() {
				mockService.EXPECT().GetUserByID(1).Return(models.User{ID: 1, Name: "Ram"}, nil)
			},
			wantStatus: http.StatusOK,
		},
		{
			id:         2,
			desc:       "Invalid method",
			url:        "/user/1",
			method:     http.MethodPost,
			mockFunc:   nil,
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			id:         3,
			desc:       "Invalid ID format",
			url:        "/user/abc",
			method:     http.MethodGet,
			mockFunc:   nil,
			wantStatus: http.StatusBadRequest,
		},
		{
			id:     4,
			desc:   "User not found",
			url:    "/user/99",
			method: http.MethodGet,
			mockFunc: func() {
				mockService.EXPECT().GetUserByID(99).Return(models.User{}, errors.New("not found"))
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(tc.method, tc.url, nil)
		w := httptest.NewRecorder()

		if tc.mockFunc != nil {
			tc.mockFunc()
		}

		handler.HandleUserByID(w, req)

		if w.Code != tc.wantStatus {
			t.Errorf("Test ID %d - %s: expected status %d, got %d", tc.id, tc.desc, tc.wantStatus, w.Code)
		}
	}
}

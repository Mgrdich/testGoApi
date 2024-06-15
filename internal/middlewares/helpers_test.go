package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MethodsTestCase struct {
	name           string
	path           string
	allowedMethods []string
	testMethods    []string
	expectedStatus int
}

func TestAllowedMethods(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	allowedMethodsTestCases := []MethodsTestCase{
		{
			name:           "Allowed method GET",
			path:           "/",
			allowedMethods: []string{http.MethodGet},
			testMethods:    []string{http.MethodGet},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Allowed methods GET, POST, PUT, DELETE",
			path:           "/",
			allowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
			testMethods:    []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Allowed methods GET RequestMethods POST, PUT, DELETE",
			path:           "/",
			allowedMethods: []string{http.MethodGet},
			testMethods:    []string{http.MethodPost, http.MethodPut, http.MethodDelete},
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, testCase := range allowedMethodsTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			handler := AllowedMethods(testCase.allowedMethods...)(next)

			for _, method := range testCase.testMethods {
				req := httptest.NewRequest(method, testCase.path, nil)
				rr := httptest.NewRecorder()
				handler.ServeHTTP(rr, req)
				fmt.Println(method)

				if status := rr.Code; status != testCase.expectedStatus {
					t.Errorf("handler returned wrong status code: got %v want %v", status, testCase.expectedStatus)
				}
			}
		})
	}
}

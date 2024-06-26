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
	methods        []string
	testMethods    []string
	expectedStatus int
}

var next = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
})

func TestAllowedMethods(t *testing.T) {
	testCases := []MethodsTestCase{
		{
			name:           "Allowed method GET",
			path:           "/",
			methods:        []string{http.MethodGet},
			testMethods:    []string{http.MethodGet},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Allowed methods GET, POST, PUT, DELETE",
			path:           "/",
			methods:        []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
			testMethods:    []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Allowed methods GET RequestMethods POST, PUT, DELETE",
			path:           "/",
			methods:        []string{http.MethodGet},
			testMethods:    []string{http.MethodPost, http.MethodPut, http.MethodDelete},
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			handler := AllowedMethods(testCase.methods...)(next)

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

func TestNotAllowedMethods(t *testing.T) {
	testCases := []MethodsTestCase{
		{
			name:           "Not Allowed Method GET",
			path:           "/",
			methods:        []string{http.MethodGet},
			testMethods:    []string{http.MethodPost, http.MethodPut, http.MethodDelete},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Not Allowed Method POST, PUT, DELETE",
			path:           "/",
			methods:        []string{http.MethodPost, http.MethodPut, http.MethodDelete},
			testMethods:    []string{http.MethodPost, http.MethodPut, http.MethodDelete},
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			handler := NotAllowedMethods(testCase.methods...)(next)

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

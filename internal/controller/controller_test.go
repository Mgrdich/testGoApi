package controller

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ExecuteRequest(
	req *http.Request,
	controllerHandler func(w http.ResponseWriter, r *http.Request),
	ctx context.Context,
) *httptest.ResponseRecorder {
	request := req
	if ctx != nil {
		request = req.WithContext(ctx)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllerHandler)
	handler.ServeHTTP(rr, request)

	return rr
}

func CheckStatusOK(t *testing.T, rr *httptest.ResponseRecorder) {
	t.Helper()

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func CheckStatusCreated(t *testing.T, rr *httptest.ResponseRecorder) {
	t.Helper()

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusCreated, status)
	}
}

func NewRequest(t *testing.T, method, url string, body io.Reader) *http.Request {
	t.Helper()

	req, err := http.NewRequestWithContext(context.Background(), method, url, body)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	return req
}

package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func ExecuteRequest(
	req *http.Request,
	controllerHandler func(w http.ResponseWriter, r *http.Request),
) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllerHandler)
	handler.ServeHTTP(rr, req)

	return rr
}

func CheckStatusOK(t *testing.T, rr *httptest.ResponseRecorder) {
	t.Helper()

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

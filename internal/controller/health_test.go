package controller

import (
	"encoding/json"
	"net/http"
	"testing"

	"testGoApi/internal/test_helpers"
)

func TestHandleGetHealth(t *testing.T) {
	req := test_helpers.NewRequest(t, http.MethodGet, "/health", nil)

	rr := test_helpers.ExecuteRequest(req, HandleGetHealth, nil)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}

	var response healthResponse

	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if response.OK != true {
		t.Errorf("Response back does not match the intention")
	}
}

package controller

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestHandleGetHealth(t *testing.T) {
	req := NewRequest(t, http.MethodGet, "/health", nil)

	rr := ExecuteRequest(req, HandleGetHealth, nil)

	CheckStatusOK(t, rr)

	var response healthResponse

	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if response.OK != true {
		t.Errorf("Response back does not match the intention")
	}
}

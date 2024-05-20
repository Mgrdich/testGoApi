package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestHandleGetHealth(t *testing.T) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/health", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	rr := ExecuteRequest(req, HandleGetHealth)

	CheckStatusOK(t, rr)

	var response healthResponse

	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if response.OK != true {
		t.Errorf("Response back does not match the intention")
	}
}

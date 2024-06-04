package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"testGoApi/internal/test_helpers"
)

var personService = test_helpers.NewMockPersonService()

func Test_GetPersonRouter(t *testing.T) {
	r := chi.NewRouter()

	GetPersonRouter(personService)(r)

	testCases := []struct {
		name   string
		method string
		path   string
	}{
		{name: "GET /", method: http.MethodGet, path: "/"},
		{name: "POST /", method: http.MethodPost, path: "/"},
		{name: "GET /{id}", method: http.MethodGet, path: "/123123123"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(testCase.method, testCase.path, nil)

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			if status := rr.Code; status == http.StatusNotFound {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
			}
		})
	}
}

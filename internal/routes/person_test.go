package routes

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"testGoApi/internal/test_helpers"
)

type HttpRouteTestCase struct {
	name   string
	method string
	path   string
}

var personService = test_helpers.NewMockPersonService()

func Test_GetPersonRouter(t *testing.T) {
	r := chi.NewRouter()
	GetPersonRouter(personService)(r)

	httpRouteTestCase := []HttpRouteTestCase{
		{name: "GET /", method: http.MethodGet, path: "/"},
		{name: "POST /", method: http.MethodPost, path: "/"},
		{name: "GET /{id}", method: http.MethodGet, path: "/123123123"},
	}

	for _, httpRoute := range httpRouteTestCase {
		t.Run(httpRoute.name, func(t *testing.T) {
			rr := validateRegisteredRoute(t, r, httpRoute.method, httpRoute.path)

			if status := rr.Code; status == http.StatusNotFound {
				t.Errorf("Route %s not found. Expected status code: %d. Got: %d.", httpRoute.path, http.StatusNotFound, rr.Code)
			}
		})
	}
}

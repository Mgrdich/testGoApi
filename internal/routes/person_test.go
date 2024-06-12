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
	status int
}

var personService = test_helpers.NewMockPersonService()

func Test_GetPersonRouter(t *testing.T) {
	r := chi.NewRouter()
	GetPersonRouter(personService)(r)

	httpRouteTestCases := []HttpRouteTestCase{
		{name: "GET /", method: http.MethodGet, path: "/", status: http.StatusOK},
		{name: "POST /", method: http.MethodPost, path: "/", status: http.StatusBadRequest},
		{name: "GET /{id}", method: http.MethodGet, path: "/123123123", status: http.StatusBadRequest},
	}

	for _, httpRoute := range httpRouteTestCases {
		t.Run(httpRoute.name, func(t *testing.T) {
			rr := createNewRequest(t, r, httpRoute.method, httpRoute.path)

			if status := rr.Code; status != httpRoute.status {
				t.Errorf(
					"Route %s not found. Expected status code: %d. Got: %d.",
					httpRoute.path, httpRoute.status, rr.Code,
				)
			}
		})
	}
}

func Test_GetPersonRouter_Not_Allowed(t *testing.T) {
	r := chi.NewRouter()
	GetPersonRouter(personService)(r)

	httpRouteTestCases := []HttpRouteTestCase{
		{name: "PUT /{id}", method: http.MethodPut, path: "/121312", status: http.StatusMethodNotAllowed},
		{name: "DELETE /{id}", method: http.MethodDelete, path: "/123123", status: http.StatusMethodNotAllowed},
	}

	for _, httpRoute := range httpRouteTestCases {
		t.Run(httpRoute.name, func(t *testing.T) {
			rr := createNewRequest(t, r, httpRoute.method, httpRoute.path)

			if status := rr.Code; status != httpRoute.status {
				t.Errorf(
					"Route %s not found. Expected status code: %d. Got: %d.",
					httpRoute.path, httpRoute.status, rr.Code,
				)
			}
		})
	}
}

package routes

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"testGoApi/internal/test_helpers"
)

var moviesService = test_helpers.NewMockMovieService()

func Test_GetMoviesRouter(t *testing.T) {
	r := chi.NewRouter()

	GetMoviesRouter(moviesService)(r)

	httpRouteTestCase := []HttpRouteTestCase{
		{name: "GET /", method: http.MethodGet, path: "/"},
		{name: "POST /", method: http.MethodPost, path: "/"},
		{name: "GET /{id}", method: http.MethodGet, path: "/123123123"},
		{name: "DELETE /{id}", method: http.MethodDelete, path: "/123123123"},
		{name: "PUT /{id}", method: http.MethodPut, path: "/123123123"},
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

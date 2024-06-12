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
		{name: "GET /", method: http.MethodGet, path: "/", status: http.StatusOK},
		{name: "POST /", method: http.MethodPost, path: "/", status: http.StatusBadRequest},
		{name: "GET /{id}", method: http.MethodGet, path: "/123123123", status: http.StatusBadRequest},
		{name: "DELETE /{id}", method: http.MethodDelete, path: "/123123123", status: http.StatusBadRequest},
		{name: "PUT /{id}", method: http.MethodPut, path: "/123123123", status: http.StatusBadRequest},
	}

	for _, httpRoute := range httpRouteTestCase {
		t.Run(httpRoute.name, func(t *testing.T) {
			rr := createNewRequest(t, r, httpRoute.method, httpRoute.path)

			if status := rr.Code; status != httpRoute.status {
				t.Errorf("Route %s not found. Expected status code: %d. Got: %d.", httpRoute.path, httpRoute.status, rr.Code)
			}
		})
	}
}

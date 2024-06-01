package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"testGoApi/configs"
	"testGoApi/internal/server"
	"testGoApi/internal/test_helpers"
)

type endpointsTestCase struct {
	name  string
	input string
}

func TestAddRoutes(t *testing.T) {
	configs.SetAppConfig(&configs.AppConfig{
		Environment: "dev",
	})

	s := &server.Server{
		Router: chi.NewRouter(),
	}
	services := &ApplicationServices{
		MovieService:  &test_helpers.MockMovieService{}, // Use mock services for testing
		PersonService: &test_helpers.MockPersonService{},
	}

	// Call the AddRoutes function
	AddRoutes(s, services)

	endpointsTestCases := []endpointsTestCase{
		{name: "Health endpoint", input: "/health"},
		{name: "Person endpoint", input: "/api/v1/person"},
		{name: "Movies endpoint", input: "/api/v1/movies"},
		{name: "Swagger endpoint", input: "/swagger/index.html"},
	}

	for _, endpoint := range endpointsTestCases {
		t.Run(endpoint.name, func(t *testing.T) {
			rr := validateRegisteredRoute(t, s.Router, endpoint.input)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("Route %s not found. Expected status code: %d. Got: %d.", endpoint.input, http.StatusOK, rr.Code)
			}
		})
	}
}

func TestAddRoutesNonDevEnvironment(t *testing.T) {
	configs.SetAppConfig(&configs.AppConfig{
		Environment: "prod",
	})

	s := &server.Server{
		Router: chi.NewRouter(),
	}

	services := &ApplicationServices{
		MovieService:  &test_helpers.MockMovieService{},
		PersonService: &test_helpers.MockPersonService{},
	}

	AddRoutes(s, services)

	endpoint := "/swagger/index.html"

	rr := validateRegisteredRoute(t, s.Router, endpoint)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Route %s not found. Expected status code: %d. Got: %d.", endpoint, http.StatusNotFound, rr.Code)
	}
}

func validateRegisteredRoute(t *testing.T, r *chi.Mux, endpoint string) *httptest.ResponseRecorder {
	t.Helper()

	// Create a request to the endpoint
	rr := createNewRequest(t, r, http.MethodGet, endpoint)

	// Check if the response status code indicates the route is registered
	return rr
}

func createNewRequest(t *testing.T, r *chi.Mux, method string, endpoint string) *httptest.ResponseRecorder {
	t.Helper()
	// Create a request to the endpoint
	req := httptest.NewRequest(method, endpoint, nil)

	// Create a test response recorder
	rr := httptest.NewRecorder()

	// Serve the request using the router
	r.ServeHTTP(rr, req)

	return rr
}

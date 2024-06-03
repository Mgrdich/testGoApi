package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"testGoApi/configs"
	"testGoApi/internal/server"
	"testGoApi/internal/test_helpers"
	"testGoApi/internal/util"
)

type endpointsTestCase struct {
	name   string
	input  string
	status int
}

func TestAddRoutes(t *testing.T) {
	configs.SetAppConfig(&configs.AppConfig{
		Environment: util.DevEnvironment,
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
		{name: "Health endpoint", input: "/health", status: http.StatusOK},
		{name: "Person endpoint", input: "/api/v1/person", status: http.StatusUnauthorized},
		{name: "Movies endpoint", input: "/api/v1/movies", status: http.StatusUnauthorized},
		{name: "Swagger endpoint", input: "/swagger/index.html", status: http.StatusOK},
	}

	for _, endpoint := range endpointsTestCases {
		t.Run(endpoint.name, func(t *testing.T) {
			rr := createNewRequest(t, s.Router, http.MethodGet, endpoint.input)
			if status := rr.Code; status != endpoint.status {
				t.Errorf("Route %s not found. Expected status code: %d. Got: %d.", endpoint.input, endpoint.status, rr.Code)
			}
		})
	}
}

func TestAddRoutesNonDevEnvironment(t *testing.T) {
	configs.SetAppConfig(&configs.AppConfig{
		Environment: util.ProdEnvironment,
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

	rr := createNewRequest(t, s.Router, http.MethodGet, endpoint)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Route %s not found. Expected status code: %d. Got: %d.", endpoint, http.StatusNotFound, rr.Code)
	}
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

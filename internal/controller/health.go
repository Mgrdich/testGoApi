package controller

import (
	"net/http"

	"github.com/go-chi/render"
)

type healthResponse struct {
	OK bool `json:"ok"`
}

func (h *healthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// HandleGetHealth checks the health of the service
// @Summary Health Check
// @Description Returns the health status of the service
// @Tags health
// @Produce json
// @Success 200 {object} healthResponse
// @Router /health [get]
func HandleGetHealth(w http.ResponseWriter, r *http.Request) {
	health := &healthResponse{OK: true}
	_ = render.Render(w, r, health)
}

package controller

import (
	"net/http"

	"github.com/go-chi/render"
)

// HandleGetHealth checks the health of the service
// @Summary Health Check
// @Description Returns the health status of the service
// @Tags health
// @Produce json
// @Success 200 {object} OKDto
// @Router /health [get]
func HandleGetHealth(w http.ResponseWriter, r *http.Request) {
	_ = render.Render(w, r, NewOkDto())
}

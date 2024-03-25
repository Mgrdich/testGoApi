package controller

import (
	"github.com/go-chi/render"
	"net/http"
)

type healthResponse struct {
	OK bool `json:"ok"`
}

func (h *healthResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func HandleGetHealth(w http.ResponseWriter, r *http.Request) {
	health := &healthResponse{OK: true}
	_ = render.Render(w, r, health)
}

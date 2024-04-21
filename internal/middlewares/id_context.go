package middlewares

import (
	"errors"
	"net/http"

	".com/internal/server"
	".com/internal/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func CheckSlugId[K any](w http.ResponseWriter, r *http.Request, getByID util.GetByIDFunc[K]) (*K, error) {
	slugId := chi.URLParam(r, "id")
	if slugId == "" {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return nil, errors.New("slug ID is empty")
	}

	id, err := uuid.Parse(slugId)
	if err != nil {
		_ = render.Render(w, r, server.ErrorBadRequest)
		return nil, errors.New("invalid slug ID")
	}

	storeValue, err := getByID(id)
	if err != nil {
		_ = render.Render(w, r, server.ErrorNotFound)
		return nil, errors.New("ID not found in store")
	}

	return storeValue, nil
}

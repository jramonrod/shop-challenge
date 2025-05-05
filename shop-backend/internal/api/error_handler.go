package api

import (
	"errors"
	"net/http"
	"shop-backend/internal/domain"
)

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidUUID):
		http.Error(w, err.Error(), 400)
	case errors.Is(err, domain.ErrProductNotFound):
		http.Error(w, err.Error(), 404)
	default:
		http.Error(w, err.Error(), 500)
	}
}

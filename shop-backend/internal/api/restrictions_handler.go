package api

import (
	"encoding/json"
	"io"
	"net/http"
	"shop-backend/internal/domain"
	"shop-backend/internal/infra/rest"
	"shop-backend/internal/services"

	"github.com/go-chi/chi"
)

func RegisterRestrictionRoutes(router chi.Router, resService services.RestrictionsService) http.Handler {
	router.Post("/products/{product_id}/restrictions/verify", func(w http.ResponseWriter, r *http.Request) {
		productID := chi.URLParam(r, "product_id")
		if productID == "" {
			http.Error(w, "missing product_id", http.StatusBadRequest)
			return
		}

		var payload struct {
			Parts []string `json:"parts"`
		}

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&payload); err != nil {
			if err == io.EOF {
				json.NewEncoder(w).Encode(rest.ViolationsDTO{
					IsValid:    false,
					Violations: []domain.Violation{},
				})
				return
			}

			http.Error(w, "invalid request payload: "+err.Error(), http.StatusBadRequest)
			return
		}

		if len(payload.Parts) == 0 {
			json.NewEncoder(w).Encode(rest.ViolationsDTO{
				IsValid:    false,
				Violations: []domain.Violation{},
			})
			return
		}

		vio, err := resService.ValidateSelections(productID, payload.Parts)
		if err != nil {
			handleError(w, r, err)
			return
		}

		// this is a verify endpoint so returning a 200 is ok here
		json.NewEncoder(w).Encode(rest.ViolationsDTO{
			IsValid:    len(vio) == 0,
			Violations: vio})
	})

	return router
}

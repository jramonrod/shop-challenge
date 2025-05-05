package api

import (
	"encoding/json"
	"net/http"
	"shop-backend/internal/services"

	"github.com/go-chi/chi"
)

func RegisterPartsRoutes(router chi.Router, partsService services.PartsService) http.Handler {
	router.Get("/products/{product_id}/categories/{category_id}/parts", func(w http.ResponseWriter, r *http.Request) {
		productID := chi.URLParam(r, "product_id")
		if productID == "" {
			http.Error(w, "missing product_id", http.StatusBadRequest)
			return
		}

		categoryID := chi.URLParam(r, "category_id")
		if productID == "" {
			http.Error(w, "missing category_id", http.StatusBadRequest)
			return
		}

		parts, err := partsService.GetAllPartsInCategory(productID, categoryID)
		if err != nil {
			handleError(w, r, err)
			return
		}

		json.NewEncoder(w).Encode(parts)
	})

	return router
}

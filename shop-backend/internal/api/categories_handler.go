package api

import (
	"encoding/json"
	"net/http"
	"shop-backend/internal/services"

	"github.com/go-chi/chi"
)

func RegisterCategoryRoutes(router chi.Router, catService services.CategoriesService) http.Handler {
	router.Get("/products/{product_id}/categories", func(w http.ResponseWriter, r *http.Request) {
		productID := chi.URLParam(r, "product_id")
		if productID == "" {
			http.Error(w, "missing product_id", http.StatusBadRequest)
			return
		}

		categories, err := catService.GetAllCategoriesForAProduct(productID)
		if err != nil {
			handleError(w, r, err)
			return
		}

		json.NewEncoder(w).Encode(categories)
	})

	return router
}

package api

import (
	"encoding/json"
	"net/http"
	"shop-backend/internal/services"

	"github.com/go-chi/chi"
)

func RegisterProductRoutes(router chi.Router, prodService services.ProductsService) http.Handler {
	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		products, err := prodService.GetAllProducts()
		if err != nil {
			handleError(w, r, err)
			return
		}

		json.NewEncoder(w).Encode(products)
	})

	router.Get("/products/{product_slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "product_slug")
		if slug == "" {
			http.Error(w, "missing product_slug", http.StatusBadRequest)
			return
		}

		products, err := prodService.GetProductBySlug(slug)
		if err != nil {
			handleError(w, r, err)
			return
		}

		json.NewEncoder(w).Encode(products)
	})

	return router
}

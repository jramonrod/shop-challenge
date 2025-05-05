package api

import (
	"encoding/json"
	"net/http"
	"shop-backend/internal/services"

	"github.com/go-chi/chi"
)

func RegisterCartRoutes(router chi.Router, cartService services.CartsService) http.Handler {
	router.Get("/carts/{cart_id}", func(w http.ResponseWriter, r *http.Request) {
		cartID := chi.URLParam(r, "cart_id")
		if cartID == "" {
			http.Error(w, "missing cart_id", http.StatusBadRequest)
			return
		}

		cart, err := cartService.GetCart(cartID)
		if err != nil {
			handleError(w, r, err)
			return
		}

		json.NewEncoder(w).Encode(cart)
	})

	router.Post("/cart", func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Parts []string `json:"parts"`
		}

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&payload); err != nil {
			http.Error(w, "invalid request payload: "+err.Error(), http.StatusBadRequest)
			return
		}

		cart, err := cartService.CreateCart(payload.Parts)
		if err != nil {
			handleError(w, r, err)
			return
		}

		json.NewEncoder(w).Encode(cart)
	})

	return router
}

package main

import (
	"log"
	"net/http"
	"os"

	"shop-backend/internal/api"
	"shop-backend/internal/infra/db"
	"shop-backend/internal/infra/repositories"
	"shop-backend/internal/services"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	db := db.MustConnectDB(os.Getenv("DB_DSN"))
	prodRepo := repositories.NewProductsRepo(db)
	catRepo := repositories.NewCategoriesRepo(db)
	partRepo := repositories.NewPartRepo(db)
	resRepo := repositories.NewRestrictionsRepo(db)
	cartRepo := repositories.NewCartsRepo(db)

	router := chi.NewRouter()

	// CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	prodSVC := services.NewProductsService(*prodRepo)
	api.RegisterProductRoutes(router, *prodSVC)

	catSVC := services.NewCategoriesService(*catRepo)
	api.RegisterCategoryRoutes(router, *catSVC)

	partSVC := services.NewPartsService(*partRepo)
	api.RegisterPartsRoutes(router, *partSVC)

	resSVC := services.NewRestrictionsService(*resRepo, *partRepo)
	api.RegisterRestrictionRoutes(router, *resSVC)

	cartSVC := services.NewCartsService(*cartRepo)
	api.RegisterCartRoutes(router, *cartSVC)

	log.Fatal(http.ListenAndServe(":8080", router))
}

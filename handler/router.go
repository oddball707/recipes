package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *Handler) NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(CORS)

	router.Get("/api/health", h.HealthHandler)
	router.Get("/api/readiness", h.ReadinessHandler)
	router.Post("/api/create", h.CreateRecipe)
	router.Post("/api/get", h.GetRecipe)

	return router
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Next
		next.ServeHTTP(w, r)
	})
}

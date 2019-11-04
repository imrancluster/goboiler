package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/imrancluster/goboiler/middlewares"
	"github.com/imrancluster/goboiler/utils"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middlewares.Recovery)
	router.Use(middleware.Logger)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(cors.Handler)
}

func sysHandler(w http.ResponseWriter, r *http.Request) {
	limit := 10
	page := 1
	if r.URL.Query().Get("limit") != "" {
		pageLimit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "cannot convert string data to number")
			return
		}
		limit = pageLimit
	}
	if r.URL.Query().Get("page") != "" {
		pageNo, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "cannot convert string data to number")
			return
		}
		page = pageNo
	}
	fmt.Println(limit / page)
	utils.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Health Check API."}, true)
}

// Router : main router function
func Router() *chi.Mux {

	router.Get("/", sysHandler)

	router.Route("/api/v1", func(r chi.Router) {

		// r.Mount("/users", userRoutes())

	})

	return router
}

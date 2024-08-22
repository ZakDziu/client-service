package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func configureRouter(api *api) *mux.Router {
	router := mux.NewRouter()

	router.Use(CORSMiddleware)

	public := router.PathPrefix("/api/v1").Subrouter()

	publicUsers := public.PathPrefix("/users").Subrouter()
	publicUsers.Use(jsonResponse)

	publicUsers.HandleFunc("/sign_in", api.Auth().SignIn).Methods(http.MethodPost)

	publicBooks := public.PathPrefix("/books").Subrouter()
	publicBooks.Use(jsonResponse)

	publicBooks.HandleFunc("", api.Books().GetAll).Methods(http.MethodGet)

	private := public.PathPrefix("").Subrouter()

	private.Use(api.auth.Authorize)

	privateUsers := private.PathPrefix("/users").Subrouter()
	privateUsers.Use(jsonResponse)

	privateUsers.HandleFunc("", api.Users().GetAll).Methods(http.MethodGet)

	return router
}

func jsonResponse(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r.Clone(r.Context()))
	})
}

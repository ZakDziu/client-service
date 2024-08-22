package api

import (
	"net/http"

	"api-gateway/pkg/api_builder"
	"api-gateway/pkg/authmiddleware"
	"api-gateway/pkg/config"

	"github.com/gorilla/mux"
)

type Server struct {
	*http.Server
}

type api struct {
	router     *mux.Router
	config     *config.ServerConfig
	auth       authmiddleware.AuthMiddleware
	apiBuilder api_builder.InternalAPI

	authHandler  *AuthHandler
	usersHandler *UsersHandler
	booksHandler *BooksHandler
}

func NewServer(
	config *config.ServerConfig,
	auth authmiddleware.AuthMiddleware,
	apiBuilder api_builder.InternalAPI,
) *Server {
	handler := newAPI(config, auth, apiBuilder)

	srv := &http.Server{
		Addr:              config.ServerPort,
		Handler:           handler,
		ReadHeaderTimeout: config.ReadTimeout.Duration,
	}

	return &Server{
		Server: srv,
	}
}

func newAPI(
	config *config.ServerConfig,
	auth authmiddleware.AuthMiddleware,
	apiBuilder api_builder.InternalAPI,
) *api {
	api := &api{
		config:     config,
		auth:       auth,
		apiBuilder: apiBuilder,
	}

	api.router = configureRouter(api)

	return api
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

func (a *api) Auth() *AuthHandler {
	if a.authHandler == nil {
		a.authHandler = NewAuthHandler(a)
	}

	return a.authHandler
}

func (a *api) Users() *UsersHandler {
	if a.usersHandler == nil {
		a.usersHandler = NewUsersHandler(a)
	}

	return a.usersHandler
}

func (a *api) Books() *BooksHandler {
	if a.booksHandler == nil {
		a.booksHandler = NewBooksHandler(a)
	}

	return a.booksHandler
}

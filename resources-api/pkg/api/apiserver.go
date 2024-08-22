package api

//nolint:revive
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"resources/pkg/config"
)

type Server struct {
	*http.Server
}

type api struct {
	router *gin.Engine
	config *config.ServerConfig

	usersHandler *UsersHandler
	booksHandler *BooksHandler
}

func NewServer(
	config *config.ServerConfig,
) *Server {
	handler := newAPI(config)

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
) *api {
	api := &api{
		config: config,
	}

	api.router = configureRouter(api)

	return api
}

//nolint:varnamelen
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding,"+
			"X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)

			return
		}

		c.Next()
	}
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
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

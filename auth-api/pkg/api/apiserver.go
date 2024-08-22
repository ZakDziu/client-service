package api

//nolint:revive
import (
	"net/http"

	"auth/pkg/authmiddleware"
	"auth/pkg/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*http.Server
}

type api struct {
	router *gin.Engine
	config *config.ServerConfig
	auth   authmiddleware.AuthMiddleware

	authHandler *AuthHandler
}

func NewServer(
	config *config.ServerConfig,
	auth authmiddleware.AuthMiddleware,
) *Server {
	handler := newAPI(config, auth)

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
) *api {
	api := &api{
		config: config,
		auth:   auth,
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

func (a *api) Auth() *AuthHandler {
	if a.authHandler == nil {
		a.authHandler = NewAuthHandler(a)
	}

	return a.authHandler
}

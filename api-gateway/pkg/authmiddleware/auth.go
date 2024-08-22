package authmiddleware

import (
	"net/http"
)

type AuthMiddleware interface {
	Authorize(next http.Handler) http.Handler
	ExtractToken(r *http.Request) string
}

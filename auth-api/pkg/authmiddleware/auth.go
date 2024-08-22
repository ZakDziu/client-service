package authmiddleware

import (
	"net/http"
)

type AuthMiddleware interface {
	CreateTokens() (*Tokens, error)
	Refresh(tokens Tokens) (*Tokens, error)
	ExtractToken(r *http.Request) string
	Validate(raw string) (*AccessClaims, error)
}

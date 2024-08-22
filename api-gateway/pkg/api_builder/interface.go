package api_builder

import (
	"api-gateway/pkg/authmiddleware"
	"api-gateway/pkg/model"
)

type InternalAPI interface {
	Auth() Auth
	Resources() Resources
}

type Auth interface {
	Token(request model.AuthUser) (*authmiddleware.Tokens, error)
	CheckToken(request model.Token) error
}

type Resources interface {
	Users(requestParams map[string]string) (*map[int]model.User, error)
	Books(requestParams map[string]string) (*map[int]model.Book, error)
}

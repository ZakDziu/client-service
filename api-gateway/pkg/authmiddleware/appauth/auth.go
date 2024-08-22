package appauth

import (
	"api-gateway/pkg/api_builder"
	"api-gateway/pkg/logger"
	"api-gateway/pkg/model"
	"net/http"
	"strings"
)

const StringsNumber = 2

type AuthMiddleware struct {
	apiBuilder api_builder.InternalAPI
}

func NewAuthMiddleware(apiBuilder api_builder.InternalAPI) *AuthMiddleware {
	var middleware = &AuthMiddleware{
		apiBuilder: apiBuilder,
	}

	return middleware
}

//nolint:varnamelen
func (m *AuthMiddleware) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := m.ExtractToken(r)

		if err := m.apiBuilder.Auth().CheckToken(model.Token{AccessToken: tokenString}); err != nil {
			logger.Errorf("Authorize.CheckToken %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *AuthMiddleware) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == StringsNumber {
		return strArr[1]
	}

	return ""
}

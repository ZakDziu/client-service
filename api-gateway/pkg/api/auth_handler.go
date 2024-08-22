package api

import (
	"encoding/json"
	"net/http"

	"api-gateway/pkg/model"
)

type AuthHandler struct {
	api *api
}

func NewAuthHandler(a *api) *AuthHandler {

	return &AuthHandler{
		api: a,
	}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var signInReq model.AuthUser

	err := json.NewDecoder(r.Body).Decode(&signInReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	tokens, err := h.api.apiBuilder.Auth().Token(signInReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(tokens)
}

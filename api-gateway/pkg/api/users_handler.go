package api

import (
	"encoding/json"
	"net/http"
)

type UsersHandler struct {
	api *api
}

func NewUsersHandler(a *api) *UsersHandler {

	return &UsersHandler{
		api: a,
	}
}

func (h *UsersHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.api.apiBuilder.Resources().Users(map[string]string{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(users)
}

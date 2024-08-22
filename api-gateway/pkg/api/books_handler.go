package api

import (
	"encoding/json"
	"net/http"
)

type BooksHandler struct {
	api *api
}

func NewBooksHandler(a *api) *BooksHandler {

	return &BooksHandler{
		api: a,
	}
}

func (h *BooksHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	books, err := h.api.apiBuilder.Resources().Books(map[string]string{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(books)
}

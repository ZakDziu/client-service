package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resources/pkg/model"
)

type BooksHandler struct {
	api *api
}

func NewBooksHandler(a *api) *BooksHandler {

	return &BooksHandler{
		api: a,
	}
}

func (h *BooksHandler) Books(c *gin.Context) {
	c.JSON(http.StatusOK, map[int]model.Book{
		1: {Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
		2: {Title: "Everyday routine", Author: "Harry"},
	})
}

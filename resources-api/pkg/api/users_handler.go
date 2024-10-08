package api

import (
	"net/http"

	"resources/pkg/model"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	api *api
}

func NewUsersHandler(a *api) *UsersHandler {

	return &UsersHandler{
		api: a,
	}
}

func (h *UsersHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, map[int]model.User{
		1: {Name: "John", Age: 30},
		2: {Name: "Jane", Age: 25},
	})
}

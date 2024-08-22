package api

import (
	"resources/pkg/model"

	"github.com/gin-gonic/gin"

	"net/http"
)

func configureRouter(api *api) *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	public := router.Group("/api/v1")

	public.GET("/users", api.Users().Users)
	public.GET("/books", api.Books().Books)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.ErrRecordNotFound)
	})

	return router
}

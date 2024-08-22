package api

import (
	"auth/pkg/model"

	"github.com/gin-gonic/gin"

	"net/http"
)

func configureRouter(api *api) *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	public := router.Group("/api/v1")

	public.POST("/token", api.Auth().Token)
	public.POST("/check_token", api.Auth().CheckToken)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.ErrRecordNotFound)
	})

	return router
}

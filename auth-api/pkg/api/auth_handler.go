package api

import (
	"net/http"

	"auth/pkg/logger"
	"auth/pkg/model"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	api *api
}

func NewAuthHandler(a *api) *AuthHandler {

	return &AuthHandler{
		api: a,
	}
}

func (h *AuthHandler) Token(c *gin.Context) {
	user := &model.AuthUser{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		logger.Errorf("Token.ShouldBindJSON", err)
		c.JSON(http.StatusBadRequest, model.ErrInvalidBody)

		return
	}
	tokens, err := h.api.auth.CreateTokens()
	if err != nil {
		logger.Errorf("Token.CreateTokens", err)
		c.JSON(http.StatusBadRequest, model.ErrUnhealthy)

		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (h *AuthHandler) CheckToken(c *gin.Context) {
	token := &model.Token{}
	err := c.ShouldBindJSON(&token)
	if err != nil {
		logger.Errorf("Token.ShouldBindJSON", err)
		c.JSON(http.StatusBadRequest, model.ErrInvalidBody)

		return
	}
	_, err = h.api.auth.Validate(token.AccessToken)
	if err != nil {
		logger.Errorf("Validate.Validate", err)
		c.JSON(http.StatusBadRequest, model.ErrUnauthorized)

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

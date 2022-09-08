package v1

import (
	"errors"
	"strconv"

	"github.com/IskanderA1/handly/iternal/service"
	"github.com/IskanderA1/handly/pkg/token"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services   *service.Services
	tokenMaker token.Maker
}

func NewHandler(services *service.Services, tokenMaker token.Maker) *Handler {
	return &Handler{
		services:   services,
		tokenMaker: tokenMaker,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAdminRoutes(v1)
	}
}

func parseUsernameFromPath(c *gin.Context, param string) (string, error) {
	idParam := c.Param(param)
	if idParam == "" {
		return "", errors.New("empty username param")
	}
	return idParam, nil
}

func parseIdFromPath(c *gin.Context, param string) (int64, error) {
	idParam := c.Param(param)
	if idParam == "" {
		return -1, errors.New("empty id param")
	}
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return -1, errors.New("invalid id param")
	}
	return id, nil
}

package http

import (
	"net/http"

	v1 "github.com/IskanderA1/handly/iternal/delivery/http/v1"
	"github.com/IskanderA1/handly/iternal/domain"
	"github.com/IskanderA1/handly/iternal/service"
	"github.com/IskanderA1/handly/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("event", validEvent)
	}

	handlerV1 := v1.NewHandler(h.services, h.tokenMaker)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}

var validEvent validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if event, ok := fieldLevel.Field().Interface().(string); ok {
		return domain.IsSupportedEvent(event)
	}
	return false
}

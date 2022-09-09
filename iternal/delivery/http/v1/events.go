package v1

import (
	"net/http"

	"github.com/IskanderA1/handly/iternal/domain"
	"github.com/IskanderA1/handly/iternal/service"
	"github.com/gin-gonic/gin"
)

type createEventInput struct {
	ProjectID int64  `json:"projectId" binding:"required"`
	Name      string `json:"name" binding:"required"`
	EventType string `json:"eventType" binding:"required,event"`
}

type updateEventInput struct {
	ID        int64  `json:"eventId" binding:"required"`
	Name      string `json:"name" binding:"required"`
	EventType string `json:"eventType" binding:"required,event"`
}

func (h *Handler) eventCreate(ctx *gin.Context) {
	var inp createEventInput
	if err := ctx.BindJSON(&inp); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.services.Events.Create(ctx.Request.Context(), service.CreateEventInput{
		ProjectID: inp.ProjectID,
		Name:      inp.Name,
		EventType: domain.EventType(inp.EventType),
	})
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) eventUpdate(ctx *gin.Context) {
	var inp updateEventInput
	if err := ctx.BindJSON(&inp); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.services.Events.Update(ctx.Request.Context(), service.UpdateEventInput{
		ID:        inp.ID,
		Name:      inp.Name,
		EventType: domain.EventType(inp.EventType),
	})
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) eventGetListByProjectId(ctx *gin.Context) {
	projectId, err := parseIdFromPath(ctx, "project-id")
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.services.Events.GetListByProjectId(ctx.Request.Context(), projectId)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) eventGetById(ctx *gin.Context) {
	id, err := parseIdFromPath(ctx, "id")
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.services.Events.GetById(ctx.Request.Context(), id)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) eventDeleteById(ctx *gin.Context) {
	id, err := parseIdFromPath(ctx, "id")
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Events.Delete(ctx.Request.Context(), id)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, successResponse{
		Success: true,
	})
}

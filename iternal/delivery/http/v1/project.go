package v1

import (
	"net/http"

	"github.com/IskanderA1/handly/iternal/service"
	"github.com/gin-gonic/gin"
)

type createInput struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) projectCreate(ctx *gin.Context) {
	var inp createInput
	if err := ctx.BindJSON(&inp); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.services.Projects.Create(ctx.Request.Context(), inp.Name)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) projectRefresh(ctx *gin.Context) {
	id, err := parseIdFromPath(ctx, "id")
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.services.Projects.RefreshTokens(ctx.Request.Context(), id)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) projectGetList(ctx *gin.Context) {
	var inp listAccountRequest

	if err := ctx.BindJSON(&inp); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.services.Projects.GetList(ctx.Request.Context(), service.ListInput{
		Limit:  inp.PageSize,
		Offset: (inp.PageID - 1) * inp.PageSize,
	})
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) projectGetById(ctx *gin.Context) {
	id, err := parseIdFromPath(ctx, "id")
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.services.Projects.GetById(ctx.Request.Context(), id)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) projectDeleteById(ctx *gin.Context) {
	id, err := parseIdFromPath(ctx, "id")
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Projects.Delete(ctx.Request.Context(), id)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, successResponse{
		Success: true,
	})
}

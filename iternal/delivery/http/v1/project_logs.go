package v1

import (
	"database/sql"
	"net/http"

	"github.com/IskanderA1/handly/iternal/domain"
	"github.com/IskanderA1/handly/iternal/service"
	"github.com/IskanderA1/handly/pkg/token"
	"github.com/gin-gonic/gin"
)

type userInput struct {
	ProjectAccountID string `json:"userID" binding:"required"`
	Name             string `json:"username" binding:"required"`
	Uuid             string `json:"deviceID" binding:"required"`
}

type logInput struct {
	EventName string `json:"eventName" binding:"required"`
	UserUUID  string `json:"deviceID" binding:"required"`
	Data      string `json:"data"`
}

func (h *Handler) initProjectLogsRoutes(api *gin.RouterGroup) {

	projectLogs := api.Group("/loger", h.projectMiddleware)
	{
		projectLogs.POST("/init-user", h.logInitUser)
		projectLogs.POST("/send-log", h.logSendLog)
	}
}

func (h *Handler) logInitUser(ctx *gin.Context) {
	var inp userInput
	if err := ctx.BindJSON(&inp); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.ProjectsLogs.InitUser(ctx.Request.Context(), service.UserInput{
		ProjectAccountID: inp.ProjectAccountID,
		Name:             inp.Name,
		Uuid:             inp.Uuid,
	})
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, successResponse{
		Success: true,
	})
}

func (h *Handler) logSendLog(ctx *gin.Context) {
	var inp logInput
	if err := ctx.BindJSON(&inp); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	payload := ctx.MustGet(domain.ProjectPayloadKey).(*token.ProjectPayload)

	err := h.services.ProjectsLogs.SendLog(ctx.Request.Context(), service.LogInput{
		ProjectId: payload.ProjectId,
		EventName: inp.EventName,
		UserUUID:  inp.UserUUID,
		Data: sql.NullString{
			String: inp.Data,
			Valid:  inp.Data != "",
		},
	})
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, successResponse{
		Success: true,
	})
}

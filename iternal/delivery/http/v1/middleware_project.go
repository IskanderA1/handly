package v1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/IskanderA1/handly/iternal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) projectMiddleware(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader(domain.AuthorizationHeaderKey)

	if len(authorizationHeader) == 0 {
		errorResponse(ctx, http.StatusUnauthorized, "authorization header is not provided")
		return
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		errorResponse(ctx, http.StatusUnauthorized, "invalid authorization header format")
		return
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != domain.AuthorizationTypeBearer {
		err := fmt.Errorf("unsupported authorization type %s", authorizationType)
		errorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	accessToken := fields[1]
	payload, err := h.projectTokenManger.VerifyToken(accessToken)
	if err != nil {
		errorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	ctx.Set(domain.ProjectPayloadKey, payload)
}

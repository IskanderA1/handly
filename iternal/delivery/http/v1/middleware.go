package v1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

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
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			errorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		accessToken := fields[1]
		payload, err := h.tokenMaker.VerifyAdminToken(accessToken)
		if err != nil {
			errorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

package v1

import (
	"database/sql"
	"net/http"

	"github.com/IskanderA1/handly/iternal/service"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initAdminRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins")
	{
		admins.POST("/sign-in", h.adminSignIn)
		admins.POST("/auth/refresh", h.adminRefresh)

		authenticated := admins.Group("/", h.authMiddleware)
		{
			authenticated.GET("/:username", h.adminGetByUsername)
			authenticated.DELETE("/:username", h.adminDeleteByUsername)
			authenticated.GET("/list", h.adminsGetList)
			authenticated.POST("/sign-up", h.adminSignUp)

			project := authenticated.Group("/projects")
			{
				project.POST("/create", h.projectCreate)
				project.POST("/refresh/:id", h.projectRefresh)
				project.GET("/list", h.projectGetList)
				project.GET("/:id", h.projectGetById)
				project.DELETE("/:id", h.projectDeleteById)
			}
		}
	}
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) adminSignIn(ctx *gin.Context) {
	var inp signInInput
	if err := ctx.BindJSON(&inp); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.services.Admins.SignIn(ctx.Request.Context(), service.AdminSingInInput{
		Username: inp.Username,
		Password: inp.Password,
	}, service.AdminConfig{
		UserAgent: ctx.Request.UserAgent(),
		ClientIp:  ctx.ClientIP(),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			errorResponse(ctx, http.StatusBadRequest, "admin not found")
		} else {
			errorResponse(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	ctx.JSON(http.StatusOK, res)
}

type refreshInput struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

func (h *Handler) adminRefresh(ctx *gin.Context) {
	var inp refreshInput

	if err := ctx.BindJSON(&inp); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.services.Admins.RefreshToken(ctx.Request.Context(), inp.RefreshToken)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) adminGetByUsername(ctx *gin.Context) {
	username, err := parseUsernameFromPath(ctx, "username")
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.services.Admins.GetByName(ctx.Request.Context(), username)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

type successResponse struct {
	Success bool `json:"success"`
}

func (h *Handler) adminDeleteByUsername(ctx *gin.Context) {
	username, err := parseUsernameFromPath(ctx, "username")
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Admins.Delete(ctx.Request.Context(), username)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, successResponse{
		Success: true,
	})
}

type listAccountRequest struct {
	PageID   int32 `json:"page_id" binding:"required,min=1"`
	PageSize int32 `json:"page_size" binding:"required,min=5,max=10"`
}

func (h *Handler) adminsGetList(ctx *gin.Context) {
	var inp listAccountRequest

	if err := ctx.BindJSON(&inp); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.services.Admins.GetList(ctx.Request.Context(), service.ListInput{
		Limit:  inp.PageSize,
		Offset: (inp.PageID - 1) * inp.PageSize,
	})
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

type signUpInput struct {
	Username string `json:"username" binding:"required"`
	FullName string `json:"fullname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) adminSignUp(ctx *gin.Context) {
	var inp signUpInput
	if err := ctx.BindJSON(&inp); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.services.Admins.SignUp(ctx.Request.Context(), service.AdminSignUpInput{
		Username: inp.Username,
		FullName: inp.FullName,
		Password: inp.Password,
	}, service.AdminConfig{
		UserAgent: ctx.Request.UserAgent(),
		ClientIp:  ctx.ClientIP(),
	})
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

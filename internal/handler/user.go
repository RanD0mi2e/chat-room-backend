package handler

import (
	"chat-app-backend-service/internal/service"
	"chat-app-backend-service/pkg/helper/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	*Handler
	userService service.UserService
}

func NewUserHandler(handler *Handler, userService service.UserService) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

func (h *UserHandler) GetUserByIdHandler(ctx *gin.Context) {
	var params struct {
		Id int `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		response.HandleError(ctx, http.StatusBadRequest, 1, err.Error(), nil)
		return
	}

	user, err := h.userService.GetUserInfoById(params.Id)
	h.logger.Info("GetUserByID", zap.Any("user", user))
	if err != nil {
		response.HandleError(ctx, http.StatusInternalServerError, 1, err.Error(), nil)
		return
	}
	response.HandleSuccess(ctx, user)
}

package server

import (
	"chat-app-backend-service/internal/handler"
	"chat-app-backend-service/pkg/helper/response"
	"chat-app-backend-service/pkg/log"

	"github.com/gin-gonic/gin"
)

func NewServerHttp(logger *log.Logger, handler *handler.UserHandler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// r.Use()
	r.GET("/", func(ctx *gin.Context) {
		response.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Nunu!",
		})
	})
	r.GET("/user", handler.GetUserByIdHandler)

	return r
}

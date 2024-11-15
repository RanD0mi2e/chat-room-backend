package handler

import "chat-app-backend-service/pkg/log"

type Handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

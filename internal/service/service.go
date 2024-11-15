package service

import "chat-app-backend-service/pkg/log"

type Service struct {
	logger *log.Logger
}

func NewService(logger *log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}

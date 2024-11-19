package service

import (
	"chat-app-backend-service/internal/repository"
	"chat-app-backend-service/pkg/log"
)

type Service struct {
	tm     repository.Transaction
	logger *log.Logger
}

func NewService(tm repository.Transaction, logger *log.Logger) *Service {
	return &Service{
		tm:     tm,
		logger: logger,
	}
}

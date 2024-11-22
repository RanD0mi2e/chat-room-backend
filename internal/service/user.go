package service

import (
	"chat-app-backend-service/internal/model"
	"chat-app-backend-service/internal/repository"
)

type userService struct {
	*Service
	userRepo repository.UserRepo
}

type UserService interface {
	GetUserInfoById(id int) (*model.User, error)
}

func NewUserService(service *Service, userRepo repository.UserRepo) UserService {
	return &userService{
		Service:  service,
		userRepo: userRepo,
	}
}

func (s *userService) GetUserInfoById(id int) (*model.User, error) {
	return s.userRepo.GetUserById(id)
}

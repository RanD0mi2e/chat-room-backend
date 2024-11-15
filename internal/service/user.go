package service

import "chat-app-backend-service/internal/repository"

type userService struct {
	*Service
	userRepo repository.UserRepo
}

type UserService interface {
	GetUserInfoById(id int) (*repository.User, error)
}

func NewUserService(service *Service, userRepo repository.UserRepo) UserService {
	return &userService{
		Service:  service,
		userRepo: userRepo,
	}
}

func (s *userService) GetUserInfoById(id int) (*repository.User, error) {
	return s.userRepo.GetUserById(id)
}

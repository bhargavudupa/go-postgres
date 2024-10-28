package service

import (
	"go-postgres-test-1/model"
	"go-postgres-test-1/repository"
)

type UserService interface {
	CreateUser(user model.NewUserRequest) (err model.Error)
	GetAllUsers() (users []model.User, err model.Error)
	GetUser(userId uint) (user model.User, err model.Error)
	UpdateUser(userId uint, user model.UpdateUserRequest) (err model.Error)
	DeleteUser(userId uint) (err model.Error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (service *userService) CreateUser(user model.NewUserRequest) (err model.Error) {
	return service.repo.CreateUser(user.Username, user.Email, user.Password)
}

func (service *userService) GetAllUsers() (users []model.User, err model.Error) {
	return service.repo.GetAllUsers()
}

func (service *userService) GetUser(userId uint) (user model.User, err model.Error) {
	return service.repo.GetUser(userId)
}

func (service *userService) UpdateUser(userId uint, user model.UpdateUserRequest) (err model.Error) {
	return service.repo.UpdateUser(userId, user.Email, user.Password)
}

func (service *userService) DeleteUser(userId uint) (err model.Error) {
	return service.repo.DeleteUser(userId)
}

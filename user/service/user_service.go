package service

import (
	"LatihanGolang/entity"
	userGorm "LatihanGolang/user/repository"
	"context"
	"fmt"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

type userService struct {
	userRepo userGorm.IUserRepository
}

func NewUserService(userRepo userGorm.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	createdUser, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("error creating user: %v", err)
	}

	return createdUser, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return entity.User{}, fmt.Errorf("user not found: %v", err)
	}

	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	updatedUser, err := s.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed updated user: %v", err)
	}

	return updatedUser, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int) error {
	err := s.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed deleted user: %v", err)
	}

	return nil
}

func (s *userService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	users, err := s.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed get all users: %v", err)
	}

	return users, nil
}

package service

import (
	"context"

	"github.com/YogaRP/finansial/user-service/internal/dto"
	"github.com/YogaRP/finansial/user-service/internal/repository"
	"github.com/google/uuid"
)

type UserServiceInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*dto.UserResponse, error)
	GetUserByID(ctx context.Context, uuid uuid.UUID) (*dto.UserResponse, error)
}

type UserService struct {
	userRepository repository.UserRepository
}

// GetUserByEmail implements [UserServiceInterface].
func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*dto.UserResponse, error) {
	return u.userRepository.GetUserByEmail(ctx, email)
}

// GetUserByID implements [UserServiceInterface].
func (u *UserService) GetUserByID(ctx context.Context, uuid uuid.UUID) (*dto.UserResponse, error) {
	return u.userRepository.GetUserByID(ctx, uuid)
}

func NewUserService(userRepository repository.UserRepository) UserServiceInterface {
	return &UserService{
		userRepository,
	}
}

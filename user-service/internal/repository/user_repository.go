package repository

import (
	"context"

	"github.com/YogaRP/finansial/user-service/internal/dto"
	"github.com/YogaRP/finansial/user-service/internal/model"
	"github.com/YogaRP/finansial/user-service/internal/pkg/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*dto.UserResponse, error)
	GetUserByID(ctx context.Context, uuid uuid.UUID) (*dto.UserResponse, error)
}

type userRepository struct {
	db *gorm.DB
}

// GetUserByEmail implements [UserRepository].
func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*dto.UserResponse, error) {
	var user model.User
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		logger.Errorf("[UserRepository] GetUserByEmail - 1", err)
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (u *userRepository) GetUserByID(ctx context.Context, uuid uuid.UUID) (*dto.UserResponse, error) {
	var user model.User
	if err := u.db.WithContext(ctx).Where("id = ?", uuid).First(&user).Error; err != nil {
		logger.Errorf("[UserRepository] GetUserById - 1", err)
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

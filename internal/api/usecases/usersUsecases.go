package usecases

import (
	"context"
	"kodeTestTask/internal/api/models"
	. "kodeTestTask/internal/api/repositories"
)

type UsersUsecase interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, userID int) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, userID int) error
}

type usersUsecase struct {
	usersRepository UsersRepository
}

func NewUsersUsecase(repository UsersRepository) UsersUsecase {
	return &usersUsecase{repository}
}

func (uc *usersUsecase) CreateUser(ctx context.Context, user *models.User) error {
	return uc.usersRepository.Create(ctx, user)
}

func (uc *usersUsecase) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	return uc.usersRepository.GetByID(ctx, userID)
}

func (uc *usersUsecase) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return uc.usersRepository.GetByUsername(ctx, username)
}

func (uc *usersUsecase) UpdateUser(ctx context.Context, user *models.User) error {
	return uc.usersRepository.Update(ctx, user)
}

func (uc *usersUsecase) DeleteUser(ctx context.Context, userID int) error {
	return uc.usersRepository.Delete(ctx, userID)
}

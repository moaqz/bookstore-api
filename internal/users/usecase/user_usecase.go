package usecase

import (
	"context"

	"github.com/techwithmat/bookery-api/internal/domain"
	"github.com/techwithmat/bookery-api/pkg/utils/hash"
	v "github.com/techwithmat/bookery-api/pkg/utils/validation"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) RegisterUser(ctx context.Context, user *domain.SignUpRequest) error {
	validationErrors := v.ValidateStruct(user)

	if validationErrors != nil {
		return validationErrors
	}

	hashedPassword, err := hash.HashPassword(user.Password)

	if err != nil {
		return err
	}

	// change the password for the hashed password
	user.Password = hashedPassword

	err = u.userRepo.RegisterUser(ctx, user)

	return err
}

func (u *userUseCase) GetUserByID(ctx context.Context, id int64) (*domain.GetUserResponse, error) {
	return &domain.GetUserResponse{}, nil
}

func (u *userUseCase) LoginUser(ctx context.Context, user *domain.LoginRequest) error {
	return nil
}

func (u *userUseCase) DeleteUser(ctx context.Context, email string) error {
	return nil
}

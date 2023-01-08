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
	user, err := u.userRepo.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUseCase) LoginUser(ctx context.Context, user *domain.LoginRequest) error {
	validationErrors := v.ValidateStruct(user)

	if validationErrors != nil {
		return validationErrors
	}

	existingUser, err := u.userRepo.LoginUser(ctx, user)

	if err != nil {
		return err
	}

	err = hash.PasswordMatch(existingUser.Password, user.Password)

	// nil means it is a match
	return err
}

func (u *userUseCase) DeleteUser(ctx context.Context, id int64) error {
	err := u.userRepo.DeleteUser(ctx, id)

	return err
}

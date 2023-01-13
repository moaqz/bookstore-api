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

	err = u.userRepo.Insert(ctx, user)

	return err
}

func (u *userUseCase) GetUserByID(ctx context.Context, id int64) (*domain.GetUserResponse, error) {
	user, err := u.userRepo.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUseCase) GetUser(ctx context.Context, user *domain.LoginRequest) (*domain.User, error) {
	validationErrors := v.ValidateStruct(user)

	if validationErrors != nil {
		return nil, validationErrors
	}

	existingUser, err := u.userRepo.FindByEmail(ctx, user.Email)

	if err != nil {
		return nil, err
	}

	err = hash.PasswordMatch(existingUser.Password, user.Password)

	// nil means it is a match
	return existingUser, err
}

func (u *userUseCase) DeleteUser(ctx context.Context, user *domain.UnregisterRequest) error {
	validationErrors := v.ValidateStruct(user)

	if validationErrors != nil {
		return validationErrors
	}

	existingUser, err := u.userRepo.FindByEmail(ctx, user.Email)

	if err != nil {
		return err
	}

	err = hash.PasswordMatch(existingUser.Password, user.Password)

	if err != nil {
		return err
	}

	err = u.userRepo.Delete(ctx, user.Email)

	return err
}

package usecase

import (
	"context"

	"github.com/techwithmat/bookstore-api/internal/domain"
	"github.com/techwithmat/bookstore-api/pkg/utils/hash"
	"github.com/techwithmat/bookstore-api/pkg/utils/jwtToken"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) RegisterUser(ctx context.Context, user *domain.SignUpRequest) (*domain.TokenResponse, error) {
	hashedPassword, err := hash.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	// change the password for the hashed password
	user.Password = hashedPassword

	// Insert the user to the database
	id, err := u.userRepo.Insert(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := jwtToken.GenerateJWT(id, false)

	if err != nil {
		return nil, err
	}

	return &domain.TokenResponse{
		Message: "User created",
		Token:   token,
	}, nil
}

func (u *userUseCase) GetUserByID(ctx context.Context, id int64) (*domain.GetUserResponse, error) {
	user, err := u.userRepo.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUseCase) LoginUser(ctx context.Context, user *domain.LoginRequest) (*domain.TokenResponse, error) {
	existingUser, err := u.userRepo.FindByEmail(ctx, user.Email)

	if err != nil {
		return nil, err
	}

	// err equal to nil means it is a match
	err = hash.PasswordMatch(existingUser.Password, user.Password)

	if err != nil {
		return nil, err
	}

	token, err := jwtToken.GenerateJWT(existingUser.ID, existingUser.IsStaff)
	if err != nil {
		return nil, err
	}

	return &domain.TokenResponse{
		Message: "Login succesful",
		Token:   token,
	}, nil
}

func (u *userUseCase) DeleteUser(ctx context.Context, id int64) error {
	err := u.userRepo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

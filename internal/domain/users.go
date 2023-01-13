package domain

import (
	"context"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsStaff  bool   `json:"is_staff" db:"is_staff"`
}

type SignUpRequest struct {
	Email                string `json:"email" form:"email" validate:"required,email"`
	Password             string `json:"password" form:"password" validate:"required,eqfield=PasswordConfirmation,min=8,max=28"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" validate:"required,min=8,max=28"`
	HashedPassword       string `json:"hashed_password" swaggerignore:"true"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=28"`
}

type TokenResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type GetUserResponse struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	IsStaff bool   `json:"is_staff" db:"is_staff"`
}

type UnregisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=28"`
}

// UserUsecase represent the User's usecases
type UserUseCase interface {
	GetUserByID(ctx context.Context, id int64) (*GetUserResponse, error)
	RegisterUser(ctx context.Context, user *SignUpRequest) error
	GetUser(ctx context.Context, user *LoginRequest) (*User, error)
	DeleteUser(ctx context.Context, user *UnregisterRequest) error
}

// UserRepository represent the User's repository contract
type UserRepository interface {
	Insert(ctx context.Context, user *SignUpRequest) error
	FindById(ctx context.Context, id int64) (*GetUserResponse, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Delete(ctx context.Context, email string) error
}

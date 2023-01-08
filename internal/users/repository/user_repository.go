package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/techwithmat/bookery-api/internal/domain"
)

type usersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) domain.UserRepository {
	return &usersRepo{
		db: db,
	}
}

func (r *usersRepo) RegisterUser(ctx context.Context, user *domain.SignUpRequest) error {
	_, err := r.db.ExecContext(ctx, registerUser, user.Email, user.Password, false)

	return err
}

func (r *usersRepo) GetUserByID(ctx context.Context, id int64) (*domain.GetUserResponse, error) {
	return &domain.GetUserResponse{}, nil
}

func (r *usersRepo) LoginUser(ctx context.Context, user *domain.LoginRequest) (*domain.User, error) {
	return &domain.User{}, nil
}

func (r *usersRepo) DeleteUser(ctx context.Context, email string) error {
	return nil
}

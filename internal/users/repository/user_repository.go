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
	var user domain.GetUserResponse

	err := r.db.GetContext(ctx, &user, getUserById, id)

	if err != nil {
		return nil, err
	}

	user.ID = id

	return &user, err
}

func (r *usersRepo) LoginUser(ctx context.Context, user *domain.LoginRequest) (*domain.User, error) {
	var existingUserbyEmail domain.User

	err := r.db.GetContext(ctx, &existingUserbyEmail, loginUser, user.Email)

	if err != nil {
		return nil, err
	}

	return &existingUserbyEmail, nil
}

func (r *usersRepo) DeleteUser(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, deleteUser, id)

	if err != nil {
		return err
	}

	return nil
}

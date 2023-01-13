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

func (r *usersRepo) Insert(ctx context.Context, user *domain.SignUpRequest) error {
	_, err := r.db.ExecContext(ctx, InsertUserQuery, user.Email, user.Password, false)

	return err
}

func (r *usersRepo) FindById(ctx context.Context, id int64) (*domain.GetUserResponse, error) {
	var user domain.GetUserResponse

	err := r.db.GetContext(ctx, &user, FindUserByIdQuery, id)

	if err != nil {
		return nil, err
	}

	user.ID = id

	return &user, err
}

func (r *usersRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	err := r.db.GetContext(ctx, &user, FindOneUserQuery, email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *usersRepo) Delete(ctx context.Context, email string) error {
	_, err := r.db.ExecContext(ctx, DeleteUserQuery, email)

	if err != nil {
		return err
	}

	return nil
}

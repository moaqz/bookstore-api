package repository

import (
	"context"
	"database/sql"

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

func (r *usersRepo) Insert(ctx context.Context, user *domain.SignUpRequest) (int64, error) {
	var id int64

	err := r.db.QueryRowContext(ctx, InsertUserQuery, user.Email, user.Password, false).Scan(&id)

	return id, err
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

	err := r.db.GetContext(ctx, &user, FindByEmailUserQuery, email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *usersRepo) Delete(ctx context.Context, userId int64) error {
	result, err := r.db.ExecContext(ctx, DeleteUserQuery, userId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return sql.ErrNoRows
	}

	return nil
}

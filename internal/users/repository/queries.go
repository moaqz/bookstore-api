package repository

const (
	registerUser = `
	INSERT INTO
    users (email, password, is_staff)
	VALUES
    ($1, $2, $3);
	`
)
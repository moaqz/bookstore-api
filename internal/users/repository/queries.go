package repository

const (
	InsertUserQuery = `
	INSERT INTO
    users (email, password, is_staff)
	VALUES
    ($1, $2, $3)
	RETURNING id
	`

	FindUserByIdQuery = `
	SELECT 
		email, 
		is_staff 
	FROM 
		users 
	WHERE 
		id = $1;
	`

	FindByEmailUserQuery = `
	SELECT 
		id,
		email,
		password,
		is_staff
	FROM
		users
	WHERE
		email = $1;
	`

	DeleteUserQuery = `
	DELETE FROM 
		users
	WHERE
		id = $1;
	`
)

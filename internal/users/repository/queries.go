package repository

const (
	registerUser = `
	INSERT INTO
    users (email, password, is_staff)
	VALUES
    ($1, $2, $3);
	`

	getUserById = `
	SELECT 
		email, 
		is_staff 
	FROM 
		users 
	WHERE 
		id = $1;
	`

	getUser = `
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

	deleteUser = `
	DELETE FROM 
		users
	WHERE
		id = $1;
	`
)

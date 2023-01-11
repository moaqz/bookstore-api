package repository

const (
	getBookByIdQuery = `
	SELECT
    id,
    title,
    subtitle,
    about_the_book,
    page_count,
    price,
    image,
    language,
    author_name,
    author_avatar,
    category_id
	FROM
    books
	WHERE
    id = $1;
	`

	getBookByCategoryQuery = `
	SELECT
    b.id,
    b.title,
    b.subtitle,
    b.image,
    b.author_name
	FROM
    books b
    INNER JOIN categories c ON b.category_id = c.id
	WHERE 
		c.name = $1
	LIMIT $2 OFFSET $3;
	`

	getBooksQuery = `
	SELECT
    id, title, subtitle, image, author_name
	FROM
    books b
	LIMIT $1 OFFSET $2;
	`

	InsertBookQuery = `
	INSERT INTO 
		books
		(title, subtitle, about_the_book, page_count, price, image, language, author_name, author_avatar, category_id)
	VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id;
	`

	DeleteBookQuery = `
	DELETE FROM
		books
	WHERE
		id = $1;
	`
)

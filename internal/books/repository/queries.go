package repository

const (
	getBookById = `
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

	getBookByCategory = `
	SELECT
    b.id,
    b.title,
    b.subtitle,
    b.image,
    b.author_name
	FROM
    books b
    INNER JOIN categories c ON b.category_id = c.id
	WHERE c.name = $1;
	`

	getBooks = `
	SELECT
    id, title, subtitle, image, author_name
	FROM
    books b;
	`
)

package domain

import "database/sql"

type Author struct {
	ID        int            `db:"id"`
	Name      string         `db:"name"`
	Avatar    string         `db:"avatar"`
	Github    sql.NullString `db:"github_name"`
	Instagram sql.NullString `db:"instagram_name"`
	Linkedin  sql.NullString `db:"linkedin_name"`
	Twitter   sql.NullString `db:"twitter_name"`
}

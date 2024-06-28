package repository

import "database/sql"

type RepoHolder struct {
	Db *sql.DB
}

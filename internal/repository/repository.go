package repository

import "database/sql"

// const (
// 	pg_err_unique_violation = "23505"
// )

type RepoHolder struct {
	Db *sql.DB
}

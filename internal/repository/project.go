package repository

import (
	"database/sql"
	"log"

	"github.com/GeekchanskiY/cv_builder/internal/schemas"
)

type ProjectRepository struct {
	db *sql.DB
}

func CreateProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (repo *ProjectRepository) Create(schema schemas.Project) (int, error) {
	new_id := 0
	err := repo.db.QueryRow("INSERT INTO projects(name, description, company_id) VALUES($1, $2, $3) RETURNING id",
		schema.Name, schema.Description, schema.CompanyId).Scan(&new_id)
	if err != nil {
		log.Println("Error creating project in repository: ", err)
		return 0, err
	}

	return int(new_id), nil
}

func (repo *ProjectRepository) Update(schema schemas.Project) error {
	_, err := repo.db.Exec("UPDATE projects SET name = $1, description = $2, company_id = $3 WHERE id = $4",
		schema.Name, schema.Description, schema.CompanyId, schema.Id)
	return err
}

func (repo *ProjectRepository) Delete(schema schemas.Project) error {
	_, err := repo.db.Exec("DELETE FROM projects WHERE id = $1", schema.Id)
	return err
}

func (repo *ProjectRepository) GetAll() (schemes []schemas.Project, err error) {
	rows, err := repo.db.Query("SELECT id, name, description, company_id FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schema schemas.Project
		if err := rows.Scan(&schema.Id, &schema.Name, &schema.Description, &schema.CompanyId); err != nil {
			return nil, err
		}
		schemes = append(schemes, schema)
	}

	if err := rows.Err(); err != nil {
		// Here's not nil, err. I'm not sure why, but:
		// https://go.dev/doc/database/querying
		// documentation has the same issue

		return schemes, err
	}

	return schemes, nil
}

func (repo *ProjectRepository) Get(id int) (schema schemas.Project, err error) {
	row := repo.db.QueryRow("SELECT id, name, description, company_id FROM projects WHERE id = $1", id)
	err = row.Scan(&schema.Id, &schema.Name, &schema.Description, &schema.CompanyId)
	return schema, err
}

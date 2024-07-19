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
	err := repo.db.QueryRow("INSERT INTO projects(name, description) VALUES($1, $2) RETURNING id",
		schema.Name, schema.Description).Scan(&new_id)
	if err != nil {
		log.Println("Error creating project in repository: ", err)
		return 0, err
	}

	return int(new_id), nil
}

func (repo *ProjectRepository) Update(schema schemas.Project) error {
	_, err := repo.db.Exec("UPDATE projects SET name = $1, description = $2 WHERE id = $3",
		schema.Name, schema.Description, schema.Id)
	return err
}

func (repo *ProjectRepository) Delete(schema schemas.Project) error {
	_, err := repo.db.Exec("DELETE FROM projects WHERE id = $1", schema.Id)
	return err
}

func (repo *ProjectRepository) GetAll() (schemes []schemas.Project, err error) {
	rows, err := repo.db.Query("SELECT id, name, description FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schema schemas.Project
		if err := rows.Scan(&schema.Id, &schema.Name, &schema.Description); err != nil {
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
	row := repo.db.QueryRow("SELECT id, name, description FROM projects WHERE id = $1", id)
	err = row.Scan(&schema.Id, &schema.Name, &schema.Description)
	return schema, err
}

func (repo *ProjectRepository) GetDomains(id int) (schemes []schemas.ProjectDomain, err error) {
	q := `SELECT id, project_id, domain_id, comments 
	FROM project_domains
	WHERE project_id = $1`

	rows, err := repo.db.Query(q, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var schema schemas.ProjectDomain
		err = rows.Scan(
			&schema.Id,
			&schema.ProjectId,
			&schema.DomainId,
			&schema.Comments,
		)
		if err != nil {
			return nil, err
		}
		schemes = append(schemes, schema)
	}

	if err := rows.Err(); err != nil {

		return schemes, err
	}

	if err := rows.Close(); err != nil {
		return schemes, err
	}
	return schemes, nil
}

func (repo *ProjectRepository) GetAllDomains() (schemes []schemas.ProjectDomain, err error) {
	q := `SELECT id, project_id, domain_id, comments 
	FROM project_domains`

	rows, err := repo.db.Query(q)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var schema schemas.ProjectDomain
		err = rows.Scan(
			&schema.Id,
			&schema.ProjectId,
			&schema.DomainId,
			&schema.Comments,
		)
		if err != nil {
			return nil, err
		}
		schemes = append(schemes, schema)
	}

	if err := rows.Err(); err != nil {

		return schemes, err
	}

	if err := rows.Close(); err != nil {
		return schemes, err
	}
	return schemes, nil
}

func (repo *ProjectRepository) CreateDomains(schema schemas.ProjectDomain) (new_id int, err error) {
	q := `INSERT INTO project_domains(project_id, domain_id, comments) VALUES($1, $2, $3) RETURNING id`

	new_id = 0
	err = repo.db.QueryRow(
		q, schema.ProjectId, schema.DomainId, schema.Comments,
	).Scan(&new_id)
	if err != nil {
		log.Println("Error creating projectDomain in repository: ", err)
		return 0, err
	}

	return new_id, nil
}

func (repo *ProjectRepository) UpdateDomains(schema schemas.ProjectDomain) error {
	q := `UPDATE project_domains SET project_id = $1, domain_id = $2, comments = $3 WHERE id = $4`
	_, err := repo.db.Exec(
		q,
		schema.ProjectId,
		schema.DomainId,
		schema.Comments,
		schema.Id,
	)
	return err
}

func (repo *ProjectRepository) DeleteDomains(schema schemas.ProjectDomain) error {
	q := `DELETE FROM project_domains WHERE id = $1`
	_, err := repo.db.Exec(q, schema.Id)
	return err
}

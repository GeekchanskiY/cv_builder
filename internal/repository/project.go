package repository

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
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
	newId := 0
	err := repo.db.QueryRow("INSERT INTO projects(name, description) VALUES($1, $2) RETURNING id",
		schema.Name, schema.Description).Scan(&newId)
	if err != nil {
		log.Println("Error creating project in repository: ", err)
		return 0, err
	}

	return newId, nil
}

func (repo *ProjectRepository) Count() (res int, err error) {
	q := `SELECT COUNT(*) FROM projects`

	err = repo.db.QueryRow(q).Scan(&res)

	if err != nil {
		log.Println("Error getting amount of projects: ", err)

		return 0, err
	}

	return res, nil
}

func (repo *ProjectRepository) CreateIfNotExists(schema schemas.Project) (created bool, err error) {
	q := `INSERT INTO projects(name, description) 
	SELECT CAST($1 AS VARCHAR), $2
	WHERE 
	    NOT EXISTS (SELECT 1 FROM projects WHERE name = $1)`

	r, err := repo.db.Exec(q, schema.Name, schema.Description)

	if err != nil {
		log.Println("Error creating project: ", err)

		return false, err
	}

	if i, _ := r.RowsAffected(); i != 0 {
		return true, nil
	}

	return false, nil
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
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing rows: ", err)
		}
	}(rows)

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

func (repo *ProjectRepository) GetMicroservicesByDomains(domainIds []int) (schemes []schemas.Project, err error) {
	q := `SELECT p.id, p.name, p.description
	FROM projects p
	JOIN project_services ps ON ps.project_id = p.id
	JOIN project_domains pd ON pd.project_id = p.id
	WHERE pd.domain_id = ANY ($1::int[])
	ORDER BY (
		SELECT COUNT(*) from project_domains pdd 
		                where pdd.project_id = p.id
		                and pdd.domain_id = ANY ($1::int[])
	)
	`
	if len(domainIds) == 0 {
		return nil, errors.New("domainIds cant be empty")
	}

	rows, err := repo.db.Query(q, pq.Array(domainIds))

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing rows: ", err)
		}
	}(rows)

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

func (repo *ProjectRepository) GetAllDomainsReadable() (schemes []schemas.ProjectDomainReadable, err error) {
	q := `SELECT p.name, d.name, comments 
	FROM project_domains pd
	JOIN projects p ON pd.project_id = p.id
	JOIN domains d on d.id = pd.domain_id`

	rows, err := repo.db.Query(q)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var schema schemas.ProjectDomainReadable
		err = rows.Scan(
			&schema.ProjectName,
			&schema.DomainName,
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

func (repo *ProjectRepository) CreateDomains(schema schemas.ProjectDomain) (newId int, err error) {
	q := `INSERT INTO project_domains(project_id, domain_id, comments) VALUES($1, $2, $3) RETURNING id`

	err = repo.db.QueryRow(
		q, schema.ProjectId, schema.DomainId, schema.Comments,
	).Scan(&newId)
	if err != nil {
		log.Println("Error creating projectDomain in repository: ", err)
		return 0, err
	}

	return newId, nil
}

func (repo *ProjectRepository) CreateDomainsIfNotExists(schema schemas.ProjectDomainReadable) (created bool, err error) {
	q := `INSERT INTO project_domains(project_id, domain_id, comments) 
	SELECT p.id, d.id, $3
	FROM projects p
	JOIN domains d ON d.name = $2::text
	WHERE 
	    p.name = $1::text 
	    AND NOT EXISTS (
		SELECT 1 
		FROM project_domains pd
		WHERE pd.project_id = p.id
		AND pd.domain_id = d.id 
		);`

	r, err := repo.db.Exec(q, schema.ProjectName, schema.DomainName, schema.Comments)

	if err != nil {
		log.Println("Error creating project domain: ", err)

		return false, err
	}

	if i, _ := r.RowsAffected(); i != 0 {
		return true, nil
	}

	return false, nil
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

func (repo *ProjectRepository) GetServices(id int) (schemes []schemas.ProjectService, err error) {
	q := `SELECT id, project_id, name, description
	FROM project_services
	WHERE project_id = $1`

	rows, err := repo.db.Query(q, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var schema schemas.ProjectService
		err = rows.Scan(
			&schema.Id,
			&schema.ProjectId,
			&schema.Name,
			&schema.Description,
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

func (repo *ProjectRepository) GetAllServices() (schemes []schemas.ProjectService, err error) {
	q := `SELECT id, project_id, name, description 
	FROM project_services`

	rows, err := repo.db.Query(q)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var schema schemas.ProjectService
		err = rows.Scan(
			&schema.Id,
			&schema.ProjectId,
			&schema.Name,
			&schema.Description,
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

func (repo *ProjectRepository) GetAllServicesReadable() (schemes []schemas.ProjectServiceReadable, err error) {
	q := `SELECT p.name, ps.name, ps.description 
	FROM project_services ps
	JOIN projects p ON ps.project_id = p.id`

	rows, err := repo.db.Query(q)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var schema schemas.ProjectServiceReadable
		err = rows.Scan(
			&schema.ProjectName,
			&schema.Name,
			&schema.Description,
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

func (repo *ProjectRepository) CreateService(schema schemas.ProjectService) (newId int, err error) {
	q := `INSERT INTO project_services(project_id, name, description) VALUES($1, $2, $3) RETURNING id`

	err = repo.db.QueryRow(
		q, schema.ProjectId, schema.Name, schema.Description,
	).Scan(&newId)
	if err != nil {
		log.Println("Error creating projectService in repository: ", err)
		return 0, err
	}

	return newId, nil
}

func (repo *ProjectRepository) CreateServiceIfNotExists(schema schemas.ProjectServiceReadable) (created bool, err error) {
	q := `INSERT INTO project_services(project_id, name, description) 
	SELECT p.id, $2::text, $3::text
	FROM projects p
	WHERE 
	    p.name = $1::text 
	    AND NOT EXISTS (
		SELECT 1 
		FROM project_services ps
		WHERE ps.project_id = p.id
		AND ps.name = $2
		);`
	r, err := repo.db.Exec(q, schema.ProjectName, schema.Name, schema.Description)

	if err != nil {
		log.Println("Error creating project service: ", err)

		return false, err
	}
	if i, _ := r.RowsAffected(); i != 0 {
		return true, nil
	}

	return false, nil
}

func (repo *ProjectRepository) UpdateService(schema schemas.ProjectService) error {
	q := `UPDATE project_services SET project_id = $1, name = $2, description = $3 WHERE id = $4`
	_, err := repo.db.Exec(
		q,
		schema.ProjectId,
		schema.Name,
		schema.Description,
		schema.Id,
	)
	return err
}

func (repo *ProjectRepository) DeleteService(schema schemas.ProjectService) error {
	q := `DELETE FROM project_services WHERE id = $1`
	_, err := repo.db.Exec(q, schema.Id)
	return err
}

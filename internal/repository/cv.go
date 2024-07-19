package repository

import (
	"database/sql"
	"log"

	"github.com/GeekchanskiY/cv_builder/internal/schemas"
)

type CVRepository struct {
	db *sql.DB
}

func CreateCVRepository(db *sql.DB) *CVRepository {
	return &CVRepository{
		db: db,
	}
}

func (repo *CVRepository) Create(schema schemas.CV) (int, error) {
	q := `INSERT INTO cvs(vacancy_id, employee_id, is_real) VALUES($1, $2, $3) RETURNING id`
	new_id := 0
	err := repo.db.QueryRow(q, schema.VacancyId, schema.EmployeeId, schema.IsReal).Scan(&new_id)
	if err != nil {
		log.Println("Error creating cv in cv repository: ", err)
		return 0, err
	}

	return new_id, nil
}

func (repo *CVRepository) Update(schema schemas.CV) error {
	q := `UPDATE cvs SET vacancy_id = $1, employee_id = $2, is_real = $3 WHERE id = $4`
	_, err := repo.db.Exec(q, schema.VacancyId, schema.EmployeeId, schema.IsReal, schema.Id)
	return err
}

func (repo *CVRepository) Delete(schema schemas.CV) error {
	_, err := repo.db.Exec("DELETE FROM cvs WHERE id = $1", schema.Id)
	return err
}

func (repo *CVRepository) GetAll() (schemes []schemas.CV, err error) {
	q := `SELECT id, vacancy_id, employee_id, is_real FROM cvs`
	rows, err := repo.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schema schemas.CV
		err = rows.Scan(&schema.Id, &schema.VacancyId, &schema.EmployeeId, &schema.IsReal)
		if err != nil {
			return nil, err
		}
		schemes = append(schemes, schema)
	}

	if err = rows.Err(); err != nil {

		return schemes, err
	}

	return schemes, nil
}

func (repo *CVRepository) Get(id int) (schema schemas.CV, err error) {
	q := `SELECT id, vacancy_id, employee_id, is_real FROM cvs WHERE id = $1`
	row := repo.db.QueryRow(q, id)
	err = row.Scan(&schema.Id, &schema.VacancyId, &schema.EmployeeId, &schema.IsReal)
	return schema, err
}

func (repo *CVRepository) GetProjects(id int) (schemes []schemas.CVProject, err error) {
	q := `SELECT id, cv_id, project_id, company_id, end_time, start_time
	FROM cv_projects
	WHERE cv_id = $1`

	rows, err := repo.db.Query(q, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var schema schemas.CVProject
		err = rows.Scan(
			&schema.Id,
			&schema.CVId,
			&schema.ProjectId,
			&schema.CompanyId,
			&schema.EndTime,
			&schema.StartTime,
		)
		if err != nil {
			return nil, err
		}
		schemes = append(schemes, schema)
	}

	if err := rows.Err(); err != nil {

		return schemes, err
	}

	return schemes, nil
}

func (repo *CVRepository) GetAllProjects() (schemes []schemas.CVProject, err error) {
	q := `SELECT id, cv_id, project_id, company_id, end_time, start_time
	FROM cv_projects`

	rows, err := repo.db.Query(q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var schema schemas.CVProject
		err = rows.Scan(
			&schema.Id,
			&schema.CVId,
			&schema.ProjectId,
			&schema.CompanyId,
			&schema.EndTime,
			&schema.StartTime,
		)
		if err != nil {
			return nil, err
		}
		schemes = append(schemes, schema)
	}

	if err := rows.Err(); err != nil {

		return schemes, err
	}

	return schemes, nil
}

func (repo *CVRepository) CreateProject(schema schemas.CVProject) (new_id int, err error) {
	q := `INSERT INTO cv_projects(cv_id, project_id, company_id, end_time, start_time) VALUES($1, $2, $3, $4, $5) RETURNING id`

	new_id = 0
	err = repo.db.QueryRow(
		q, schema.CVId, schema.ProjectId, schema.CompanyId, schema.EndTime, schema.StartTime,
	).Scan(&new_id)
	if err != nil {
		return 0, err
	}

	return new_id, nil
}

func (repo *CVRepository) UpdateProject(schema schemas.CVProject) error {
	q := `UPDATE cv_projects SET
    cv_id = $1, project_id = $2, company_id = $3, end_time = $4, start_time = $5
    WHERE id = $6`
	_, err := repo.db.Exec(
		q,
		schema.CVId,
		schema.ProjectId,
		schema.CompanyId,
		schema.EndTime,
		schema.StartTime,
		schema.Id,
	)
	return err
}

func (repo *CVRepository) DeleteProject(schema schemas.CVProject) error {
	q := `DELETE FROM cv_projects WHERE id = $1`
	_, err := repo.db.Exec(q, schema.Id)
	return err
}

func (repo *CVRepository) GetProjectsResponsibilities(id int) (schemes []schemas.CVProjectResponsibility, err error) {
	q := `SELECT id, cv_project_id, responsibility_id, priority
	FROM project_responsibilities
	WHERE cv_project_id = $1`

	rows, err := repo.db.Query(q, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var schema schemas.CVProjectResponsibility
		err = rows.Scan(
			&schema.Id,
			&schema.CVProjectId,
			&schema.ResponsibilityId,
			&schema.Priority,
		)
		if err != nil {
			return nil, err
		}
		schemes = append(schemes, schema)
	}

	if err := rows.Err(); err != nil {

		return schemes, err
	}

	return schemes, nil
}

func (repo *CVRepository) GetAllProjectResponsibilities() (schemes []schemas.CVProjectResponsibility, err error) {
	q := `SELECT id, cv_project_id, responsibility_id, priority
	FROM project_responsibilities`

	rows, err := repo.db.Query(q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var schema schemas.CVProjectResponsibility
		err = rows.Scan(
			&schema.Id,
			&schema.CVProjectId,
			&schema.ResponsibilityId,
			&schema.Priority,
		)
		if err != nil {
			return nil, err
		}
		schemes = append(schemes, schema)
	}

	if err := rows.Err(); err != nil {

		return schemes, err
	}

	return schemes, nil
}

func (repo *CVRepository) CreateProjectResponsibility(schema schemas.CVProjectResponsibility) (new_id int, err error) {
	q := `INSERT INTO project_responsibilities(cv_project_id, responsibility_id, priority) VALUES($1, $2, $3) RETURNING id`

	new_id = 0
	err = repo.db.QueryRow(
		q, schema.CVProjectId, schema.ResponsibilityId, schema.Priority,
	).Scan(&new_id)
	if err != nil {
		return 0, err
	}

	return new_id, nil
}

func (repo *CVRepository) UpdateProjectResponsibility(schema schemas.CVProjectResponsibility) error {
	q := `UPDATE project_responsibilities SET
    cv_project_id = $1, responsibility_id = $2, priority = $3
    WHERE id = $4`
	_, err := repo.db.Exec(
		q,
		schema.CVProjectId,
		schema.ResponsibilityId,
		schema.Priority,
		schema.Id,
	)
	return err
}

func (repo *CVRepository) DeleteProjectResponsibility(schema schemas.CVProjectResponsibility) error {
	q := `DELETE FROM project_responsibilities WHERE id = $1`
	_, err := repo.db.Exec(q, schema.Id)
	return err
}

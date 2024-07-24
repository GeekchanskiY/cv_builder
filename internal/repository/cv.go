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
	q := `INSERT INTO cvs(name, vacancy_id, employee_id, is_real) VALUES($1, $2, $3, $4) RETURNING id`
	newId := 0
	err := repo.db.QueryRow(q, schema.Name, schema.VacancyId, schema.EmployeeId, schema.IsReal).Scan(&newId)
	if err != nil {
		log.Println("Error creating cv in cv repository: ", err)
		return 0, err
	}

	return newId, nil
}

func (repo *CVRepository) CreateIfNotExists(schema schemas.CV) (created bool, err error) {
	q := `INSERT INTO cvs(name, vacancy_id, employee_id, is_real) 
	SELECT $1, $2, $3, $4
	WHERE 
	    NOT EXISTS (SELECT 1 FROM cvs WHERE vacancy_id = $1 AND employee_id = $2)`

	r, err := repo.db.Exec(q, schema.VacancyId, schema.EmployeeId, schema.IsReal)

	if err != nil {
		log.Println("Error creating cv: ", err)

		return false, err
	}

	if i, _ := r.RowsAffected(); i != 0 {
		return true, nil
	}

	return false, nil
}

func (repo *CVRepository) Update(schema schemas.CV) error {
	q := `UPDATE cvs SET name=$1, vacancy_id = $2, employee_id = $3, is_real = $4 WHERE id = $5`
	_, err := repo.db.Exec(q, schema.Name, schema.VacancyId, schema.EmployeeId, schema.IsReal, schema.Id)
	return err
}

func (repo *CVRepository) Delete(schema schemas.CV) error {
	_, err := repo.db.Exec("DELETE FROM cvs WHERE id = $1", schema.Id)
	return err
}

func (repo *CVRepository) GetAll() (schemes []schemas.CV, err error) {
	q := `SELECT id, name, vacancy_id, employee_id, is_real FROM cvs`
	rows, err := repo.db.Query(q)
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
		var schema schemas.CV
		err = rows.Scan(&schema.Id, &schema.Name, &schema.VacancyId, &schema.EmployeeId, &schema.IsReal)
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
	q := `SELECT id, name, vacancy_id, employee_id, is_real FROM cvs WHERE id = $1`
	row := repo.db.QueryRow(q, id)
	err = row.Scan(&schema.Id, &schema.Name, &schema.VacancyId, &schema.EmployeeId, &schema.IsReal)
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

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println()
		}
	}(rows)

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

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing rows: ", err)
		}
	}(rows)

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

func (repo *CVRepository) CreateProject(schema schemas.CVProject) (newId int, err error) {
	q := `INSERT INTO cv_projects(cv_id, project_id, company_id, end_time, start_time) VALUES($1, $2, $3, $4, $5) RETURNING id`

	err = repo.db.QueryRow(
		q, schema.CVId, schema.ProjectId, schema.CompanyId, schema.EndTime, schema.StartTime,
	).Scan(&newId)
	if err != nil {
		return 0, err
	}

	return newId, nil
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

func (repo *CVRepository) GetCVServices(id int) (schemes []schemas.CVService, err error) {
	q := `SELECT id, cv_project_id, name, order_num
	FROM cv_services
	WHERE cv_project_id = $1`

	rows, err := repo.db.Query(q, id)

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
		var schema schemas.CVService
		err = rows.Scan(
			&schema.Id,
			&schema.CVProjectId,
			&schema.Name,
			&schema.OrderNum,
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

func (repo *CVRepository) GetAllCVServices() (schemes []schemas.CVService, err error) {
	q := `SELECT id, cv_project_id, name, order_num
	FROM cv_services`

	rows, err := repo.db.Query(q)

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
		var schema schemas.CVService
		err = rows.Scan(
			&schema.Id,
			&schema.CVProjectId,
			&schema.Name,
			&schema.OrderNum,
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

func (repo *CVRepository) CreateCVService(schema schemas.CVService) (newId int, err error) {
	q := `INSERT INTO cv_services(cv_project_id, name, order_num) VALUES($1, $2, $3) RETURNING id`

	err = repo.db.QueryRow(
		q, schema.CVProjectId, schema.Name, schema.OrderNum,
	).Scan(&newId)
	if err != nil {
		return 0, err
	}

	return newId, nil
}

func (repo *CVRepository) UpdateCVService(schema schemas.CVService) error {
	q := `UPDATE cv_services SET
    cv_project_id = $1, name = $2, order_num = $3
    WHERE id = $4`
	_, err := repo.db.Exec(
		q,
		schema.CVProjectId,
		schema.Name,
		schema.OrderNum,
		schema.Id,
	)
	return err
}

func (repo *CVRepository) DeleteCVService(schema schemas.CVService) error {
	q := `DELETE FROM cv_services WHERE id = $1`
	_, err := repo.db.Exec(q, schema.Id)
	return err
}

func (repo *CVRepository) GetCVServiceResponsibilities(id int) (schemes []schemas.CVServiceResponsibility, err error) {
	q := `SELECT id, cv_service_id, responsibility_id, order_num
	FROM cv_service_responsibilities
	WHERE cv_service_id = $1`

	rows, err := repo.db.Query(q, id)

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
		var schema schemas.CVServiceResponsibility
		err = rows.Scan(
			&schema.Id,
			&schema.CVServiceId,
			&schema.ResponsibilityId,
			&schema.OrderNum,
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

func (repo *CVRepository) GetAllCVServiceResponsibilities() (schemes []schemas.CVServiceResponsibility, err error) {
	q := `SELECT id, cv_service_id, responsibility_id, order_num
	FROM cv_service_responsibilities`

	rows, err := repo.db.Query(q)

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
		var schema schemas.CVServiceResponsibility
		err = rows.Scan(
			&schema.Id,
			&schema.CVServiceId,
			&schema.ResponsibilityId,
			&schema.OrderNum,
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

func (repo *CVRepository) CreateCVServiceResponsibility(schema schemas.CVServiceResponsibility) (newId int, err error) {
	q := `INSERT INTO cv_service_responsibilities(cv_service_id, responsibility_id, order_num) VALUES($1, $2, $3) RETURNING id`

	err = repo.db.QueryRow(
		q, schema.CVServiceId, schema.ResponsibilityId, schema.OrderNum,
	).Scan(&newId)
	if err != nil {
		return 0, err
	}

	return newId, nil
}

func (repo *CVRepository) UpdateCVServiceResponsibility(schema schemas.CVServiceResponsibility) error {
	q := `UPDATE cv_service_responsibilities SET
    cv_service_id = $1, responsibility_id = $2, order_num = $3
    WHERE id = $4`
	_, err := repo.db.Exec(
		q,
		schema.CVServiceId,
		schema.ResponsibilityId,
		schema.OrderNum,
		schema.Id,
	)
	return err
}

func (repo *CVRepository) DeleteCVServiceResponsibility(schema schemas.CVServiceResponsibility) error {
	q := `DELETE FROM cv_service_responsibilities WHERE id = $1`
	_, err := repo.db.Exec(q, schema.Id)
	return err
}

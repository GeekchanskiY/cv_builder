package repository

import (
	"database/sql"
	"log"

	"github.com/GeekchanskiY/cv_builder/internal/schemas"
)

type EmployeeRepository struct {
	db *sql.DB
}

func CreateEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (repo *EmployeeRepository) Create(employee schemas.Employee) (newId int, err error) {
	q := `INSERT INTO employees(name, about_me, image_url, real_experience) VALUES($1, $2, $3, $4) RETURNING id`
	err = repo.db.QueryRow(
		q, employee.Name, employee.AboutMe, employee.ImageUrl, employee.RealExperience,
	).Scan(&newId)
	if err != nil {
		log.Println("Error creating employee in employee repository: ", err)
		return 0, err
	}

	return newId, nil
}

func (repo *EmployeeRepository) Count() (res int, err error) {
	q := `SELECT COUNT(*) FROM employees`

	err = repo.db.QueryRow(q).Scan(&res)

	if err != nil {
		log.Println("Error getting amount of employees: ", err)

		return 0, err
	}

	return res, nil
}

func (repo *EmployeeRepository) CreateIfNotExists(schema schemas.Employee) (created bool, err error) {
	q := `INSERT INTO employees(name, about_me, image_url, real_experience) 
	SELECT CAST($1 AS VARCHAR), $2, $3, $4
	WHERE 
	    NOT EXISTS (SELECT 1 FROM employees WHERE name = $1)`

	r, err := repo.db.Exec(q, schema.Name, schema.AboutMe, schema.ImageUrl, schema.RealExperience)

	if err != nil {
		log.Println("Error creating employee: ", err)

		return false, err
	}

	if i, _ := r.RowsAffected(); i != 0 {
		return true, nil
	}

	return false, nil
}

func (repo *EmployeeRepository) Update(employee schemas.Employee) error {
	q := `UPDATE employees SET name = $1, about_me = $2, image_url = $3, real_experience = $4 WHERE id = $5`
	_, err := repo.db.Exec(q, employee.Name, employee.AboutMe, employee.ImageUrl, employee.RealExperience, employee.Id)
	return err
}

func (repo *EmployeeRepository) Delete(employee schemas.Employee) error {
	q := `DELETE FROM employees WHERE id = $1`
	_, err := repo.db.Exec(q, employee.Id)
	return err
}

func (repo *EmployeeRepository) GetAll() ([]schemas.Employee, error) {
	var employees []schemas.Employee
	q := `SELECT id, name, about_me, image_url, real_experience FROM employees`
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
		var employee schemas.Employee
		err = rows.Scan(
			&employee.Id,
			&employee.Name,
			&employee.AboutMe,
			&employee.ImageUrl,
			&employee.RealExperience,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return employees, err
	}

	return employees, nil
}

func (repo *EmployeeRepository) Get(id int) (schemas.Employee, error) {
	var employee schemas.Employee
	q := `SELECT id, name, about_me, image_url, real_experience FROM employees WHERE id = $1`
	row := repo.db.QueryRow(q, id)
	err := row.Scan(&employee.Id, &employee.Name, &employee.AboutMe, &employee.ImageUrl, &employee.RealExperience)
	return employee, err
}

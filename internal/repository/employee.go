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

func (repo *EmployeeRepository) Create(employee schemas.Employee) (int, error) {
	new_id := 0
	err := repo.db.QueryRow("INSERT INTO employees(name) VALUES($1) RETURNING id", employee.Name).Scan(&new_id)
	if err != nil {
		log.Println("Error creating employee in employee repository: ", err)
		// log.Println(err)
		// pgerr, ok := utils.PQErrorHandler(err)
		// if ok {
		// 	// log.Println(pgerr.Code)
		// 	// log.Println(pgerr.Message)

		// 	if pgerr.Code == pg_err_unique_violation {
		// 		return nil
		// 	}
		// }
		return 0, err
	}

	return int(new_id), nil
}

func (repo *EmployeeRepository) Update(employee schemas.Employee) error {
	_, err := repo.db.Exec("UPDATE employees SET name = $1 WHERE id = $2", employee.Name, employee.Id)
	return err
}

func (repo *EmployeeRepository) Delete(employee schemas.Employee) error {
	_, err := repo.db.Exec("DELETE FROM employees WHERE id = $1", employee.Id)
	return err
}

func (repo *EmployeeRepository) GetAll() ([]schemas.Employee, error) {
	var employees []schemas.Employee

	rows, err := repo.db.Query("SELECT id, name FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee schemas.Employee
		if err := rows.Scan(&employee.Id, &employee.Name); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		// Here's not nil, err. I'm not sure why, but:
		// https://go.dev/doc/database/querying
		// documentation has the same issue

		return employees, err
	}

	return employees, nil
}

func (repo *EmployeeRepository) Get(id int) (schemas.Employee, error) {
	var employee schemas.Employee
	row := repo.db.QueryRow("SELECT id, name FROM employees WHERE id = $1", id)
	err := row.Scan(&employee.Id, &employee.Name)
	return employee, err
}

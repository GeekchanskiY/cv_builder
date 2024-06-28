package repository

import (
	"database/sql"
	"log"

	"github.com/GeekchanskiY/cv_builder/internal/schemas"
	"github.com/GeekchanskiY/cv_builder/internal/utils"
)

type EmployeeRepository struct {
	db *sql.DB
}

func CreateEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (repo *EmployeeRepository) CreateEmployee(employee schemas.Employee) error {
	_, err := repo.db.Exec("INSERT INTO employees(name) VALUES($1)", employee.Name)
	if err != nil {
		log.Println("Error creating employee")
		log.Println(err)
		pgerr, ok := utils.PQErrorHandler(err)
		if ok {
			// log.Println(pgerr.Code)
			// log.Println(pgerr.Message)

			if pgerr.Code == pg_err_unique_violation {
				return nil
			}
		}
		return err
	}
	return err
}

func (repo *EmployeeRepository) GetEmployees() ([]schemas.Employee, error) {
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

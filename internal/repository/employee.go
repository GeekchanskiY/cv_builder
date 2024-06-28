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
	if err := repo.db.QueryRow("SELECT * FROM employees").Scan(&employees); err != nil {
		if err == sql.ErrNoRows {
			return employees, nil
		}
		log.Println("Error getting employees")
		log.Println(err)
		return nil, err
	}
	return employees, nil
}

package repository

import (
	"database/sql"
	"fmt"

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

func (repo *EmployeeRepository) CreateEmployee(employee schemas.Employee) error {
	res, err := repo.db.Exec("INSERT INTO employees(name) VALUES($1)", employee.Name)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return err
}

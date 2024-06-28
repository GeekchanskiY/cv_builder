package internal

import (
	"database/sql"
	"fmt"

	"github.com/GeekchanskiY/cv_builder/internal/repository"
	"github.com/GeekchanskiY/cv_builder/internal/schemas"
)

func Samples(db *sql.DB) {
	a := repository.CreateEmployeeRepository(db)
	err := a.CreateEmployee(schemas.Employee{
		Name: "test1",
	})
	fmt.Println(err)
}

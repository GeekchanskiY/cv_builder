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

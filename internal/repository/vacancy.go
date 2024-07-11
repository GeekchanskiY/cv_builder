package repository

import (
	"database/sql"
	"log"

	"github.com/GeekchanskiY/cv_builder/internal/schemas"
)

type VacanciesRepository struct {
	db *sql.DB
}

func CreateVacanciesRepository(db *sql.DB) *VacanciesRepository {
	return &VacanciesRepository{
		db: db,
	}
}

func (repo *VacanciesRepository) Create(schema schemas.Vacancy) (int, error) {
	q := `INSERT INTO vacancies(name, company_id, link, description, published_at, experience) VALUES($1) RETURNING id`
	new_id := 0
	err := repo.db.QueryRow(q, schema.Name, schema.CompanyId, schema.Link, schema.Description, schema.PublishedAt, schema.Experience).Scan(&new_id)
	if err != nil {
		log.Println("Error creating schema in repository: ", err)

		return 0, err
	}

	return int(new_id), nil
}

func (repo *VacanciesRepository) Update(schema schemas.Vacancy) error {
	q := `UPDATE vacancies SET name = $1, company_id = $2, link = $3, description = $4, published_at = $5, experience = $6 WHERE id = $7`
	_, err := repo.db.Exec(q, schema.Name, schema.CompanyId, schema.Link, schema.Description, schema.PublishedAt, schema.Experience, schema.Id)
	return err
}

func (repo *VacanciesRepository) Delete(schema schemas.Vacancy) error {
	q := `DELETE FROM vacancies WHERE id = $1`
	_, err := repo.db.Exec(q, schema.Id)
	return err
}

func (repo *VacanciesRepository) GetAll() (schemasArr []schemas.Vacancy, err error) {
	q := `SELECT id, name, company_id, link, description, published_at, experience FROM vacancies`
	rows, err := repo.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schema schemas.Vacancy
		if err := rows.Scan(&schema.Id, &schema.Name, &schema.CompanyId, &schema.Link, &schema.Description, &schema.PublishedAt, &schema.Experience); err != nil {
			return nil, err
		}
		schemasArr = append(schemasArr, schema)
	}

	if err := rows.Err(); err != nil {
		// Here's not nil, err. I'm not sure why, but:
		// https://go.dev/doc/database/querying
		// documentation has the same issue

		return schemasArr, err
	}

	return schemasArr, nil
}

func (repo *VacanciesRepository) Get(id int) (schema schemas.Vacancy, err error) {
	row := repo.db.QueryRow("SELECT id, name, company_id, link, description, published_at, experience FROM vacancies WHERE id = $1", id)
	err = row.Scan(&schema.Id, &schema.Name, &schema.CompanyId, &schema.Link, &schema.Description, &schema.PublishedAt, &schema.Experience)
	return schema, err
}

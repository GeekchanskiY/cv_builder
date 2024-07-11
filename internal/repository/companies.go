package repository

import (
	"database/sql"
	"log"

	"github.com/GeekchanskiY/cv_builder/internal/schemas"
)

type CompanyRepository struct {
	db *sql.DB
}

func CreateCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (repo *CompanyRepository) Create(company schemas.Company) (int, error) {
	new_id := 0
	err := repo.db.QueryRow("INSERT INTO companies(name) VALUES($1) RETURNING id", company.Name).Scan(&new_id)
	if err != nil {
		log.Println("Error creating company in company repository: ", err)

		return 0, err
	}

	return int(new_id), nil
}

func (repo *CompanyRepository) Update(company schemas.Company) error {
	_, err := repo.db.Exec("UPDATE companies SET name = $1 WHERE id = $2", company.Name, company.Id)
	return err
}

func (repo *CompanyRepository) Delete(company schemas.Company) error {
	_, err := repo.db.Exec("DELETE FROM companies WHERE id = $1", company.Id)
	return err
}

func (repo *CompanyRepository) GetAll() (companies []schemas.Company, err error) {
	rows, err := repo.db.Query("SELECT id, name FROM companies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var company schemas.Company
		if err := rows.Scan(&company.Id, &company.Name); err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	if err := rows.Err(); err != nil {

		return companies, err
	}

	return companies, nil
}

func (repo *CompanyRepository) Get(id int) (company schemas.Company, err error) {
	row := repo.db.QueryRow("SELECT id, name FROM companies WHERE id = $1", id)
	err = row.Scan(&company.Id, &company.Name)
	return company, err
}

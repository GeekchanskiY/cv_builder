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
	q := `INSERT INTO companies(name, description, homepage, is_trusted) VALUES($1, $2, $3, $4) RETURNING id`

	err := repo.db.QueryRow(q, company.Name, company.Description, company.Homepage, company.IsTrusted).Scan(&new_id)

	if err != nil {
		log.Println("Error creating company in company repository: ", err)

		return 0, err
	}

	return new_id, nil
}

func (repo *CompanyRepository) CreateIfNotExists(company schemas.Company) (created bool, err error) {
	// Cast is required
	// https://stackoverflow.com/questions/31733790/postgresql-parameter-issue-1
	q := `INSERT INTO companies(name, description, homepage, is_trusted) 
	SELECT CAST($1 AS VARCHAR) AS name, $2 AS description, $3 AS homepage, $4 AS is_trusted
	WHERE 
	    NOT EXISTS (SELECT 1 FROM companies WHERE name = $1)
	RETURNING id`

	r, err := repo.db.Exec(q, company.Name, company.Description, company.Homepage, company.IsTrusted)

	if err != nil {
		log.Println("Error creating company in company repository: ", err)

		return false, err
	}

	if i, _ := r.RowsAffected(); i != 0 {
		return true, nil
	}

	return false, nil
}

func (repo *CompanyRepository) Update(company schemas.Company) error {
	q := `UPDATE companies SET name = $1, description = $2, homepage = $3, is_trusted = $4 WHERE id = $5`
	_, err := repo.db.Exec(q, company.Name, company.Description, company.Homepage, company.IsTrusted, company.Id)
	return err
}

func (repo *CompanyRepository) Delete(company schemas.Company) error {
	_, err := repo.db.Exec(`DELETE FROM companies WHERE id = $1`, company.Id)
	return err
}

func (repo *CompanyRepository) GetAll() (companies []schemas.Company, err error) {
	q := `SELECT id, name, description, homepage, is_trusted FROM companies`
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
		var company schemas.Company
		err = rows.Scan(&company.Id, &company.Name, &company.Description, &company.Homepage, &company.IsTrusted)
		if err != nil {
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
	row := repo.db.QueryRow("SELECT id, name, description, homepage, is_trusted FROM companies WHERE id = $1", id)
	err = row.Scan(&company.Id, &company.Name, &company.Description, &company.Homepage, &company.IsTrusted)
	return company, err
}

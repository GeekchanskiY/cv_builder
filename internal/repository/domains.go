package repository

import (
	"database/sql"
	"log"

	"github.com/GeekchanskiY/cv_builder/internal/schemas"
)

type DomainRepository struct {
	db *sql.DB
}

func CreateDomainRepository(db *sql.DB) *DomainRepository {
	return &DomainRepository{
		db: db,
	}
}

func (repo *DomainRepository) Create(domain schemas.Domain) (int, error) {
	newId := 0
	q := `INSERT INTO domains(name, description) VALUES($1, $2) RETURNING id`
	err := repo.db.QueryRow(q, domain.Name, domain.Description).Scan(&newId)
	if err != nil {
		log.Println("Error creating domain in domain repository: ", err)
		return 0, err
	}

	return newId, nil
}

func (repo *DomainRepository) Update(domain schemas.Domain) error {
	q := `UPDATE domains SET name = $1, description = $2 WHERE id = $3`
	_, err := repo.db.Exec(q, domain.Name, domain.Description, domain.Id)
	return err
}

func (repo *DomainRepository) Delete(domain schemas.Domain) error {
	_, err := repo.db.Exec("DELETE FROM domains WHERE id = $1", domain.Id)
	return err
}

func (repo *DomainRepository) GetAll() (domains []schemas.Domain, err error) {
	q := `SELECT id, name, description FROM domains`
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
		var domain schemas.Domain
		err = rows.Scan(&domain.Id, &domain.Name, &domain.Description)
		if err != nil {
			return nil, err
		}
		domains = append(domains, domain)
	}

	if err := rows.Err(); err != nil {
		return domains, err
	}

	return domains, nil
}

func (repo *DomainRepository) Get(id int) (domain schemas.Domain, err error) {
	q := `SELECT id, name, description FROM domains WHERE id = $1`
	row := repo.db.QueryRow(q, id)
	err = row.Scan(&domain.Id, &domain.Name, &domain.Description)
	return domain, err
}

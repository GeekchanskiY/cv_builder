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
	new_id := 0
	err := repo.db.QueryRow("INSERT INTO domains(name) VALUES($1) RETURNING id", domain.Name).Scan(&new_id)
	if err != nil {
		log.Println("Error creating domain in domain repository: ", err)
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

func (repo *DomainRepository) Update(domain schemas.Domain) error {
	_, err := repo.db.Exec("UPDATE domains SET name = $1 WHERE id = $2", domain.Name, domain.Id)
	return err
}

func (repo *DomainRepository) Delete(domain schemas.Domain) error {
	_, err := repo.db.Exec("DELETE FROM domains WHERE id = $1", domain.Id)
	return err
}

func (repo *DomainRepository) GetAll() (domains []schemas.Domain, err error) {
	rows, err := repo.db.Query("SELECT id, name FROM domains")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var domain schemas.Domain
		if err := rows.Scan(&domain.Id, &domain.Name); err != nil {
			return nil, err
		}
		domains = append(domains, domain)
	}

	if err := rows.Err(); err != nil {
		// Here's not nil, err. I'm not sure why, but:
		// https://go.dev/doc/database/querying
		// documentation has the same issue

		return domains, err
	}

	return domains, nil
}

func (repo *DomainRepository) Get(id int) (domain schemas.Domain, err error) {
	row := repo.db.QueryRow("SELECT id, name FROM domains WHERE id = $1", id)
	err = row.Scan(&domain.Id, &domain.Name)
	return domain, err
}

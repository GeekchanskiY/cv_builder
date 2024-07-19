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
	q := `INSERT INTO vacancies(name, company_id, link, description, published_at, experience) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
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

func (repo *VacanciesRepository) GetSkills(id int) (vacancySkills []schemas.VacancySkill, err error) {
	q := `SELECT id, skill_id, vacancy_id, priority FROM vacancy_skills
	WHERE vacancy_id = $1`

	rows, err := repo.db.Query(q, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var vacancySkill schemas.VacancySkill
		err = rows.Scan(&vacancySkill.Id, &vacancySkill.SkillId, &vacancySkill.VacancyId, &vacancySkill.Priority)
		if err != nil {
			return nil, err
		}
		vacancySkills = append(vacancySkills, vacancySkill)
	}

	if err := rows.Err(); err != nil {
		// Here's not nil, err. I'm not sure why, but:
		// https://go.dev/doc/database/querying
		// documentation has the same issue

		return vacancySkills, err
	}

	return vacancySkills, nil
}

func (repo *VacanciesRepository) GetAllSkills() (vacancySkills []schemas.VacancySkill, err error) {
	q := `SELECT id, skill_id, vacancy_id, priority FROM vacancy_skills`

	rows, err := repo.db.Query(q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var vacancySkill schemas.VacancySkill
		err = rows.Scan(&vacancySkill.Id, &vacancySkill.SkillId, &vacancySkill.VacancyId, &vacancySkill.Priority)
		if err != nil {
			return nil, err
		}
		vacancySkills = append(vacancySkills, vacancySkill)
	}

	if err := rows.Err(); err != nil {
		// Here's not nil, err. I'm not sure why, but:
		// https://go.dev/doc/database/querying
		// documentation has the same issue

		return vacancySkills, err
	}

	return vacancySkills, nil
}

func (repo *VacanciesRepository) AddSkill(schema schemas.VacancySkill) (new_id int, err error) {
	q := `INSERT INTO vacancy_skills(vacancy_id, skill_id, priority) VALUES($1, $2, $3) returning id`
	err = repo.db.QueryRow(q, schema.VacancyId, schema.SkillId, schema.Priority).Scan(&new_id)
	if err != nil {
		log.Println("Error creating schema in repository: ", err)

		return 0, err
	}

	return new_id, nil

}

func (repo *VacanciesRepository) DeleteSkill(schema schemas.VacancySkill) (err error) {
	q := `DELETE FROM vacancy_skills WHERE id = $1`
	_, err = repo.db.Exec(q, schema.Id)
	return err
}

func (repo *VacanciesRepository) UpdateSkill(schema schemas.VacancySkill) error {
	q := `UPDATE vacancy_skills SET skill_id = $1, vacancy_id = $2, priority = $3 WHERE id = $4`
	_, err := repo.db.Exec(
		q,
		schema.SkillId,
		schema.VacancyId,
		schema.Priority,
		schema.Id,
	)
	return err
}

func (repo *VacanciesRepository) GetDomains(id int) (schemes []schemas.VacancyDomain, err error) {
	q := `SELECT id, vacancy_id, domain_id, priority FROM vacancy_domains
	WHERE vacancy_id = $1`

	rows, err := repo.db.Query(q, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var schema schemas.VacancyDomain
		err = rows.Scan(&schema.Id, &schema.VacancyId, &schema.DomainId, &schema.Priority)
		if err != nil {
			return nil, err
		}
		schemes = append(schemes, schema)
	}

	if err := rows.Err(); err != nil {
		// Here's not nil, err. I'm not sure why, but:
		// https://go.dev/doc/database/querying
		// documentation has the same issue

		return schemes, err
	}

	return schemes, nil
}

func (repo *VacanciesRepository) GetAllDomains() (schemes []schemas.VacancyDomain, err error) {
	q := `SELECT id, vacancy_id, domain_id, priority FROM vacancy_domains`

	rows, err := repo.db.Query(q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var schema schemas.VacancyDomain
		err = rows.Scan(&schema.Id, &schema.VacancyId, &schema.DomainId, &schema.Priority)
		if err != nil {
			return nil, err
		}
		schemes = append(schemes, schema)
	}

	if err := rows.Err(); err != nil {
		// Here's not nil, err. I'm not sure why, but:
		// https://go.dev/doc/database/querying
		// documentation has the same issue

		return schemes, err
	}

	return schemes, nil
}

func (repo *VacanciesRepository) AddDomain(schema schemas.VacancyDomain) (new_id int, err error) {
	q := `INSERT INTO vacancy_domains(vacancy_id, domain_id, priority) VALUES($1, $2, $3) returning id`
	err = repo.db.QueryRow(q, schema.VacancyId, schema.DomainId, schema.Priority).Scan(&new_id)
	if err != nil {
		log.Println("Error creating schema in repository: ", err)

		return 0, err
	}

	return new_id, nil

}

func (repo *VacanciesRepository) DeleteDomain(schema schemas.VacancyDomain) (err error) {
	q := `DELETE FROM vacancy_domains WHERE id = $1`
	_, err = repo.db.Exec(q, schema.Id)
	return err
}

func (repo *VacanciesRepository) UpdateDomain(schema schemas.VacancyDomain) error {
	q := `UPDATE vacancy_domains SET vacancy_id = $1, domain_id = $2, priority = $3 WHERE id = $4`
	_, err := repo.db.Exec(
		q,
		schema.VacancyId,
		schema.DomainId,
		schema.Priority,
		schema.Id,
	)
	return err
}

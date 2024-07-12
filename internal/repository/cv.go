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
	q := `INSERT INTO cvs(vacancy_id, employee_id) VALUES($1, $2) RETURNING id`
	new_id := 0
	err := repo.db.QueryRow(q, schema.VacancyId, schema.EmployeeId).Scan(&new_id)
	if err != nil {
		log.Println("Error creating cv in cv repository: ", err)
		return 0, err
	}

	return int(new_id), nil
}

func (repo *CVRepository) Update(schema schemas.CV) error {
	q := `UPDATE cvs SET vacancy_id = $1, employee_id = $2 WHERE id = $3`
	_, err := repo.db.Exec(q, schema.VacancyId, schema.EmployeeId, schema.Id)
	return err
}

func (repo *CVRepository) Delete(schema schemas.CV) error {
	_, err := repo.db.Exec("DELETE FROM cvs WHERE id = $1", schema.Id)
	return err
}

func (repo *CVRepository) GetAll() (schemes []schemas.CV, err error) {
	rows, err := repo.db.Query("SELECT id, vacancy_id, employee_id FROM cvs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schema schemas.CV
		if err := rows.Scan(&schema.Id, &schema.VacancyId, &schema.EmployeeId); err != nil {
			return nil, err
		}
		schemes = append(schemes, schema)
	}

	if err := rows.Err(); err != nil {

		return schemes, err
	}

	return schemes, nil
}

func (repo *CVRepository) Get(id int) (schema schemas.CV, err error) {
	row := repo.db.QueryRow("SELECT id, vacancy_id, employee_id FROM cvs WHERE id = $1", id)
	err = row.Scan(&schema.Id, &schema.VacancyId, &schema.EmployeeId)
	return schema, err
}

func (repo *CVRepository) AddSkill(schema schemas.CvSkill) (err error) {
	q := `INSERT INTO cv_skills(cv_id, skill_id, years) VALUES($1, $2)`
	_, err = repo.db.Exec(q, schema.CvId, schema.SkillId, schema.Years)
	return err
}

func (repo *CVRepository) GetSkills(id int) (schemes []schemas.CVSkillExtension, err error) {
	q := `SELECT skill.id, skill.name, skill.description, skill.parent_id, cv_skills.years
	FROM skills skill 
	JOIN cv_skills ON cv_skills.skill_id = skill.id
	WHERE cv_skills.cv_id = $1`

	rows, err := repo.db.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var schema schemas.CVSkillExtension
		if err := rows.Scan(&schema.SkillId, &schema.Name, &schema.Description, &schema.ParentId, &schema.Years); err != nil {
			return nil, err
		}
		schemes = append(schemes, schema)
	}

	if err := rows.Err(); err != nil {

		return schemes, err
	}

	return schemes, nil
}

func (repo *CVRepository) DeleteSkill(schema schemas.CvSkill) (err error) {
	q := `DELETE FROM cv_skills WHERE cv_id = $1 AND skill_id = $2`
	_, err = repo.db.Exec(q, schema.CvId, schema.SkillId)
	return err
}

package repository

import (
	"database/sql"
	"log"

	"github.com/GeekchanskiY/cv_builder/internal/schemas"
)

type ResponsibilityRepository struct {
	db *sql.DB
}

func CreateResponsibilityRepository(db *sql.DB) *ResponsibilityRepository {
	return &ResponsibilityRepository{
		db: db,
	}
}

func (repo *ResponsibilityRepository) Create(schema schemas.Responsibility) (new_id int, err error) {
	q := `INSERT INTO responsibilities(name, priority, skill_id, experience, comments) VALUES($1, $2, $3, $4, $5) RETURNING id`
	err = repo.db.QueryRow(
		q, schema.Name, schema.Priority, schema.SkillId, schema.Experience, schema.Comments,
	).Scan(&new_id)
	if err != nil {
		log.Println("Error creating responsibility: ", err)
		return 0, err
	}

	return new_id, nil
}

func (repo *ResponsibilityRepository) Update(schema schemas.Responsibility) error {
	q := `UPDATE responsibilities SET name = $1, priority = $2, skill_id = $3, experience = $4, comments = $5 WHERE id = $6`
	_, err := repo.db.Exec(q, schema.Name, schema.Priority, schema.SkillId, schema.Experience, schema.Comments, schema.Id)
	return err
}

func (repo *ResponsibilityRepository) Delete(schema schemas.Responsibility) error {
	q := `DELETE FROM responsibilities WHERE id = $1`
	_, err := repo.db.Exec(q, schema.Id)
	return err
}

func (repo *ResponsibilityRepository) GetAll() (schemes []schemas.Responsibility, err error) {
	q := `SELECT id, name, priority, skill_id, experience, comments FROM responsibilities`
	rows, err := repo.db.Query(q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var schema schemas.Responsibility
		err = rows.Scan(
			&schema.Id, &schema.Name, &schema.Priority, &schema.SkillId, &schema.Experience, &schema.Comments,
		)
		if err != nil {
			return nil, err
		}
		schemes = append(schemes, schema)
	}

	if err = rows.Err(); err != nil {
		return schemes, err
	}

	if err = rows.Close(); err != nil {
		return schemes, err
	}

	return schemes, err
}

func (repo *ResponsibilityRepository) Get(id int) (schema schemas.Responsibility, err error) {
	q := `SELECT id, name, priority, skill_id, experience, comments FROM responsibilities WHERE id = $1`
	row := repo.db.QueryRow(q, id)
	err = row.Scan(&schema.Id, &schema.Name, &schema.Priority, &schema.SkillId, &schema.Experience, &schema.Comments)
	return schema, err
}

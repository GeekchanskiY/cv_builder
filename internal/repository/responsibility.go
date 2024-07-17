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

func (repo *ResponsibilityRepository) GetConflicts(id int) (conflicts []schemas.ResponsibilityConflict, err error) {
	q := `SELECT id, responsibility_1_id, responsibility_2_id, comment, priority 
	FROM responsibility_conflicts
	WHERE responsibility_1_id = $1 OR responsibility_2_id = $1`

	rows, err := repo.db.Query(q, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var conflict schemas.ResponsibilityConflict
		if err := rows.Scan(&conflict.Id, &conflict.Responsibility1Id, &conflict.Responsibility2Id, &conflict.Comment, &conflict.Priority); err != nil {
			return nil, err
		}
		conflicts = append(conflicts, conflict)
	}

	if err := rows.Err(); err != nil {
		// Here's not nil, err. I'm not sure why, but:
		// https://go.dev/doc/database/querying
		// documentation has the same issue

		return conflicts, err
	}

	return conflicts, nil
}

func (repo *ResponsibilityRepository) CreateConflict(conflict schemas.ResponsibilityConflict) (new_id int, err error) {
	q := `INSERT INTO responsibility_conflicts(responsibility_1_id, responsibility_2_id, comment, priority) VALUES($1, $2, $3, $4) RETURNING id`

	new_id = 0
	err = repo.db.QueryRow(
		q, conflict.Responsibility1Id, conflict.Responsibility2Id, conflict.Comment, conflict.Priority,
	).Scan(&new_id)
	if err != nil {
		log.Println("Error creating skillConflict in skillConflict repository: ", err)
		return 0, err
	}

	return new_id, nil
}

func (repo *ResponsibilityRepository) UpdateConflict(conflict schemas.ResponsibilityConflict) error {
	q := `UPDATE responsibility_conflicts SET responsibility_1_id = $1, responsibility_2_id = $2, comment = $3, priority = $4 WHERE id = $5`
	_, err := repo.db.Exec(
		q,
		conflict.Responsibility1Id,
		conflict.Responsibility2Id,
		conflict.Comment,
		conflict.Priority,
		conflict.Id,
	)
	return err
}

func (repo *ResponsibilityRepository) DeleteConflict(conflict schemas.ResponsibilityConflict) error {
	q := `DELETE FROM responsibility_conflicts WHERE id = $1`
	_, err := repo.db.Exec(q, conflict.Id)
	return err
}

func (repo *ResponsibilityRepository) GetSynonyms(id int) (schemes []schemas.ResponsibilitySynonym, err error) {
	q := `SELECT id, responsibility_id, name
	FROM responsibility_synonyms
	WHERE responsibility_id = $1`

	rows, err := repo.db.Query(q, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var schema schemas.ResponsibilitySynonym
		if err := rows.Scan(&schema.Id, &schema.ResponsibilityId, &schema.Name); err != nil {
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

func (repo *ResponsibilityRepository) CreateSynonym(schema schemas.ResponsibilitySynonym) (new_id int, err error) {
	q := `INSERT INTO responsibility_synonyms(responsibility_id, name) VALUES($1, $2) RETURNING id`

	new_id = 0
	err = repo.db.QueryRow(
		q, schema.ResponsibilityId, schema.Name,
	).Scan(&new_id)
	if err != nil {
		log.Println("Error creating responsibility synonym: ", err)
		return 0, err
	}

	return new_id, nil
}

func (repo *ResponsibilityRepository) UpdateSynonym(schema schemas.ResponsibilitySynonym) error {
	q := `UPDATE responsibility_synonyms SET responsibility_id = $1, name = $2 WHERE id = $3`
	_, err := repo.db.Exec(
		q,
		schema.ResponsibilityId,
		schema.Name,
		schema.Id,
	)
	return err
}

func (repo *ResponsibilityRepository) DeleteSynonym(schema schemas.ResponsibilitySynonym) error {
	q := `DELETE FROM responsibility_synonyms WHERE id = $1`
	_, err := repo.db.Exec(q, schema.Id)
	return err
}
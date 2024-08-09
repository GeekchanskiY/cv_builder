package repository

import (
	"database/sql"
	"log"

	"github.com/solndev/cv_builder/internal/schemas"
)

type ResponsibilityRepository struct {
	db *sql.DB
}

func CreateResponsibilityRepository(db *sql.DB) *ResponsibilityRepository {
	return &ResponsibilityRepository{
		db: db,
	}
}

func (repo *ResponsibilityRepository) Create(schema schemas.Responsibility) (newId int, err error) {
	q := `INSERT INTO responsibilities(name, priority, skill_id, experience, comments) VALUES($1, $2, $3, $4, $5) RETURNING id`
	err = repo.db.QueryRow(
		q, schema.Name, schema.Priority, schema.SkillId, schema.Experience, schema.Comments,
	).Scan(&newId)
	if err != nil {
		log.Println("Error creating responsibility: ", err)
		return 0, err
	}

	return newId, nil
}

func (repo *ResponsibilityRepository) Count() (res int, err error) {
	q := `SELECT COUNT(*) FROM responsibilities`

	err = repo.db.QueryRow(q).Scan(&res)

	if err != nil {
		log.Println("Error getting amount of responsibilities: ", err)

		return 0, err
	}

	return res, nil
}

func (repo *ResponsibilityRepository) CreateIfNotExists(schema schemas.ResponsibilityReadable) (created bool, err error) {
	// Cast is required
	// https://stackoverflow.com/questions/31733790/postgresql-parameter-issue-1
	q := `INSERT INTO responsibilities(name, priority, skill_id, experience, comments) 
	SELECT CAST($1 AS VARCHAR), $2, s.id, $4, $5
	FROM skills s 
	WHERE s.name = $3::text
	AND NOT EXISTS(
	    select 1
	    FROM responsibilities r
	    WHERE r.name = CAST($1 AS VARCHAR)
	);`

	r, err := repo.db.Exec(q, schema.Name, schema.Priority, schema.SkillName, schema.Experience, schema.Comments)

	if err != nil {
		log.Println("Error creating responsibility: ", err)

		return false, err
	}

	if i, _ := r.RowsAffected(); i != 0 {
		return true, nil
	}

	return false, nil
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

func (repo *ResponsibilityRepository) GetAllReadable() (schemes []schemas.ResponsibilityReadable, err error) {
	q := `SELECT r.name, r.priority, s.name, r.experience, r.comments FROM responsibilities r
	JOIN skills s ON s.id = r.skill_id`
	rows, err := repo.db.Query(q)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var schema schemas.ResponsibilityReadable
		err = rows.Scan(
			&schema.Name, &schema.Priority, &schema.SkillName, &schema.Experience, &schema.Comments,
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

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing rows: ", err)
		}
	}(rows)

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

func (repo *ResponsibilityRepository) GetAllConflicts() (conflicts []schemas.ResponsibilityConflict, err error) {
	q := `SELECT id, responsibility_1_id, responsibility_2_id, comment, priority 
	FROM responsibility_conflicts`

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

func (repo *ResponsibilityRepository) GetAllConflictsReadable() (conflicts []schemas.ResponsibilityConflictReadable, err error) {
	q := `SELECT r1.name, r2.name, rs.comment, rs.priority 
	FROM responsibility_conflicts rs
	JOIN responsibilities r1 ON r1.id = rs.responsibility_1_id
	JOIN responsibilities r2 ON r2.id = rs.responsibility_2_id`

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
		var conflict schemas.ResponsibilityConflictReadable
		if err := rows.Scan(&conflict.Responsibility1Name, &conflict.Responsibility2Name, &conflict.Comment, &conflict.Priority); err != nil {
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

func (repo *ResponsibilityRepository) CreateConflict(conflict schemas.ResponsibilityConflict) (newId int, err error) {
	q := `INSERT INTO responsibility_conflicts(responsibility_1_id, responsibility_2_id, comment, priority) VALUES($1, $2, $3, $4) RETURNING id`

	err = repo.db.QueryRow(
		q, conflict.Responsibility1Id, conflict.Responsibility2Id, conflict.Comment, conflict.Priority,
	).Scan(&newId)
	if err != nil {
		log.Println("Error creating skillConflict in skillConflict repository: ", err)
		return 0, err
	}

	return newId, nil
}

func (repo *ResponsibilityRepository) CreateConflictIfNotExists(schema schemas.ResponsibilityConflictReadable) (created bool, err error) {
	q := `INSERT INTO responsibility_conflicts(responsibility_1_id, responsibility_2_id, comment, priority) 
	SELECT r1.id, r2.id, $3, $4
	FROM responsibilities r1
	JOIN responsibilities r2 ON r2.name = $2::text
	WHERE 
	    r1.name = $1::text 
	    AND NOT EXISTS (
		SELECT 1 
		FROM responsibility_conflicts rc
		WHERE rc.responsibility_1_id = r1.id
		AND rc.responsibility_2_id = r2.id 
		);`

	r, err := repo.db.Exec(q, schema.Responsibility1Name, schema.Responsibility2Name, schema.Comment, schema.Priority)

	if err != nil {
		log.Println("Error creating responsibility conflict: ", err)

		return false, err
	}

	if i, _ := r.RowsAffected(); i != 0 {
		return true, nil
	}

	return false, nil
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

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing rows: ", err)
		}
	}(rows)

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

func (repo *ResponsibilityRepository) GetAllSynonyms() (schemes []schemas.ResponsibilitySynonym, err error) {
	q := `SELECT id, responsibility_id, name
	FROM responsibility_synonyms`

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

func (repo *ResponsibilityRepository) GetAllSynonymsReadable() (schemes []schemas.ResponsibilitySynonymReadable, err error) {
	q := `SELECT r.name, rs.name
	FROM responsibility_synonyms rs JOIN responsibilities r on rs.responsibility_id = r.id`

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
		var schema schemas.ResponsibilitySynonymReadable
		if err := rows.Scan(&schema.ResponsibilityName, &schema.Name); err != nil {
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

func (repo *ResponsibilityRepository) CreateSynonym(schema schemas.ResponsibilitySynonym) (newId int, err error) {
	q := `INSERT INTO responsibility_synonyms(responsibility_id, name) VALUES($1, $2) RETURNING id`

	err = repo.db.QueryRow(
		q, schema.ResponsibilityId, schema.Name,
	).Scan(&newId)
	if err != nil {
		log.Println("Error creating responsibility synonym: ", err)
		return 0, err
	}

	return newId, nil
}

func (repo *ResponsibilityRepository) CreateSynonymIfNotExists(schema schemas.ResponsibilitySynonymReadable) (created bool, err error) {
	q := `INSERT INTO responsibility_synonyms(responsibility_id, name) 
	SELECT r.id, $2
	FROM responsibilities r
	WHERE 
	    r.name = $1::text 
	    AND NOT EXISTS (
		SELECT 1 
		FROM responsibility_synonyms rs
		WHERE rs.responsibility_id = r.id
		);`

	r, err := repo.db.Exec(q, schema.ResponsibilityName, schema.Name)

	if err != nil {
		log.Println("Error creating responsibility conflict: ", err)

		return false, err
	}

	if i, _ := r.RowsAffected(); i != 0 {
		return true, nil
	}

	return false, nil
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

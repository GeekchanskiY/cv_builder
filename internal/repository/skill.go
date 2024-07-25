package repository

import (
	"database/sql"
	"log"

	"github.com/GeekchanskiY/cv_builder/internal/schemas"
)

type SkillRepository struct {
	db *sql.DB
}

func CreateSkillRepository(db *sql.DB) *SkillRepository {
	return &SkillRepository{
		db: db,
	}
}

func (repo *SkillRepository) Create(skill schemas.Skill) (int, error) {
	newId := 0
	err := repo.db.QueryRow(
		"INSERT INTO skills(name, description, parent_id) VALUES($1, $2, $3) RETURNING id",
		skill.Name, skill.Description, skill.ParentId,
	).Scan(&newId)

	if err != nil {
		log.Println("Error creating skill in skill repository: ", err)
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

	return newId, nil
}

func (repo *SkillRepository) CreateIfNotExists(schema schemas.SkillReadable) (created bool, err error) {
	// Cast is required
	// https://stackoverflow.com/questions/31733790/postgresql-parameter-issue-1
	q := `INSERT INTO skills(name, description, parent_id) 
	SELECT CAST($1 AS VARCHAR) AS name, $2 AS description, (SELECT id from skills where name = CAST($3 AS VARCHAR))
	WHERE NOT EXISTS(
		select 1
		FROM skills
		WHERE name = CAST($1 AS VARCHAR)
	);`

	r, err := repo.db.Exec(q, schema.Name, schema.Description, schema.ParentName)

	if err != nil {
		log.Println("Error creating skill in skill repository: ", err)

		return false, err
	}

	if i, _ := r.RowsAffected(); i != 0 {
		return true, nil
	}

	return false, nil
}

func (repo *SkillRepository) Update(skill schemas.Skill) error {
	_, err := repo.db.Exec("UPDATE skills SET name = $1, description = $2, parent_id = $3 WHERE id = $4", skill.Name, skill.Description, skill.ParentId, skill.Id)
	return err
}

func (repo *SkillRepository) Delete(skill schemas.Skill) error {
	_, err := repo.db.Exec("DELETE FROM skills WHERE id = $1", skill.Id)
	return err
}

func (repo *SkillRepository) GetAll() (skills []schemas.Skill, err error) {
	rows, err := repo.db.Query("SELECT id, name, description, parent_id FROM skills")
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
		var skill schemas.Skill
		if err := rows.Scan(&skill.Id, &skill.Name, &skill.Description, &skill.ParentId); err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	if err := rows.Err(); err != nil {
		// Here's not nil, err. I'm not sure why, but:
		// https://go.dev/doc/database/querying
		// documentation has the same issue

		return skills, err
	}

	return skills, nil
}

func (repo *SkillRepository) GetAllReadable() (skills []schemas.SkillReadable, err error) {
	q := `SELECT s1.name, s1.description, s2.name FROM skills s1 left join skills s2 on s2.id = s1.parent_id`
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
		var skill schemas.SkillReadable
		if err := rows.Scan(&skill.Name, &skill.Description, &skill.ParentName); err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	if err := rows.Err(); err != nil {
		// Here's not nil, err. I'm not sure why, but:
		// https://go.dev/doc/database/querying
		// documentation has the same issue

		return skills, err
	}

	return skills, nil
}

func (repo *SkillRepository) Get(id int) (skill schemas.Skill, err error) {
	row := repo.db.QueryRow("SELECT id, name, description, parent_id FROM skills WHERE id = $1", id)
	err = row.Scan(&skill.Id, &skill.Name, &skill.Description, &skill.ParentId)
	return skill, err
}

func (repo *SkillRepository) GetConflicts(id int) (conflicts []schemas.SkillConflict, err error) {
	q := `SELECT id, skill_1_id, skill_2_id, comment, priority 
	FROM skill_conflicts
	WHERE skill_1_id = $1 OR skill_2_id = $1`

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
		var conflict schemas.SkillConflict
		if err := rows.Scan(&conflict.Id, &conflict.Skill1Id, &conflict.Skill2Id, &conflict.Comment, &conflict.Priority); err != nil {
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

func (repo *SkillRepository) GetAllConflicts() (conflicts []schemas.SkillConflict, err error) {
	q := `SELECT id, skill_1_id, skill_2_id, comment, priority 
	FROM skill_conflicts`

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
		var conflict schemas.SkillConflict
		if err := rows.Scan(&conflict.Id, &conflict.Skill1Id, &conflict.Skill2Id, &conflict.Comment, &conflict.Priority); err != nil {
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

func (repo *SkillRepository) GetAllConflictsReadable() (conflicts []schemas.SkillConflictReadable, err error) {
	q := `SELECT s1.name, s2.name, comment, priority 
	FROM skill_conflicts sd
	join skills s1 on sd.skill_1_id = s1.id
	join skills s2 on sd.skill_2_id = s2.id`

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
		var conflict schemas.SkillConflictReadable
		if err := rows.Scan(&conflict.Skill1Name, &conflict.Skill2Name, &conflict.Comment, &conflict.Priority); err != nil {
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

func (repo *SkillRepository) CreateConflict(conflict schemas.SkillConflict) (newId int, err error) {
	q := `INSERT INTO skill_conflicts(skill_1_id, skill_2_id, comment, priority) VALUES($1, $2, $3, $4) RETURNING id`

	err = repo.db.QueryRow(
		q, conflict.Skill1Id, conflict.Skill2Id, conflict.Comment, conflict.Priority,
	).Scan(&newId)
	if err != nil {
		log.Println("Error creating skillConflict in skillConflict repository: ", err)
		return 0, err
	}

	return newId, nil
}

func (repo *SkillRepository) CreateConflictIfNotExists(schema schemas.SkillConflictReadable) (created bool, err error) {
	q := `INSERT INTO skill_conflicts(skill_1_id, skill_2_id, comment, priority) 
	SELECT s1.id, s2.id, $3, $4
	FROM skills s1
	JOIN skills s2 ON s2.name = $2::text
	WHERE 
	    s1.name = $1::text 
	    AND NOT EXISTS (
		SELECT 1 
		FROM skill_conflicts sc
		WHERE sc.skill_1_id = s1.id
		AND sc.skill_2_id = s2.id 
		);`

	r, err := repo.db.Exec(q, schema.Skill1Name, schema.Skill2Name, schema.Comment, schema.Priority)

	if err != nil {
		log.Println("Error creating skill conflict: ", err)

		return false, err
	}

	if i, _ := r.RowsAffected(); i != 0 {
		return true, nil
	}

	return false, nil
}

func (repo *SkillRepository) UpdateConflict(conflict schemas.SkillConflict) error {
	q := `UPDATE skill_conflicts SET skill_1_id = $1, skill_2_id = $2, comment = $3, priority = $4 WHERE id = $5`
	_, err := repo.db.Exec(
		q,
		conflict.Skill1Id,
		conflict.Skill2Id,
		conflict.Comment,
		conflict.Priority,
		conflict.Id,
	)
	return err
}

func (repo *SkillRepository) DeleteConflict(conflict schemas.SkillConflict) error {
	q := `DELETE FROM skill_conflicts WHERE id = $1`
	_, err := repo.db.Exec(q, conflict.Id)
	return err
}

func (repo *SkillRepository) GetDomains(id int) (skillDomains []schemas.SkillDomain, err error) {
	q := `SELECT id, skill_id, domain_id, comments, priority 
	FROM skill_domains
	WHERE skill_id = $1`

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
		var skillDomain schemas.SkillDomain
		err = rows.Scan(
			&skillDomain.Id,
			&skillDomain.SkillId,
			&skillDomain.DomainId,
			&skillDomain.Comments,
			&skillDomain.Priority,
		)
		if err != nil {
			return nil, err
		}
		skillDomains = append(skillDomains, skillDomain)
	}

	if err := rows.Err(); err != nil {

		return skillDomains, err
	}

	return skillDomains, nil
}

func (repo *SkillRepository) GetAllDomains() (skillDomains []schemas.SkillDomain, err error) {
	q := `SELECT id, skill_id, domain_id, comments, priority 
	FROM skill_domains`

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
		var skillDomain schemas.SkillDomain
		err = rows.Scan(
			&skillDomain.Id,
			&skillDomain.SkillId,
			&skillDomain.DomainId,
			&skillDomain.Comments,
			&skillDomain.Priority,
		)
		if err != nil {
			return nil, err
		}
		skillDomains = append(skillDomains, skillDomain)
	}

	if err := rows.Err(); err != nil {

		return skillDomains, err
	}

	return skillDomains, nil
}

func (repo *SkillRepository) GetAllDomainsReadable() (skillDomains []schemas.SkillDomainReadable, err error) {
	q := `SELECT s.name, d.name, sd.comments, sd.priority 
	FROM skill_domains sd
	join skills s on sd.skill_id = s.id
	join domains d on sd.domain_id = d.id`

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
		var skillDomain schemas.SkillDomainReadable
		err = rows.Scan(
			&skillDomain.SkillName,
			&skillDomain.DomainName,
			&skillDomain.Comments,
			&skillDomain.Priority,
		)
		if err != nil {
			return nil, err
		}
		skillDomains = append(skillDomains, skillDomain)
	}

	if err := rows.Err(); err != nil {

		return skillDomains, err
	}

	return skillDomains, nil
}

func (repo *SkillRepository) CreateDomains(skillDomain schemas.SkillDomain) (newId int, err error) {
	q := `INSERT INTO skill_domains(skill_id, domain_id, comments, priority) VALUES($1, $2, $3, $4) RETURNING id`

	err = repo.db.QueryRow(
		q, skillDomain.SkillId, skillDomain.DomainId, skillDomain.Comments, skillDomain.Priority,
	).Scan(&newId)
	if err != nil {
		log.Println("Error creating skillDomain in skillDomain repository: ", err)
		return 0, err
	}

	return newId, nil
}

func (repo *SkillRepository) CreateDomainIfNotExists(schema schemas.SkillDomainReadable) (created bool, err error) {
	q := `INSERT INTO skill_domains(skill_id, domain_id, comments, priority) 
	SELECT s.id, d.id, $3, $4
	FROM skills s
	JOIN domains d ON d.name = $2::text
	WHERE 
	    s.name = $1::text 
	    AND NOT EXISTS (
		SELECT 1 
		FROM skill_domains sd
		WHERE sd.skill_id = s.id
		AND sd.domain_id = d.id 
		);`

	r, err := repo.db.Exec(q, schema.SkillName, schema.DomainName, schema.Comments, schema.Priority)

	if err != nil {
		log.Println("Error creating skill domain: ", err)

		return false, err
	}

	if i, _ := r.RowsAffected(); i != 0 {
		return true, nil
	}

	return false, nil
}

func (repo *SkillRepository) UpdateDomains(skillDomain schemas.SkillDomain) error {
	q := `UPDATE skill_domains SET skill_id = $1, domain_id = $2, comments = $3, priority = $4 WHERE id = $5`
	_, err := repo.db.Exec(
		q,
		skillDomain.SkillId,
		skillDomain.DomainId,
		skillDomain.Comments,
		skillDomain.Priority,
		skillDomain.Id,
	)
	return err
}

func (repo *SkillRepository) DeleteDomains(skillDomain schemas.SkillDomain) error {
	q := `DELETE FROM skill_domains WHERE id = $1`
	_, err := repo.db.Exec(q, skillDomain.Id)
	return err
}

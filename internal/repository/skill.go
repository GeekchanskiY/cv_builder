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

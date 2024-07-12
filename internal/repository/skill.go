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
	new_id := 0
	err := repo.db.QueryRow(
		"INSERT INTO skills(name, description, parent_id) VALUES($1, $2, $3) RETURNING id",
		skill.Name, skill.Description, skill.ParentId,
	).Scan(&new_id)

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

	return int(new_id), nil
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
	defer rows.Close()

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

	defer rows.Close()

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

func (repo *SkillRepository) CreateConflict(conflict schemas.SkillConflict) (new_id int, err error) {
	q := `INSERT INTO skill_conflicts(skill_1_id, skill_2_id, comment, priority) VALUES($1, $2, $3, $4) RETURNING id`

	new_id = 0
	err = repo.db.QueryRow(
		q, conflict.Skill1Id, conflict.Skill2Id, conflict.Comment, conflict.Priority,
	).Scan(&new_id)
	if err != nil {
		log.Println("Error creating skill in skill repository: ", err)
		return 0, err
	}

	return int(new_id), nil
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

func (repo *SkillRepository) GetByVacancyId(id int) (skills []schemas.Skill, err error) {
	q := `SELECT s.id, s.name, s.description, s.parent_id FROM skills s
	JOIN vacancy_skills vs ON s.id = vs.skill_id
	WHERE vs.vacancy_id = $1`

	rows, err := repo.db.Query(q, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

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

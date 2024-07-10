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

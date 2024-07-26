package requests

type BuildRequest struct {
	UserID    int `json:"user_id"`
	VacancyID int `json:"vacancy_id"`
}

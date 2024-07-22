package schemas

type BuildRequest struct {
	IsMicroservice bool `json:"is_microservice"`
	UserID         int  `json:"user_id"`
	VacancyID      int  `json:"vacancy_id"`
}

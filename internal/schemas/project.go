package schemas

type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

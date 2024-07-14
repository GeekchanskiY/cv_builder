package schemas

type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectDomain struct {
	Id        int    `json:"id"`
	DomainId  int    `json:"domain_id"`
	ProjectId int    `json:"project_id"`
	Comments  string `json:"comments"`
}

type ProjectRespobsibilities struct {
	Id               int `json:"id"`
	ResponsibilityId int `json:"responsibility_id"`
	ProjectId        int `json:"project_id"`
	Priority         int `json:"priority"`
}

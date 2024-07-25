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

type ProjectDomainReadable struct {
	DomainName  string `json:"domain_name"`
	ProjectName string `json:"project_name"`
	Comments    string `json:"comments"`
}

type ProjectService struct {
	Id          int    `json:"id"`
	ProjectId   int    `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectServiceReadable struct {
	ProjectName string `json:"project_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

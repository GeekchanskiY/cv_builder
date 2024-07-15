package schemas

type Company struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`

	// Used to check if this company may be mentioned in a CV
	IsTrusted bool `json:"is_trusted"`
}

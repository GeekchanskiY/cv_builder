package schemas

type Employee struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	// Used for CV decoration
	AboutMe  string `json:"about_me"`
	ImageUrl string `json:"image_url"`

	// Used to notify if difference in experience is too large
	RealExperience int `json:"real_experience"`
}

package schemas

type Employee struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	// Used to notify if difference in experience is too large
	RealExperience int `json:"real_experience"`
}

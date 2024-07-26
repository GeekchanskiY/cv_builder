package interfaces

type Validatable interface {
	// Validate must return list of field names, where errors found
	Validate() []string
}

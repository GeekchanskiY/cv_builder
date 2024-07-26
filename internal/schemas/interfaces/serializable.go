package interfaces

type Serializable interface {
	// Serialize must return list of field names, where errors found
	Serialize() []string
}

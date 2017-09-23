package declension

// Declension represents a declension class.
type Declension int

//go:generate jsonenums -type=Declension
//go:generate stringer -type=Declension
const (
	Strong Declension = iota
	Weak   Declension = iota
)

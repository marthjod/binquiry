package cases

// Case is an enum representing case names.
type Case int

//go:generate jsonenums -type=Case
//go:generate stringer -type=Case
const (
	Nominative Case = iota
	Accusative Case = iota
	Dative     Case = iota
	Genitive   Case = iota
	Unknown    Case = iota
)

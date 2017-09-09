package wordtype

type Number int

//go:generate stringer -type=Number
const (
	Eintala   Number = iota
	Fleirtala Number = iota
)

package number

// Number is an enum representing grammatical number.
type Number int

//go:generate jsonenums -type=Number
//go:generate stringer -type=Number
const (
	Singular Number = iota
	Plural   Number = iota
)

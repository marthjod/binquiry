package wordtype

type Case int

//go:generate stringer -type=Case
const (
	Nefnifall  Case = iota
	Þolfall    Case = iota
	Þágufall   Case = iota
	Eignarfall Case = iota
)

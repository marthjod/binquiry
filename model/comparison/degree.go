package comparison

import "strings"

// Degree represents a degree of comparison.
type Degree int

//go:generate jsonenums -type=Degree
//go:generate stringer -type=Degree
const (
	Positive    Degree = iota
	Comparative Degree = iota
	Superlative Degree = iota
	Unknown     Degree = iota
)

var degrees = map[string]Degree{
	"frumstig":   Positive,
	"miðstig":    Comparative,
	"efsta stig": Superlative,
}

// GetDegree determines a Degree type based on the input string.
func GetDegree(header string) Degree {
	for k, v := range degrees {
		if strings.Contains(strings.ToLower(header), k) {
			return v
		}
	}

	return Unknown
}

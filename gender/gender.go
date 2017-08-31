package gender

import (
	"strings"
)

type Gender int

//go:generate stringer -type=Gender
const (
	Masculine Gender = iota
	Feminine  = iota
	Neuter    = iota
	Unknown   = iota
)

var genders = map[string]Gender{
	"karlkyn":   Masculine,
	"kvenkyn":   Feminine,
	"hvorugkyn": Neuter,
}

func GetGender(header string) Gender {
	for k, v := range genders {
		if strings.Contains(strings.ToLower(header), k) {
			return v
		}
	}

	return Unknown
}

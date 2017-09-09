package gender

import (
	"strings"
)

type Gender int

//go:generate jsonenums -type=Gender
//go:generate stringer -type=Gender
const (
	Masculine Gender = iota
	Feminine  Gender = iota
	Neuter    Gender = iota
	Unknown   Gender = iota
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

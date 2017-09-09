package gender

import (
	"strings"
)

type Gender int

//go:generate stringer -type=Gender
const (
	Karlkyn   Gender = iota
	Kvenkyn   Gender = iota
	Hvorugkyn Gender = iota
	Unknown   Gender = iota
)

var genders = map[string]Gender{
	"karlkyn":   Karlkyn,
	"kvenkyn":   Kvenkyn,
	"hvorugkyn": Hvorugkyn,
}

func GetGender(header string) Gender {
	for k, v := range genders {
		if strings.Contains(strings.ToLower(header), k) {
			return v
		}
	}

	return Unknown
}

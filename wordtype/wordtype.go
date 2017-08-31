package wordtype

import (
    "strings"
)

type WordType int

//go:generate stringer -type=WordType
const (
	Noun    WordType = iota
	Unknown = iota
)

var wordTypes = map[string]WordType{
	"nafnorð": Noun,
}

func GetWordType(header string) WordType {
	for k, v := range wordTypes {
		if strings.Contains(strings.ToLower(header), k) {
			return v
		}
	}

	return Unknown
}

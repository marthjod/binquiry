package wordtype

import (
	"strings"
)

type WordType int

//go:generate stringer -type=WordType
const (
	NounType WordType = iota
	Unknown  WordType = iota
)

var wordTypes = map[string]WordType{
	"nafnorð": NounType,
}

func GetWordType(header string) WordType {
	for k, v := range wordTypes {
		if strings.Contains(strings.ToLower(header), k) {
			return v
		}
	}

	return Unknown
}

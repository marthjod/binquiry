package wordtype

import (
	"strings"
)

type Word interface {
	Json() string
	List() []string
}

type WordType int

//go:generate stringer -type=WordType
const (
	Noun      WordType = iota
	Adjective WordType = iota
	Verb      WordType = iota
	Unknown   WordType = iota
)

var wordTypes = map[string]WordType{
	"nafnorð":     Noun,
	"lýsingarorð": Adjective,
	"sagnorð":     Verb,
}

func GetWordType(header string) WordType {
	for k, v := range wordTypes {
		if strings.Contains(strings.ToLower(header), k) {
			return v
		}
	}

	return Unknown
}

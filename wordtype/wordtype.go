package wordtype

import (
	"strings"
)

// Word represents features every word type must exhibit.
type Word interface {
	JSON() string
	List() []string
}

// WordType is an enum representing word types.
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

// GetWordType determines a word type based on the input string.
func GetWordType(header string) WordType {
	for k, v := range wordTypes {
		if strings.Contains(strings.ToLower(header), k) {
			return v
		}
	}

	return Unknown
}

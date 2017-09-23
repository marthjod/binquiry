package convert

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/marthjod/bingo/getter"
	"github.com/marthjod/bingo/model/noun"
	"github.com/marthjod/bingo/model/wordtype"
	"github.com/marthjod/bingo/reader"
	"gopkg.in/xmlpath.v2"
)

// Converter is the interface definition for converters.
type Converter interface {
	Convert(g *getter.Getter, query string) string
}

// BaseConverter is an embeddable struct providing base functionality for other converters.
type BaseConverter struct{}

// Convert converts query input to structs.
func (bc *BaseConverter) Convert(g *getter.Getter, query string) (*wordtype.Words, error) {
	var (
		words                wordtype.Words
		errNotImplementedYet = "not implemented yet"
	)

	responses, err := g.GetWord(query)
	if err != nil {
		return &words, err
	}

	for _, resp := range responses {
		header, wordType, xmlRoot, err := reader.Read(bytes.NewReader(resp))
		if err != nil {
			return &words, fmt.Errorf("failed to parse response from %s: %s", g.WordQuery(query), err)
		}

		switch wordType {
		case wordtype.Noun:
			path := xmlpath.MustCompile("//tr/td[2]")
			word := noun.ParseNoun(header, path.Iter(xmlRoot))
			words = append(words, word)
		case wordtype.Adjective:
			return &words, fmt.Errorf("%s: %s", wordtype.Adjective, errNotImplementedYet)
		case wordtype.Verb:
			return &words, fmt.Errorf("%s: %s", wordtype.Verb, errNotImplementedYet)
		default:
			return &words, errors.New(errNotImplementedYet)
		}

	}

	return &words, nil
}

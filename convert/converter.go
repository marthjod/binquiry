package convert

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/marthjod/binquiry/getter"
	"github.com/marthjod/binquiry/model/adjective"
	"github.com/marthjod/binquiry/model/noun"
	"github.com/marthjod/binquiry/model/wordtype"
	"github.com/marthjod/binquiry/reader"
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
	var words wordtype.Words

	responses, err := g.GetWord(query)
	if err != nil {
		return &words, err
	}

	return convert(responses, g.WordQuery(query))
}

func convert(input [][]byte, origin string) (*wordtype.Words, error) {
	var (
		words                = wordtype.Words{}
		errNotImplementedYet = "not implemented yet"
	)

	if len(input) == 0 {
		return &words, fmt.Errorf("cannot convert empty input")
	}

	for _, resp := range input {
		header, wordType, xmlRoot, err := reader.Read(bytes.NewReader(resp))
		if err != nil {
			return &words, fmt.Errorf("failed to parse input from %s: %s", origin, err)
		}

		switch wordType {
		case wordtype.Noun:
			path := xmlpath.MustCompile("//tr/td[2]")
			word := noun.ParseNoun(header, path.Iter(xmlRoot))
			words = append(words, word)
		case wordtype.Adjective:
			path := xmlpath.MustCompile("//tr/td/span")
			word := adjective.ParseAdjective(path.Iter(xmlRoot))
			words = append(words, word)
		case wordtype.Verb:
			return &words, fmt.Errorf("%s: %s", wordtype.Verb, errNotImplementedYet)
		default:
			return &words, errors.New(errNotImplementedYet)
		}

	}

	return &words, nil
}

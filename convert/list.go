package convert

import (
	"bytes"
	"fmt"

	"github.com/marthjod/bingo/getter"
	"github.com/marthjod/bingo/model/noun"
	"github.com/marthjod/bingo/model/wordtype"
	"github.com/marthjod/bingo/reader"
	"gopkg.in/xmlpath.v2"
)

// ListConverter converts input to simple (Go) lists.
type ListConverter struct{}

// Convert implements the Converter interface.
func (lc *ListConverter) Convert(g *getter.Getter, query string) string {
	var (
		words                wordtype.Words
		errNotImplementedYet = "not implemented yet"
	)

	err := g.GetWord(query)
	if err != nil {
		return err.Error()
	}

	for _, resp := range g.ResponseBodies {
		header, wordType, xmlRoot, err := reader.Read(bytes.NewReader(resp))
		if err != nil {
			return fmt.Sprintf("Failed to parse response from %s: %s", g.WordQuery(query), err)
		}

		switch wordType {
		case wordtype.Noun:
			path := xmlpath.MustCompile("//tr/td[2]")
			word := noun.ParseNoun(header, path.Iter(xmlRoot))
			words = append(words, word)
		case wordtype.Adjective:
			return fmt.Sprint("Adjective: ", errNotImplementedYet)
		case wordtype.Verb:
			return fmt.Sprint("Verb: ", errNotImplementedYet)
		default:
			return errNotImplementedYet
		}

	}

	return words.List()
}

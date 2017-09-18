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

// JSONConverter converts input to a JSON array containing the converted items in JSON format.
type JSONConverter struct{}

// Convert implements the Converter interface.
func (jc *JSONConverter) Convert(g *getter.Getter, query string) string {
	var (
		words                wordtype.Words
		errNotImplementedYet = "not implemented yet"
	)

	err := g.GetWord(query)
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err.Error())
	}

	for _, resp := range g.ResponseBodies {
		header, wordType, xmlRoot, err := reader.Read(bytes.NewReader(resp))
		if err != nil {
			return fmt.Sprintf(`{"error": "failed to parse response from %s: %s`, g.WordQuery(query), err)
		}

		switch wordType {
		case wordtype.Noun:
			path := xmlpath.MustCompile("//tr/td[2]")
			word := noun.ParseNoun(header, path.Iter(xmlRoot))
			words = append(words, word)
		case wordtype.Adjective:
			return fmt.Sprintf(`{"error": "%s: %s"}`, wordtype.Adjective, errNotImplementedYet)
		case wordtype.Verb:
			return fmt.Sprintf(`{"error": "%s: %s"}`, wordtype.Verb, errNotImplementedYet)
		default:
			return fmt.Sprintf(`{"error": "%s"}`, errNotImplementedYet)
		}

	}

	return words.JSON()

}

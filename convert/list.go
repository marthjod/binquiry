package convert

import "github.com/marthjod/bingo/getter"

// ListConverter converts input to simple (Go) lists.
type ListConverter struct {
	BaseConverter
}

// Convert implements the Converter interface.
func (lc *ListConverter) Convert(g *getter.Getter, query string) string {
	words, err := lc.BaseConverter.Convert(g, query)

	if err != nil {
		return err.Error()
	}

	return words.List()
}

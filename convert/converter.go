package convert

import "github.com/marthjod/bingo/getter"

// Converter is the interface definition for converters.
type Converter interface {
	Convert(g *getter.Getter, query string) string
}

package convert

import (
	"fmt"

	"github.com/marthjod/binquiry/getter"
)

// JSONConverter converts input to a JSON array containing the converted items in JSON format.
type JSONConverter struct {
	BaseConverter
}

// Convert implements the Converter interface.
func (jc *JSONConverter) Convert(g *getter.Getter, query string) string {

	words, err := jc.BaseConverter.Convert(g, query)
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err.Error())
	}

	return words.JSON()

}

// ConvertBytes is a convenience wrapper for Convert to be used for http.ResponseWriter vel sim.
func (jc *JSONConverter) ConvertBytes(g *getter.Getter, query string) []byte {
	return []byte(jc.Convert(g, query))
}

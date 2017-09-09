package wordtype

import (
	"testing"
)

var expected = []struct {
	in  string
	out WordType
}{
	{
		in:  "nafnorð",
		out: Noun,
	},
	{
		in:  "lýsingarorð",
		out: Adjective,
	},
	{
		in:  "sagnorð",
		out: Verb,
	},
	{
		in:  "Eitthvað annað",
		out: Unknown,
	},
}

func TestGetWordType(t *testing.T) {
	for _, exp := range expected {
		actual := GetWordType(exp.in)
		if exp.out != actual {
			t.Errorf("Expected: %v, actual: %v", exp.out, actual)
		}
	}
}

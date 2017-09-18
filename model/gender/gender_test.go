package gender

import (
	"testing"
)

var expected = []struct {
	in  string
	out Gender
}{
	{
		in:  "Hvorugkynsnafnorð",
		out: Neuter,
	},
	{
		in:  " Kvenkyn ",
		out: Feminine,
	},
	{
		in:  "Eitthvað annað",
		out: Unknown,
	},
}

func TestGetGender(t *testing.T) {
	for _, exp := range expected {
		actual := GetGender(exp.in)
		if exp.out != actual {
			t.Errorf("Expected: %v, actual: %v", exp.out, actual)
		}
	}
}

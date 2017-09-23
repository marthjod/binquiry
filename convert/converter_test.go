package convert

import (
	"errors"
	"github.com/marthjod/bingo/model/case"
	"github.com/marthjod/bingo/model/gender"
	"github.com/marthjod/bingo/model/noun"
	"github.com/marthjod/bingo/model/number"
	"github.com/marthjod/bingo/model/wordtype"
	"io/ioutil"
	"reflect"
	"testing"
)

var expected = []struct {
	input     [][]byte
	converted *wordtype.Words
	err       error
}{
	{
		input:     [][]byte{},
		converted: &wordtype.Words{},
		err:       errors.New("cannot convert empty input"),
	},
	{
		input: [][]byte{
			[]byte("foo bar"),
		},
		converted: &wordtype.Words{},
		err:       errors.New("failed to parse input from http://query.icelandic.example.com: cannot find header"),
	},
	{
		input: [][]byte{
			fromFile("testdata/penni.xml"),
			fromFile("testdata/kona.xml"),
		},
		converted: &wordtype.Words{
			&noun.Noun{
				Gender: gender.Masculine,
				CaseForms: []noun.CaseForm{
					{Case: cases.Nominative, Number: number.Singular, Form: "penni"},
					{Case: cases.Accusative, Number: number.Singular, Form: "penna"},
					{Case: cases.Dative, Number: number.Singular, Form: "penna"},
					{Case: cases.Genitive, Number: number.Singular, Form: "penna"},
					{Case: cases.Nominative, Number: number.Plural, Form: "pennar"},
					{Case: cases.Accusative, Number: number.Plural, Form: "penna"},
					{Case: cases.Dative, Number: number.Plural, Form: "pennum"},
					{Case: cases.Genitive, Number: number.Plural, Form: "penna"},
				},
			},
			&noun.Noun{
				Gender: gender.Feminine,
				CaseForms: []noun.CaseForm{
					{Case: cases.Nominative, Number: number.Singular, Form: "kona"},
					{Case: cases.Accusative, Number: number.Singular, Form: "konu"},
					{Case: cases.Dative, Number: number.Singular, Form: "konu"},
					{Case: cases.Genitive, Number: number.Singular, Form: "konu"},
					{Case: cases.Nominative, Number: number.Plural, Form: "konur"},
					{Case: cases.Accusative, Number: number.Plural, Form: "konur"},
					{Case: cases.Dative, Number: number.Plural, Form: "konum"},
					{Case: cases.Genitive, Number: number.Plural, Form: "kvenna"},
				},
			},
		},
		err: nil,
	},
	{
		input: [][]byte{
			fromFile("testdata/gamall.xml"),
		},
		converted: &wordtype.Words{},
		err:       errors.New("Adjective: not implemented yet"),
	},
}

func fromFile(path string) []byte {
	content, _ := ioutil.ReadFile(path)
	return content
}

func TestBaseConverter_convert(t *testing.T) {
	for _, exp := range expected {
		actual, err := convert(exp.input, "http://query.icelandic.example.com")

		if !reflect.DeepEqual(err, exp.err) {
			t.Errorf("Errors do not match: expected: %v,\nactual: %v", exp.err, err)
		}

		if !reflect.DeepEqual(actual, exp.converted) {
			t.Errorf("Expected: %v,\nactual: %v", exp.converted.JSON(), actual.JSON())
		}
	}
}

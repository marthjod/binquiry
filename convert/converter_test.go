package convert

import (
	"errors"
	"github.com/marthjod/binquiry/model/case"
	"github.com/marthjod/binquiry/model/gender"
	"github.com/marthjod/binquiry/model/noun"
	"github.com/marthjod/binquiry/model/number"
	"github.com/marthjod/binquiry/model/wordtype"
	"io/ioutil"
	"reflect"
	"testing"
	"github.com/marthjod/binquiry/model/declension"
	"github.com/marthjod/binquiry/model/comparison"
	"github.com/marthjod/binquiry/model/adjective"
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
		converted: &wordtype.Words{
			&adjective.Adjective{
				Type: wordtype.Adjective,
				CaseForms: []adjective.CaseForm{
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Nominative, Number: number.Singular, Gender: gender.Masculine, Form: "gamall"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Nominative, Number: number.Singular, Gender: gender.Feminine, Form: "gömul"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Nominative, Number: number.Singular, Gender: gender.Neuter, Form: "gamalt"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Accusative, Number: number.Singular, Gender: gender.Masculine, Form: "gamlan"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Accusative, Number: number.Singular, Gender: gender.Feminine, Form: "gamla"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Accusative, Number: number.Singular, Gender: gender.Neuter, Form: "gamalt"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Dative, Number: number.Singular, Gender: gender.Masculine, Form: "gömlum"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Dative, Number: number.Singular, Gender: gender.Feminine, Form: "gamalli"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Dative, Number: number.Singular, Gender: gender.Neuter, Form: "gömlu"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Genitive, Number: number.Singular, Gender: gender.Masculine, Form: "gamals"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Genitive, Number: number.Singular, Gender: gender.Feminine, Form: "gamallar"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Genitive, Number: number.Singular, Gender: gender.Neuter, Form: "gamals"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Nominative, Number: number.Plural, Gender: gender.Masculine, Form: "gamlir"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Nominative, Number: number.Plural, Gender: gender.Feminine, Form: "gamlar"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Nominative, Number: number.Plural, Gender: gender.Neuter, Form: "gömul"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Accusative, Number: number.Plural, Gender: gender.Masculine, Form: "gamla"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Accusative, Number: number.Plural, Gender: gender.Feminine, Form: "gamlar"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Accusative, Number: number.Plural, Gender: gender.Neuter, Form: "gömul"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Dative, Number: number.Plural, Gender: gender.Masculine, Form: "gömlum"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Dative, Number: number.Plural, Gender: gender.Feminine, Form: "gömlum"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Dative, Number: number.Plural, Gender: gender.Neuter, Form: "gömlum"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Genitive, Number: number.Plural, Gender: gender.Masculine, Form: "gamalla"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Genitive, Number: number.Plural, Gender: gender.Feminine, Form: "gamalla"},
					{Declension: declension.Strong, Degree: comparison.Positive, Case: cases.Genitive, Number: number.Plural, Gender: gender.Neuter, Form: "gamalla"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Nominative, Number: number.Singular, Gender: gender.Masculine, Form: "gamli"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Nominative, Number: number.Singular, Gender: gender.Feminine, Form: "gamla"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Nominative, Number: number.Singular, Gender: gender.Neuter, Form: "gamla"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Accusative, Number: number.Singular, Gender: gender.Masculine, Form: "gamla"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Accusative, Number: number.Singular, Gender: gender.Feminine, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Accusative, Number: number.Singular, Gender: gender.Neuter, Form: "gamla"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Dative, Number: number.Singular, Gender: gender.Masculine, Form: "gamla"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Dative, Number: number.Singular, Gender: gender.Feminine, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Dative, Number: number.Singular, Gender: gender.Neuter, Form: "gamla"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Genitive, Number: number.Singular, Gender: gender.Masculine, Form: "gamla"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Genitive, Number: number.Singular, Gender: gender.Feminine, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Genitive, Number: number.Singular, Gender: gender.Neuter, Form: "gamla"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Nominative, Number: number.Plural, Gender: gender.Masculine, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Nominative, Number: number.Plural, Gender: gender.Feminine, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Nominative, Number: number.Plural, Gender: gender.Neuter, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Accusative, Number: number.Plural, Gender: gender.Masculine, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Accusative, Number: number.Plural, Gender: gender.Feminine, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Accusative, Number: number.Plural, Gender: gender.Neuter, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Dative, Number: number.Plural, Gender: gender.Masculine, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Dative, Number: number.Plural, Gender: gender.Feminine, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Dative, Number: number.Plural, Gender: gender.Neuter, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Genitive, Number: number.Plural, Gender: gender.Masculine, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Genitive, Number: number.Plural, Gender: gender.Feminine, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Positive, Case: cases.Genitive, Number: number.Plural, Gender: gender.Neuter, Form: "gömlu"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Nominative, Number: number.Singular, Gender: gender.Masculine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Nominative, Number: number.Singular, Gender: gender.Feminine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Nominative, Number: number.Singular, Gender: gender.Neuter, Form: "eldra"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Accusative, Number: number.Singular, Gender: gender.Masculine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Accusative, Number: number.Singular, Gender: gender.Feminine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Accusative, Number: number.Singular, Gender: gender.Neuter, Form: "eldra"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Dative, Number: number.Singular, Gender: gender.Masculine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Dative, Number: number.Singular, Gender: gender.Feminine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Dative, Number: number.Singular, Gender: gender.Neuter, Form: "eldra"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Genitive, Number: number.Singular, Gender: gender.Masculine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Genitive, Number: number.Singular, Gender: gender.Feminine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Genitive, Number: number.Singular, Gender: gender.Neuter, Form: "eldra"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Nominative, Number: number.Plural, Gender: gender.Masculine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Nominative, Number: number.Plural, Gender: gender.Feminine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Nominative, Number: number.Plural, Gender: gender.Neuter, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Accusative, Number: number.Plural, Gender: gender.Masculine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Accusative, Number: number.Plural, Gender: gender.Feminine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Accusative, Number: number.Plural, Gender: gender.Neuter, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Dative, Number: number.Plural, Gender: gender.Masculine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Dative, Number: number.Plural, Gender: gender.Feminine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Dative, Number: number.Plural, Gender: gender.Neuter, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Genitive, Number: number.Plural, Gender: gender.Masculine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Genitive, Number: number.Plural, Gender: gender.Feminine, Form: "eldri"},
					{Declension: declension.Weak, Degree: comparison.Comparative, Case: cases.Genitive, Number: number.Plural, Gender: gender.Neuter, Form: "eldri"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Nominative, Number: number.Singular, Gender: gender.Masculine, Form: "elstur"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Nominative, Number: number.Singular, Gender: gender.Feminine, Form: "elst"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Nominative, Number: number.Singular, Gender: gender.Neuter, Form: "elst"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Accusative, Number: number.Singular, Gender: gender.Masculine, Form: "elstan"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Accusative, Number: number.Singular, Gender: gender.Feminine, Form: "elsta"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Accusative, Number: number.Singular, Gender: gender.Neuter, Form: "elst"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Dative, Number: number.Singular, Gender: gender.Masculine, Form: "elstum"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Dative, Number: number.Singular, Gender: gender.Feminine, Form: "elstri"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Dative, Number: number.Singular, Gender: gender.Neuter, Form: "elstu"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Genitive, Number: number.Singular, Gender: gender.Masculine, Form: "elsts"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Genitive, Number: number.Singular, Gender: gender.Feminine, Form: "elstrar"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Genitive, Number: number.Singular, Gender: gender.Neuter, Form: "elsts"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Nominative, Number: number.Plural, Gender: gender.Masculine, Form: "elstir"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Nominative, Number: number.Plural, Gender: gender.Feminine, Form: "elstar"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Nominative, Number: number.Plural, Gender: gender.Neuter, Form: "elst"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Accusative, Number: number.Plural, Gender: gender.Masculine, Form: "elsta"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Accusative, Number: number.Plural, Gender: gender.Feminine, Form: "elstar"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Accusative, Number: number.Plural, Gender: gender.Neuter, Form: "elst"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Dative, Number: number.Plural, Gender: gender.Masculine, Form: "elstum"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Dative, Number: number.Plural, Gender: gender.Feminine, Form: "elstum"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Dative, Number: number.Plural, Gender: gender.Neuter, Form: "elstum"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Genitive, Number: number.Plural, Gender: gender.Masculine, Form: "elstra"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Genitive, Number: number.Plural, Gender: gender.Feminine, Form: "elstra"},
					{Declension: declension.Strong, Degree: comparison.Superlative, Case: cases.Genitive, Number: number.Plural, Gender: gender.Neuter, Form: "elstra"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Nominative, Number: number.Singular, Gender: gender.Masculine, Form: "elsti"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Nominative, Number: number.Singular, Gender: gender.Feminine, Form: "elsta"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Nominative, Number: number.Singular, Gender: gender.Neuter, Form: "elsta"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Accusative, Number: number.Singular, Gender: gender.Masculine, Form: "elsta"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Accusative, Number: number.Singular, Gender: gender.Feminine, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Accusative, Number: number.Singular, Gender: gender.Neuter, Form: "elsta"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Dative, Number: number.Singular, Gender: gender.Masculine, Form: "elsta"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Dative, Number: number.Singular, Gender: gender.Feminine, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Dative, Number: number.Singular, Gender: gender.Neuter, Form: "elsta"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Genitive, Number: number.Singular, Gender: gender.Masculine, Form: "elsta"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Genitive, Number: number.Singular, Gender: gender.Feminine, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Genitive, Number: number.Singular, Gender: gender.Neuter, Form: "elsta"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Nominative, Number: number.Plural, Gender: gender.Masculine, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Nominative, Number: number.Plural, Gender: gender.Feminine, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Nominative, Number: number.Plural, Gender: gender.Neuter, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Accusative, Number: number.Plural, Gender: gender.Masculine, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Accusative, Number: number.Plural, Gender: gender.Feminine, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Accusative, Number: number.Plural, Gender: gender.Neuter, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Dative, Number: number.Plural, Gender: gender.Masculine, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Dative, Number: number.Plural, Gender: gender.Feminine, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Dative, Number: number.Plural, Gender: gender.Neuter, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Genitive, Number: number.Plural, Gender: gender.Masculine, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Genitive, Number: number.Plural, Gender: gender.Feminine, Form: "elstu"},
					{Declension: declension.Weak, Degree: comparison.Superlative, Case: cases.Genitive, Number: number.Plural, Gender: gender.Neuter, Form: "elstu"},
				},
			},
		},
		err:       nil,
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

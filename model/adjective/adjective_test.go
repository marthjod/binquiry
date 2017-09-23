package adjective

import (
	"github.com/marthjod/binquiry/model/case"
	"github.com/marthjod/binquiry/model/comparison"
	"github.com/marthjod/binquiry/model/declension"
	"github.com/marthjod/binquiry/model/gender"
	"github.com/marthjod/binquiry/model/number"
	"github.com/marthjod/binquiry/model/wordtype"
	"github.com/marthjod/binquiry/reader"
	"gopkg.in/xmlpath.v2"
	"os"
	"reflect"
	"testing"
)

var (
	adj = Adjective{
		Type: wordtype.Adjective,
		CaseForms: []CaseForm{
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
	}
)

func TestAdjective_JSON(t *testing.T) {
	expected := `{
  "type": "Adjective",
  "cases": [
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "gamall"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "gömul"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "gamalt"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "gamlan"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "gamla"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "gamalt"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Dative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "gömlum"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Dative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "gamalli"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Dative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "gömlu"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Masculine",
      "form": "gamals"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Feminine",
      "form": "gamallar"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Neuter",
      "form": "gamals"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "gamlir"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "gamlar"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "gömul"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "gamla"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "gamlar"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "gömul"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Dative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "gömlum"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Dative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "gömlum"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Dative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "gömlum"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Masculine",
      "form": "gamalla"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Feminine",
      "form": "gamalla"
    },
    {
      "declension": "Strong",
      "degree": "Positive",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Neuter",
      "form": "gamalla"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "gamli"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "gamla"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "gamla"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "gamla"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "gamla"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Dative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "gamla"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Dative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Dative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "gamla"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Masculine",
      "form": "gamla"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Feminine",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Neuter",
      "form": "gamla"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Dative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Dative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Dative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Masculine",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Feminine",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Positive",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Neuter",
      "form": "gömlu"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "eldra"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "eldra"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Dative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Dative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Dative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "eldra"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Masculine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Feminine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Neuter",
      "form": "eldra"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Dative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Dative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Dative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Masculine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Feminine",
      "form": "eldri"
    },
    {
      "declension": "Weak",
      "degree": "Comparative",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Neuter",
      "form": "eldri"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "elstur"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "elst"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "elst"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "elstan"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "elsta"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "elst"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Dative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "elstum"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Dative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "elstri"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Dative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "elstu"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Masculine",
      "form": "elsts"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Feminine",
      "form": "elstrar"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Neuter",
      "form": "elsts"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "elstir"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "elstar"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "elst"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "elsta"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "elstar"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "elst"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Dative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "elstum"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Dative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "elstum"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Dative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "elstum"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Masculine",
      "form": "elstra"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Feminine",
      "form": "elstra"
    },
    {
      "declension": "Strong",
      "degree": "Superlative",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Neuter",
      "form": "elstra"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "elsti"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "elsta"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Nominative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "elsta"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "elsta"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Accusative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "elsta"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Dative",
      "number": "Singular",
      "gender": "Masculine",
      "form": "elsta"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Dative",
      "number": "Singular",
      "gender": "Feminine",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Dative",
      "number": "Singular",
      "gender": "Neuter",
      "form": "elsta"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Masculine",
      "form": "elsta"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Feminine",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Genitive",
      "number": "Singular",
      "gender": "Neuter",
      "form": "elsta"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Nominative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Accusative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Dative",
      "number": "Plural",
      "gender": "Masculine",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Dative",
      "number": "Plural",
      "gender": "Feminine",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Dative",
      "number": "Plural",
      "gender": "Neuter",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Masculine",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Feminine",
      "form": "elstu"
    },
    {
      "declension": "Weak",
      "degree": "Superlative",
      "case": "Genitive",
      "number": "Plural",
      "gender": "Neuter",
      "form": "elstu"
    }
  ]
}`
	actual := adj.JSON()
	if expected != actual {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}

func TestParseAdjective(t *testing.T) {
	expected := &adj

	f, err := os.Open("testdata/gamall.xml")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer f.Close()

	_, _, root, _ := reader.Read(f)
	path := xmlpath.MustCompile("//tr/td/span")

	actual := ParseAdjective(path.Iter(root))
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v,\nactual: %v", expected, actual)
	}
}

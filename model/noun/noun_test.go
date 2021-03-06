package noun

import (
	"fmt"
	"github.com/marthjod/binquiry/model/case"
	"github.com/marthjod/binquiry/model/gender"
	"github.com/marthjod/binquiry/model/number"
	"github.com/marthjod/binquiry/reader"
	"gopkg.in/xmlpath.v2"
	"os"
	"reflect"
	"testing"
)

var (
	noun = Noun{
		Gender: gender.Masculine,
		CaseForms: []CaseForm{
			{Case: cases.Nominative, Number: number.Singular, Form: "penni"},
			{Case: cases.Accusative, Number: number.Singular, Form: "penna"},
			{Case: cases.Dative, Number: number.Singular, Form: "penna"},
			{Case: cases.Genitive, Number: number.Singular, Form: "penna"},
			{Case: cases.Nominative, Number: number.Plural, Form: "pennar"},
			{Case: cases.Accusative, Number: number.Plural, Form: "penna"},
			{Case: cases.Dative, Number: number.Plural, Form: "pennum"},
			{Case: cases.Genitive, Number: number.Plural, Form: "penna"},
		},
	}
)

func TestNoun_Json(t *testing.T) {
	expected := `{
  "type": "Noun",
  "gender": "Masculine",
  "cases": [
    {
      "case": "Nominative",
      "number": "Singular",
      "form": "penni"
    },
    {
      "case": "Accusative",
      "number": "Singular",
      "form": "penna"
    },
    {
      "case": "Dative",
      "number": "Singular",
      "form": "penna"
    },
    {
      "case": "Genitive",
      "number": "Singular",
      "form": "penna"
    },
    {
      "case": "Nominative",
      "number": "Plural",
      "form": "pennar"
    },
    {
      "case": "Accusative",
      "number": "Plural",
      "form": "penna"
    },
    {
      "case": "Dative",
      "number": "Plural",
      "form": "pennum"
    },
    {
      "case": "Genitive",
      "number": "Plural",
      "form": "penna"
    }
  ]
}`

	actual := noun.JSON()
	if expected != actual {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}

}

func ExampleNoun_JSON() {

	n := Noun{
		Gender: gender.Masculine,
		CaseForms: []CaseForm{
			{Case: cases.Nominative, Number: number.Singular, Form: "penni"},
			{Case: cases.Accusative, Number: number.Singular, Form: "penna"},
			{Case: cases.Dative, Number: number.Singular, Form: "penna"},
			{Case: cases.Genitive, Number: number.Singular, Form: "penna"},
			{Case: cases.Nominative, Number: number.Plural, Form: "pennar"},
			{Case: cases.Accusative, Number: number.Plural, Form: "penna"},
			{Case: cases.Dative, Number: number.Plural, Form: "pennum"},
			{Case: cases.Genitive, Number: number.Plural, Form: "penna"},
		},
	}
	fmt.Println(n.JSON())
	// Output: {
	//   "type": "Noun",
	//   "gender": "Masculine",
	//   "cases": [
	//     {
	//       "case": "Nominative",
	//       "number": "Singular",
	//       "form": "penni"
	//     },
	//     {
	//       "case": "Accusative",
	//       "number": "Singular",
	//       "form": "penna"
	//     },
	//     {
	//       "case": "Dative",
	//       "number": "Singular",
	//       "form": "penna"
	//     },
	//     {
	//       "case": "Genitive",
	//       "number": "Singular",
	//       "form": "penna"
	//     },
	//     {
	//       "case": "Nominative",
	//       "number": "Plural",
	//       "form": "pennar"
	//     },
	//     {
	//       "case": "Accusative",
	//       "number": "Plural",
	//       "form": "penna"
	//     },
	//     {
	//       "case": "Dative",
	//       "number": "Plural",
	//       "form": "pennum"
	//     },
	//     {
	//       "case": "Genitive",
	//       "number": "Plural",
	//       "form": "penna"
	//     }
	//   ]
	// }
}

func TestParseNoun(t *testing.T) {
	expected := &noun
	f, err := os.Open("testdata/penni.xml")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer f.Close()

	header, _, root, _ := reader.Read(f)
	path := xmlpath.MustCompile("//tr/td[2]")

	actual := ParseNoun(header, path.Iter(root))
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v,\nactual: %v", expected, actual)
	}
}

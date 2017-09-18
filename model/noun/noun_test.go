package noun

import (
	"fmt"
	"github.com/marthjod/bingo/model/case"
	"github.com/marthjod/bingo/model/gender"
	"github.com/marthjod/bingo/model/number"
	"github.com/marthjod/bingo/reader"
	"gopkg.in/xmlpath.v2"
	"os"
	"reflect"
	"testing"
)

var (
	noun = Noun{
		Gender: gender.Masculine,
		CaseForms: []CaseForm{
			{Name: cases.Nominative, Number: number.Singular, Form: "penni"},
			{Name: cases.Accusative, Number: number.Singular, Form: "penna"},
			{Name: cases.Dative, Number: number.Singular, Form: "penna"},
			{Name: cases.Genitive, Number: number.Singular, Form: "penna"},
			{Name: cases.Nominative, Number: number.Plural, Form: "pennar"},
			{Name: cases.Accusative, Number: number.Plural, Form: "penna"},
			{Name: cases.Dative, Number: number.Plural, Form: "pennum"},
			{Name: cases.Genitive, Number: number.Plural, Form: "penna"},
		},
	}
)

func TestNoun_Json(t *testing.T) {
	expected := `{
  "gender": "Masculine",
  "cases": [
    {
      "name": "Nominative",
      "number": "Singular",
      "form": "penni"
    },
    {
      "name": "Accusative",
      "number": "Singular",
      "form": "penna"
    },
    {
      "name": "Dative",
      "number": "Singular",
      "form": "penna"
    },
    {
      "name": "Genitive",
      "number": "Singular",
      "form": "penna"
    },
    {
      "name": "Nominative",
      "number": "Plural",
      "form": "pennar"
    },
    {
      "name": "Accusative",
      "number": "Plural",
      "form": "penna"
    },
    {
      "name": "Dative",
      "number": "Plural",
      "form": "pennum"
    },
    {
      "name": "Genitive",
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

func TestNoun_List(t *testing.T) {
	expected := []string{
		"penni",
		"penna",
		"penna",
		"penna",
		"pennar",
		"penna",
		"pennum",
		"penna",
	}
	actual := noun.List()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, actual: %v", expected, actual)
	}
}

func ExampleNoun_JSON() {

	n := Noun{
		Gender: gender.Masculine,
		CaseForms: []CaseForm{
			{Name: cases.Nominative, Number: number.Singular, Form: "penni"},
			{Name: cases.Accusative, Number: number.Singular, Form: "penna"},
			{Name: cases.Dative, Number: number.Singular, Form: "penna"},
			{Name: cases.Genitive, Number: number.Singular, Form: "penna"},
			{Name: cases.Nominative, Number: number.Plural, Form: "pennar"},
			{Name: cases.Accusative, Number: number.Plural, Form: "penna"},
			{Name: cases.Dative, Number: number.Plural, Form: "pennum"},
			{Name: cases.Genitive, Number: number.Plural, Form: "penna"},
		},
	}
	fmt.Println(n.JSON())
	// Output: {
	//   "gender": "Masculine",
	//   "cases": [
	//     {
	//       "name": "Nominative",
	//       "number": "Singular",
	//       "form": "penni"
	//     },
	//     {
	//       "name": "Accusative",
	//       "number": "Singular",
	//       "form": "penna"
	//     },
	//     {
	//       "name": "Dative",
	//       "number": "Singular",
	//       "form": "penna"
	//     },
	//     {
	//       "name": "Genitive",
	//       "number": "Singular",
	//       "form": "penna"
	//     },
	//     {
	//       "name": "Nominative",
	//       "number": "Plural",
	//       "form": "pennar"
	//     },
	//     {
	//       "name": "Accusative",
	//       "number": "Plural",
	//       "form": "penna"
	//     },
	//     {
	//       "name": "Dative",
	//       "number": "Plural",
	//       "form": "pennum"
	//     },
	//     {
	//       "name": "Genitive",
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

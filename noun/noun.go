package noun

import (
	"encoding/json"
	"fmt"
	"github.com/marthjod/bingo/case"
	"github.com/marthjod/bingo/gender"
	"github.com/marthjod/bingo/number"
	"gopkg.in/xmlpath.v2"
)

// CaseForm represents a single case form, i.e. case name, number, and actual form.
type CaseForm struct {
	Name   cases.Case    `json:"name"`
	Number number.Number `json:"number"`
	Form   string        `json:"form"`
}

// Noun is defined as a combination of a gender and a list of case forms.
type Noun struct {
	Gender    gender.Gender `json:"gender"`
	CaseForms []CaseForm    `json:"cases"`
}

// ParseNoun parses XML input into a Noun struct.
func ParseNoun(header string, iter *xmlpath.Iter) *Noun {
	n := Noun{
		Gender: gender.GetGender(header),
	}
	count := 1
	for iter.Next() {
		node := iter.Node()
		switch count {
		case 1:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   cases.Nominative,
				Number: number.Singular,
				Form:   node.String(),
			})
		case 2:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   cases.Accusative,
				Number: number.Singular,
				Form:   node.String(),
			})
		case 3:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   cases.Dative,
				Number: number.Singular,
				Form:   node.String(),
			})
		case 4:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   cases.Genitive,
				Number: number.Singular,
				Form:   node.String(),
			})
		case 5:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   cases.Nominative,
				Number: number.Plural,
				Form:   node.String(),
			})
		case 6:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   cases.Accusative,
				Number: number.Plural,
				Form:   node.String(),
			})
		case 7:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   cases.Dative,
				Number: number.Plural,
				Form:   node.String(),
			})
		case 8:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   cases.Genitive,
				Number: number.Plural,
				Form:   node.String(),
			})
		}
		count++
	}

	return &n
}

// JSON representation of a Noun.
func (n *Noun) JSON() string {
	j, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err.Error())
	}
	return string(j)
}

// List of a Noun's forms.
func (n *Noun) List() []string {
	l := []string{}
	for _, c := range n.CaseForms {
		l = append(l, c.Form)
	}

	return l
}

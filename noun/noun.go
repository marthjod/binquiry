package noun

import (
	"encoding/json"
	"fmt"
	"github.com/marthjod/bingo/case"
	"github.com/marthjod/bingo/gender"
	"github.com/marthjod/bingo/number"
	"gopkg.in/xmlpath.v2"
)

type CaseForm struct {
	Name   case_.Case    `json:"name"`
	Number number.Number `json:"number"`
	Form   string        `json:"form"`
}

type Noun struct {
	Gender    gender.Gender `json:"gender"`
	CaseForms []CaseForm    `json:"cases"`
}

func ParseNoun(header string, iter *xmlpath.Iter) *Noun {
	n := Noun{
		Gender: gender.GetGender(header),
	}
	count := 1
	for {
		if !iter.Next() {
			break
		}
		node := iter.Node()
		switch count {
		case 1:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   case_.Nominative,
				Number: number.Singular,
				Form:   node.String(),
			})
		case 2:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   case_.Accusative,
				Number: number.Singular,
				Form:   node.String(),
			})
		case 3:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   case_.Dative,
				Number: number.Singular,
				Form:   node.String(),
			})
		case 4:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   case_.Genitive,
				Number: number.Singular,
				Form:   node.String(),
			})
		case 5:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   case_.Nominative,
				Number: number.Plural,
				Form:   node.String(),
			})
		case 6:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   case_.Accusative,
				Number: number.Plural,
				Form:   node.String(),
			})
		case 7:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   case_.Dative,
				Number: number.Plural,
				Form:   node.String(),
			})
		case 8:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   case_.Genitive,
				Number: number.Plural,
				Form:   node.String(),
			})
		}
		count++
	}

	return &n
}

func (n *Noun) Json() string {
	j, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err.Error())
	}
	return string(j)
}

func (n *Noun) List() []string {
	l := []string{}
	for _, c := range n.CaseForms {
		l = append(l, c.Form)
	}

	return l
}

package wordtype

import (
	"gopkg.in/xmlpath.v2"
)

type CaseForm struct {
	Name   Case
	Number Number
	Form   string
}

type Noun struct {
	CaseForms []CaseForm
}

func ParseNoun(iter *xmlpath.Iter) *Noun {
	n := Noun{}
	count := 1
	for {
		if !iter.Next() {
			break
		}
		node := iter.Node()
		switch count {
		case 1:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   Nefnifall,
				Number: Eintala,
				Form:   node.String(),
			})
		case 2:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   Þolfall,
				Number: Eintala,
				Form:   node.String(),
			})
		case 3:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   Þágufall,
				Number: Eintala,
				Form:   node.String(),
			})
		case 4:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   Eignarfall,
				Number: Eintala,
				Form:   node.String(),
			})
		case 5:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   Nefnifall,
				Number: Fleirtala,
				Form:   node.String(),
			})
		case 6:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   Þolfall,
				Number: Fleirtala,
				Form:   node.String(),
			})
		case 7:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   Þágufall,
				Number: Fleirtala,
				Form:   node.String(),
			})
		case 8:
			n.CaseForms = append(n.CaseForms, CaseForm{
				Name:   Eignarfall,
				Number: Fleirtala,
				Form:   node.String(),
			})
		}
		count++
	}

	return &n
}

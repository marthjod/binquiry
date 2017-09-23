package adjective

import (
	"encoding/json"
	"fmt"
	"github.com/marthjod/binquiry/model/case"
	"github.com/marthjod/binquiry/model/comparison"
	"github.com/marthjod/binquiry/model/declension"
	"github.com/marthjod/binquiry/model/gender"
	"github.com/marthjod/binquiry/model/number"
	"github.com/marthjod/binquiry/model/wordtype"
	"gopkg.in/xmlpath.v2"
)

const (
	degrees          = 3
	declensions      = 2
	genders          = 3
	numCases         = 4
	numbers          = 2
	afterPositive    = declensions * genders * numCases * numbers
	afterComparative = afterPositive + afterPositive/2
)

// CaseForm represents a single case form with all its attributes.
type CaseForm struct {
	Declension declension.Declension `json:"declension"`
	Degree     comparison.Degree     `json:"degree"`
	Case       cases.Case            `json:"case"`
	Number     number.Number         `json:"number"`
	Gender     gender.Gender         `json:"gender"`
	Form       string                `json:"form"`
}

// Adjective is defined as a combination of adjective case forms.
type Adjective struct {
	Type      wordtype.WordType `json:"type"`
	CaseForms []CaseForm        `json:"cases"`
}

// ParseAdjective parses XML input into an Adjective struct.
func ParseAdjective(iter *xmlpath.Iter) *Adjective {
	var (
		adj = Adjective{
			Type: wordtype.Adjective,
		}
		degreeCounter     comparison.Degree
		declensionCounter declension.Declension
		genderCounter     gender.Gender
		caseCounter       cases.Case
		numberCounter     number.Number
		formCounter       int
	)

	for iter.Next() {

		caseForm := CaseForm{
			Degree:     degreeCounter % degrees,
			Declension: declensionCounter % declensions,
			Case:       caseCounter % numCases,
			Gender:     genderCounter % genders,
			Number:     numberCounter % numbers,
			Form:       iter.Node().String(),
		}

		// override for comparative
		if caseForm.Degree == comparison.Comparative {
			caseForm.Declension = declension.Weak
		}

		// fmt.Printf("(%3d) %+v\n", formCounter+1, caseForm)

		formCounter++
		genderCounter++

		switch formCounter {
		case afterPositive:
			degreeCounter++
		case afterComparative:
			// override for comparative
			declensionCounter--
			degreeCounter++
		}

		switch formCounter % (numCases * genders * numbers) {
		case 0:
			declensionCounter++
		}

		switch formCounter % (numCases * genders) {
		case 0:
			numberCounter++
		}

		switch formCounter % genders {
		case 0:
			caseCounter++
		}

		adj.CaseForms = append(adj.CaseForms, caseForm)
	}

	return &adj
}

// JSON representation of an Adjective.
func (a *Adjective) JSON() string {
	j, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err.Error())
	}
	return string(j)
}

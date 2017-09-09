// generated by jsonenums -type=Case; DO NOT EDIT

package case_

import (
	"encoding/json"
	"fmt"
)

var (
	_CaseNameToValue = map[string]Case{
		"Nominative": Nominative,
		"Accusative": Accusative,
		"Dative":     Dative,
		"Genitive":   Genitive,
	}

	_CaseValueToName = map[Case]string{
		Nominative: "Nominative",
		Accusative: "Accusative",
		Dative:     "Dative",
		Genitive:   "Genitive",
	}
)

func init() {
	var v Case
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_CaseNameToValue = map[string]Case{
			interface{}(Nominative).(fmt.Stringer).String(): Nominative,
			interface{}(Accusative).(fmt.Stringer).String(): Accusative,
			interface{}(Dative).(fmt.Stringer).String():     Dative,
			interface{}(Genitive).(fmt.Stringer).String():   Genitive,
		}
	}
}

// MarshalJSON is generated so Case satisfies json.Marshaler.
func (r Case) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _CaseValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Case: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Case satisfies json.Unmarshaler.
func (r *Case) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Case should be a string, got %s", data)
	}
	v, ok := _CaseNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Case %q", s)
	}
	*r = v
	return nil
}

// generated by jsonenums -type=Number; DO NOT EDIT

package number

import (
	"encoding/json"
	"fmt"
)

var (
	_NumberNameToValue = map[string]Number{
		"Singular": Singular,
		"Plural":   Plural,
	}

	_NumberValueToName = map[Number]string{
		Singular: "Singular",
		Plural:   "Plural",
	}
)

func init() {
	var v Number
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_NumberNameToValue = map[string]Number{
			interface{}(Singular).(fmt.Stringer).String(): Singular,
			interface{}(Plural).(fmt.Stringer).String():   Plural,
		}
	}
}

// MarshalJSON is generated so Number satisfies json.Marshaler.
func (r Number) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _NumberValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Number: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Number satisfies json.Unmarshaler.
func (r *Number) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Number should be a string, got %s", data)
	}
	v, ok := _NumberNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Number %q", s)
	}
	*r = v
	return nil
}
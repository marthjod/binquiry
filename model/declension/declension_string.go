// Code generated by "stringer -type=Declension"; DO NOT EDIT.

package declension

import "fmt"

const _Declension_name = "StrongWeak"

var _Declension_index = [...]uint8{0, 6, 10}

func (i Declension) String() string {
	if i < 0 || i >= Declension(len(_Declension_index)-1) {
		return fmt.Sprintf("Declension(%d)", i)
	}
	return _Declension_name[_Declension_index[i]:_Declension_index[i+1]]
}

// Code generated by "stringer -type=WordType"; DO NOT EDIT.

package wordtype

import "fmt"

const _WordType_name = "NounAdjectiveUnknown"

var _WordType_index = [...]uint8{0, 4, 13, 20}

func (i WordType) String() string {
	if i < 0 || i >= WordType(len(_WordType_index)-1) {
		return fmt.Sprintf("WordType(%d)", i)
	}
	return _WordType_name[_WordType_index[i]:_WordType_index[i+1]]
}

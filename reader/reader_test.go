package reader

import (
	"errors"
	"github.com/marthjod/bingo/wordtype"
	"gopkg.in/xmlpath.v2"
	"os"
	"strings"
	"testing"
)

var expectedHeader = []struct {
	in    string
	found bool
	out   string
}{
	{
		in:    `<div class="page-header"><h2>penni <small>Karlkynsnafnorð</small></h2></div>`,
		found: true,
		out:   "penni Karlkynsnafnorð",
	},
	{
		in:    `<div class="page-header">penni Karlkynsnafnorð</div>`,
		found: false,
		out:   "",
	},
}

var expectedReader = []struct {
	inputPath  string
	header     string
	wordType   wordtype.WordType
	xmlRootNil bool
	err        error
}{
	{
		inputPath:  "testdata/kona.xml",
		header:     "kona Kvenkynsnafnorð",
		wordType:   wordtype.Noun,
		xmlRootNil: false,
		err:        nil,
	},
	{
		inputPath:  "testdata/kona-no-header.xml",
		header:     "",
		wordType:   wordtype.Unknown,
		xmlRootNil: true,
		err:        errors.New("cannot find header"),
	},
	{
		inputPath:  "testdata/invalid.xml",
		header:     "",
		wordType:   wordtype.Unknown,
		xmlRootNil: true,
		err:        errors.New("XML syntax error on line 1: unexpected EOF"),
	},
}

func TestGetHeader(t *testing.T) {

	for _, exp := range expectedHeader {
		root, err := xmlpath.Parse(strings.NewReader(exp.in))
		if err != nil {
			t.Fatal(err.Error())
		}
		actual, ok := GetHeader(root)
		if ok != exp.found {
			t.Error("Header unexpectedly (not) found.")
		}
		if actual != exp.out {
			t.Errorf("Expected: %v,\nactual: %v", exp.out, actual)
		}
	}
}

func TestRead(t *testing.T) {

	for _, exp := range expectedReader {
		f, err := os.Open(exp.inputPath)
		if err != nil {
			t.Fatal(err.Error())
		}
		defer f.Close()

		header, wordType, root, err := Read(f)
		if header != exp.header {
			t.Errorf("Headers do not match. Expected: %v,\nactual: %v", exp.header, header)
		}
		if wordType != exp.wordType {
			t.Errorf("Word types do not match. Expected: %v,\nactual: %v", exp.wordType, wordType)
		}
		if err != nil && err.Error() != exp.err.Error() {
			t.Errorf("Errors do not match. Expected: %v,\nactual: %v", exp.err.Error(), err.Error())
		}
		if exp.xmlRootNil && root != nil {
			t.Errorf("XML root unexpectedly not nil.")
		}
	}
}

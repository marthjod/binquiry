package reader

import (
	"bytes"
	"errors"
	"github.com/marthjod/bingo/wordtype"
	"gopkg.in/xmlpath.v2"
	"html"
	"io"
	"io/ioutil"
	"strings"
)

func GetHeader(root *xmlpath.Node) (header string, found bool) {
	qHeader := xmlpath.MustCompile("/div/h2")
	return qHeader.String(root)
}

func Read(r io.Reader) (header string, wordType wordtype.WordType, xmlRoot *xmlpath.Node, err error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return "", wordtype.Unknown, nil, err
	}

	buf := bytes.NewBuffer(body)
	escaped := html.UnescapeString(buf.String())

	root, err := xmlpath.Parse(strings.NewReader(escaped))
	if err != nil {
		return "", wordtype.Unknown, nil, err
	}

	header, ok := GetHeader(root)
	if !ok {
		return "", wordtype.Unknown, nil, errors.New("cannot find header")
	}

	wordType = wordtype.GetWordType(header)

	return header, wordType, root, nil
}

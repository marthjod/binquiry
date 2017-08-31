package main

import (
	"fmt"
	"github.com/marthjod/bingo/gender"
	"github.com/marthjod/bingo/wordtype"
	"gopkg.in/xmlpath.v2"
	"os"
)

func main() {
	f, err := os.Open("orð.xml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	root, err := xmlpath.Parse(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	header := xmlpath.MustCompile("/div/h2")
	if h, ok := header.String(root); ok {
		wordType := wordtype.GetWordType(h)
		fmt.Printf("Word type: %s\n", wordType)
		if wordType == wordtype.Noun {
			fmt.Printf("Gender: %s\n", gender.GetGender(h))
		}
	}

	path := xmlpath.MustCompile("//tr/td[2]")

	iter := path.Iter(root)
	for {
		if !iter.Next() {
			break
		}
		node := iter.Node()
		fmt.Println(node.String())
	}
}

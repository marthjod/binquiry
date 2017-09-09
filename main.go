package main

import (
	"flag"
	"fmt"
	"github.com/marthjod/bingo/getter"
	"github.com/marthjod/bingo/noun"
	"github.com/marthjod/bingo/reader"
	"github.com/marthjod/bingo/wordtype"
	"gopkg.in/xmlpath.v2"
	"os"
)

func main() {

	var (
		word         wordtype.Word
		query        = flag.String("q", "orð", "Word to query.")
		urlPrefix    = flag.String("url-prefix", "http://dev.phpbin.ja.is/ajax_leit.php", "Query URL prefix.")
		outputFormat = flag.String("f", "json", "Output format (json|list|plain).")
	)
	flag.Parse()

	g := getter.Getter{URLPrefix: *urlPrefix}
	resp, err := g.GetWord(*query)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	header, wordType, xmlRoot, err := reader.Read(resp.Body)
	if err != nil {
		fmt.Printf("Failed to parse response from %s: %s\n", g.WordQuery(*query), err)
		os.Exit(1)
	}

	// have we got a paradigm or a multiple-choice interstitial?

	switch wordType {
	case wordtype.Noun:
		path := xmlpath.MustCompile("//tr/td[2]")
		word = noun.ParseNoun(header, path.Iter(xmlRoot))
	case wordtype.Adjective:
		fmt.Println("Not implemented yet")
		os.Exit(1)
	case wordtype.Verb:
		fmt.Println("Not implemented yet")
		os.Exit(1)
	default:
		fmt.Println("Unknown word type")
		os.Exit(1)
	}

	switch *outputFormat {
	case "json":
		fmt.Println(word.JSON())
	case "list":
		fmt.Println(word.List())
	case "plain":
		fmt.Println(word)
	default:
		fmt.Println("Unknown output format", *outputFormat)
	}
}

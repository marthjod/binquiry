package main

import (
	"bytes"
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"os"
	"strings"

	"github.com/marthjod/bingo/getter"
	"github.com/marthjod/bingo/noun"
	"github.com/marthjod/bingo/wordtype"
	"gopkg.in/xmlpath.v2"
)

func main() {

	var (
		word         wordtype.Word
		query        = flag.String("q", "orð", "Word to query.")
		urlPrefix    = flag.String("url-prefix", "http://dev.phpbin.ja.is/ajax_leit.php", "Query URL prefix.")
		debug        = flag.Bool("debug", false, "Enable debug output.")
		outputFormat = flag.String("f", "json", "Output format (json|list|plain).")
	)
	flag.Parse()

	g := getter.Getter{UrlPrefix: *urlPrefix}
	resp, err := g.GetWord(*query)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buf := bytes.NewBuffer(body)
	escaped := html.UnescapeString(buf.String())

	if *debug {
		fmt.Println(escaped)
	}

	root, err := xmlpath.Parse(strings.NewReader(escaped))
	if err != nil {
		fmt.Printf("Failed to parse response from %s: %s\n", g.WordQuery(*query), err)
		os.Exit(1)
	}

	// have we got a paradigm or a multiple-choice interstitial?

	qHeader := xmlpath.MustCompile("/div/h2")
	header, ok := qHeader.String(root)
	if !ok {
		fmt.Println("Cannot determine word type")
		os.Exit(1)
	}

	wordType := wordtype.GetWordType(header)
	if *debug {
		fmt.Println("Word type:", wordType)
	}

	switch wordType {
	case wordtype.Noun:
		path := xmlpath.MustCompile("//tr/td[2]")
		word = noun.ParseNoun(header, path.Iter(root))
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
		fmt.Println(word.Json())
	case "list":
		fmt.Println(word.List())
	case "plain":
		fmt.Println(word)
	default:
		fmt.Println("Unknown output format", *outputFormat)
	}
}

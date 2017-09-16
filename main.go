package main

import (
	"bytes"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/marthjod/bingo/getter"
	"github.com/marthjod/bingo/noun"
	"github.com/marthjod/bingo/reader"
	"github.com/marthjod/bingo/wordtype"
	"gopkg.in/xmlpath.v2"
	"os"
)

func main() {

	var (
		words        wordtype.Words
		query        = flag.String("q", "orð", "Word to query.")
		urlPrefix    = flag.String("url-prefix", "http://dev.phpbin.ja.is/ajax_leit.php", "Query URL prefix.")
		outputFormat = flag.String("f", "json", "Output format (json|list|plain).")
		debug        = flag.Bool("debug", false, "Enable debug output.")
	)
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	g := getter.Getter{URLPrefix: *urlPrefix}
	resp, err := g.GetWord(*query)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// errcheck
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	log.Debugf("%d response(s) overall", len(g.ResponseBodies))

	// discard first response if we have more than one (it contains no word data itself)
	if len(g.ResponseBodies) > 1 {
		g.ResponseBodies = g.ResponseBodies[1:]
	}

	for _, resp := range g.ResponseBodies {
		header, wordType, xmlRoot, err := reader.Read(bytes.NewReader(resp))
		if err != nil {
			log.Errorf("failed to parse response from %s: %s", g.WordQuery(*query), err)
			os.Exit(1)
		}

		switch wordType {
		case wordtype.Noun:
			path := xmlpath.MustCompile("//tr/td[2]")
			word := noun.ParseNoun(header, path.Iter(xmlRoot))
			words = append(words, word)
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

	}

	log.Debugf("got %d word(s)", len(words))
	switch *outputFormat {
	case "json":
		fmt.Println(words.JSON())
	case "list":
		for _, word := range words {
			fmt.Println(word.List())
		}
	case "plain":
		fmt.Println(words)
	default:
		fmt.Println("Unknown output format", *outputFormat)
	}

}

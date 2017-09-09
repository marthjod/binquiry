package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/marthjod/bingo/gender"
	"github.com/marthjod/bingo/getter"
	"github.com/marthjod/bingo/wordtype"
	"gopkg.in/xmlpath.v2"
	"html"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	var (
		query     = flag.String("q", "orð", "Word to query.")
		urlPrefix = flag.String("url-prefix", "http://dev.phpbin.ja.is/ajax_leit.php", "Query URL prefix.")
		debug     = flag.Bool("debug", false, "Enable debug output.")
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
		fmt.Println(err)
		os.Exit(1)
	}

	// have we got a paradigm or a multiple-choice interstitial?

	header := xmlpath.MustCompile("/div/h2")
	if h, ok := header.String(root); ok {
		wordType := wordtype.GetWordType(h)
		fmt.Printf("Word type: %s\n", wordType)
		if wordType == wordtype.NounType {
			fmt.Printf("Gender: %s\n", gender.GetGender(h))

			path := xmlpath.MustCompile("//tr/td[2]")
			iter := path.Iter(root)
			n := wordtype.ParseNoun(iter)
			fmt.Printf("%s\n", n)
		}
	}
}

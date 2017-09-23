package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/marthjod/bingo/convert"
	"github.com/marthjod/bingo/getter"
)

func main() {

	var (
		converter    convert.Converter
		query        = flag.String("q", "orð", "Word to query.")
		urlPrefix    = flag.String("url-prefix", "http://dev.phpbin.ja.is/ajax_leit.php", "Query URL prefix.")
		outputFormat = flag.String("f", "list", "Output format (json|list).")
	)

	flag.Parse()

	g := getter.Getter{URLPrefix: *urlPrefix}

	switch *outputFormat {
	case "json":
		converter = &convert.JSONConverter{}
	case "list":
		converter = &convert.ListConverter{}
	default:
		fmt.Printf("Unknown output format: %s", *outputFormat)
		os.Exit(1)
	}

	fmt.Print(converter.Convert(&g, *query))

}

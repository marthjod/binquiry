package main

import (
	"flag"
	"fmt"

	"github.com/marthjod/binquiry/convert"
	"github.com/marthjod/binquiry/getter"
)

func main() {

	var (
		converter    convert.Converter
		query        = flag.String("q", "orð", "Word to query.")
		urlPrefix    = flag.String("url-prefix", "http://dev.phpbin.ja.is/ajax_leit.php", "Query URL prefix.")
	)

	flag.Parse()

	converter = &convert.JSONConverter{}
	fmt.Print(converter.Convert(&getter.Getter{URLPrefix: *urlPrefix}, *query))

}

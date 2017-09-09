package getter

import (
	"fmt"
	"net/http"
	"net/url"
)

// Getter builds query URLs and HTTP requests against a data source.
type Getter struct {
	URLPrefix string
}

// WordQuery builds a URL for querying a word.
func (g *Getter) WordQuery(word string) (query string) {
	v := url.Values{}
	v.Set("q", word)
	return g.URLPrefix + "?" + v.Encode()
}

// IDQuery builds a URL for querying a search ID.
func (g *Getter) IDQuery(id int) (query string) {
	return fmt.Sprintf("%s?id=%d", g.URLPrefix, id)
}

// GetWord makes an HTTP request for a word against the data source.
func (g *Getter) GetWord(word string) (*http.Response, error) {
	return http.Get(g.WordQuery(word))
}

// GetID makes an HTTP request for a search ID against the data source.
func (g *Getter) GetID(id int) (*http.Response, error) {
	return http.Get(g.IDQuery(id))
}

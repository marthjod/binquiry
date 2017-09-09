package getter

import (
	"fmt"
	"net/http"
	"net/url"
)

type Getter struct {
	UrlPrefix string
}

func (g *Getter) queryWord(query string) string {
	v := url.Values{}
	v.Set("q", query)
	return g.UrlPrefix + "?" + v.Encode()
}

func (g *Getter) queryId(id int) string {
	return fmt.Sprintf("%s?id=%d", g.UrlPrefix, id)
}

func (g *Getter) GetWord(word string) (*http.Response, error) {
	return http.Get(g.queryWord(word))
}

func (g *Getter) getId(id int) (*http.Response, error) {
	return http.Get(g.queryId(id))
}

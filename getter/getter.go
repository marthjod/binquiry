package getter

import (
	"fmt"
	"net/http"
	"net/url"
)

type Getter struct {
	UrlPrefix string
}

func (g *Getter) WordQuery(word string) (query string) {
	v := url.Values{}
	v.Set("q", word)
	return g.UrlPrefix + "?" + v.Encode()
}

func (g *Getter) IdQuery(id int) (query string) {
	return fmt.Sprintf("%s?id=%d", g.UrlPrefix, id)
}

func (g *Getter) GetWord(word string) (*http.Response, error) {
	return http.Get(g.WordQuery(word))
}

func (g *Getter) GetId(id int) (*http.Response, error) {
	return http.Get(g.IdQuery(id))
}

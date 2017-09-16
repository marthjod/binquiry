package getter

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/marthjod/bingo/reader"
	"golang.org/x/net/html"
	"gopkg.in/xmlpath.v2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// Getter builds query URLs and HTTP requests against a data source.
type Getter struct {
	URLPrefix      string
	ResponseBodies [][]byte
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
	query := g.WordQuery(word)
	log.Debug("query: ", query)
	r, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	body := reader.Sanitize(b)

	g.ResponseBodies = append(g.ResponseBodies, body)

	addtlResponseData, err := g.dispatch(body)
	if err != nil {
		return nil, err
	}
	g.ResponseBodies = append(g.ResponseBodies, addtlResponseData...)

	return r, err
}

// GetID makes an HTTP request for a search ID against the data source.
func (g *Getter) GetID(id int) (*http.Response, error) {
	query := g.IDQuery(id)
	log.Debug("query: ", query)
	return http.Get(query)
}

func (g *Getter) dispatch(r []byte) (responseBodies [][]byte, err error) {
	var reponseBodies = [][]byte{}

	root, err := xmlpath.Parse(bytes.NewReader(r))
	if err != nil {
		return [][]byte{}, err
	}
	qLinks := xmlpath.MustCompile("/ul/li/strong/a")
	if qLinks.Exists(root) {
		doc, err := html.Parse(bytes.NewReader(r))
		if err != nil {
			return [][]byte{}, err
		}

		links := getLinkNodes(doc)
		for _, link := range links {
			id, err := getSearchID(link)
			if err != nil {
				return [][]byte{}, err
			}
			r, err := g.GetID(id)
			if err != nil {
				return [][]byte{}, err
			}

			body, err := readSanitized(r.Body)
			if err != nil {
				return [][]byte{}, err
			}

			reponseBodies = append(reponseBodies, body)
		}
	}

	return reponseBodies, nil
}

func getSearchID(val string) (int, error) {
	log.Debugf("obtaining search ID from %q", val)
	searchID := regexp.MustCompile(`leit_id\('(\d+)'\)`)
	groups := searchID.FindStringSubmatch(val)
	if len(groups) > 1 {
		return strconv.Atoi(groups[1])
	}
	return 0, errors.New("unable to convert search ID")
}

func getLinkNodes(doc *html.Node) []string {

	var (
		f     func(*html.Node)
		links []string
	)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "onclick" {
					links = append(links, a.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return links
}

func readSanitized(r io.Reader) ([]byte, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return []byte{}, err
	}
	return reader.Sanitize(b), nil
}

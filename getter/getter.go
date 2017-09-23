package getter

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/marthjod/binquiry/reader"
	"golang.org/x/net/html"
	"gopkg.in/xmlpath.v2"
)

// Getter builds query URLs and HTTP requests against a data source.
type Getter struct {
	URLPrefix      string
	responseBodies [][]byte
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
func (g *Getter) GetWord(word string) (responses [][]byte, err error) {
	query := g.WordQuery(word)
	log.Debug("query: ", query)
	r, err := http.Get(query)
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	return g.dispatch(reader.Sanitize(b))
}

// GetID makes an HTTP request for a search ID against the data source.
func (g *Getter) GetID(id int) (*http.Response, error) {
	query := g.IDQuery(id)
	log.Debug("query: ", query)
	return http.Get(query)
}

func (g *Getter) fetchLink(link string, w *sync.WaitGroup) {
	id, err := getSearchID(link)
	if err != nil {
		w.Done()
		return
	}

	r, err := g.GetID(id)
	if err != nil {
		w.Done()
		return
	}

	body, err := readSanitized(r.Body)
	if err != nil {
		w.Done()
		return
	}

	g.responseBodies = append(g.responseBodies, body)
	w.Done()
}

func (g *Getter) fetchLinks(links []string, w *sync.WaitGroup) {
	for _, link := range links {
		go g.fetchLink(link, w)
	}
}

func (g *Getter) dispatch(r []byte) (responses [][]byte, err error) {
	var w sync.WaitGroup

	root, err := xmlpath.Parse(bytes.NewReader(r))
	if err != nil {
		return
	}

	// did we land on a multiple-choice page?
	qLinks := xmlpath.MustCompile("/ul/li/strong/a")
	if qLinks.Exists(root) {
		doc, err := html.Parse(bytes.NewReader(r))
		if err != nil {
			return [][]byte{}, err
		}

		links := getLinkNodes(doc)
		w.Add(len(links))
		g.fetchLinks(links, &w)
		w.Wait()

		return g.responseBodies, nil
	}

	// add original response body if we did not land on a multiple-choice page
	g.responseBodies = append(g.responseBodies, r)

	return g.responseBodies, nil
}

func getSearchID(val string) (int, error) {
	var searchID = regexp.MustCompile(`leit_id\('(\d+)'\)`)
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

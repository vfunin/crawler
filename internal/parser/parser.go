package parser

import (
	"fmt"
	"io"
	netURL "net/url"

	"github.com/pkg/errors"

	"github.com/PuerkitoBio/goquery"
)

type Page interface {
	Parse(url string, reader io.Reader) (Page, error)
	Title() string
	Links() []string
	parseTitle(doc *goquery.Document) string
	parseLinks(doc *goquery.Document, baseURL string) (urls []string, err error)
	formatURL(url string, baseURL string) (string, error)
}

type page struct {
	title string
	links []string
}

func New() Page {
	return &page{title: "", links: nil}
}

// Parse - returns page title with links
func (p *page) Parse(url string, reader io.Reader) (Page, error) {
	var (
		doc       *goquery.Document
		err       error
		parsedURL *netURL.URL
		links     []string
	)

	if doc, err = goquery.NewDocumentFromReader(reader); err != nil {
		return nil, errors.Wrap(err, "parser document creation")
	}

	if parsedURL, err = netURL.Parse(url); err != nil {
		return nil, errors.Wrap(err, "parser url parsing")
	}

	baseURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)

	if links, err = p.parseLinks(doc, baseURL); err != nil {
		return nil, errors.Wrap(err, "parser links list")
	}

	p.title = p.parseTitle(doc)
	p.links = links

	return p, nil
}

func (p *page) Title() string {
	return p.title
}

func (p *page) Links() []string {
	return p.links
}

func (p *page) parseTitle(doc *goquery.Document) string {
	return doc.Find("title").First().Text()
}

func (p *page) parseLinks(doc *goquery.Document, baseURL string) (urls []string, err error) {
	fURL := ""

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if !ok {
			return
		}
		if fURL, err = p.formatURL(url, baseURL); err != nil {
			return
		}
		urls = append(urls, fURL)
	})

	return
}

func (p *page) formatURL(url string, baseURL string) (string, error) {
	parsedURL, err := netURL.Parse(url)
	if err != nil {
		return "", err
	}

	if parsedURL.Host == "" {
		return baseURL + url, err
	}

	if parsedURL.Scheme == "" {
		url = "https:" + url
	}

	return url, nil
}

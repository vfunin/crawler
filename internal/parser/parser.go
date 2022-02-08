package parser

import (
	"fmt"
	"io"
	netURL "net/url"

	"github.com/pkg/errors"

	"github.com/PuerkitoBio/goquery"
)

type Page struct {
	Title string
	Links []string
}

//Parse - returns page title with links
func Parse(url string, reader io.Reader) (*Page, error) {
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

	if links, err = getLinks(doc, baseURL); err != nil {
		return nil, errors.Wrap(err, "parser links list")
	}

	return &Page{
		Title: getTitle(doc),
		Links: links,
	}, nil
}

func getTitle(doc *goquery.Document) string {
	return doc.Find("title").First().Text()
}

func getLinks(doc *goquery.Document, baseURL string) (urls []string, err error) {
	fURL := ""

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, ok := s.Attr("href")
		if !ok {
			return
		}
		if fURL, err = formatURL(url, baseURL); err != nil {
			return
		}
		urls = append(urls, fURL)
	})

	return
}

func formatURL(url string, baseURL string) (string, error) {
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

package parser

import (
	"bufio"
	"os"
	"reflect"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	p := New()
	assert.NotNil(t, p)
}

func TestLinks(t *testing.T) {
	var (
		r   *os.File
		err error
	)

	r, err = os.Open("../../mocks/crawler_crawl_filled.html")
	assert.Nil(t, err)

	p := New()

	p, err = p.Parse("http://test.go", bufio.NewReader(r))

	assert.Nil(t, err)

	var ex []string

	assert.Equal(t, ex, p.Links())
}

func TestParse(t *testing.T) {
	var (
		r   *os.File
		err error
	)

	r, err = os.Open("../../mocks/crawler_crawl_filled.html")
	assert.Nil(t, err)

	p := New()

	_, err = p.Parse("http://test.go", bufio.NewReader(r))

	assert.Nil(t, err)
}

func TestTitle(t *testing.T) {
	var (
		r   *os.File
		err error
	)

	r, err = os.Open("../../mocks/crawler_crawl_filled.html")
	assert.Nil(t, err)

	p := New()

	p, err = p.Parse("http://test.go", bufio.NewReader(r))

	assert.Nil(t, err)

	assert.Equal(t, "Home page", p.Title())
}

func Test_page_formatURL(t *testing.T) {
	type fields struct {
		title string
		links []string
	}

	type args struct {
		uri     string
		baseURL string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "normal case",
			fields: fields{
				title: "Home page",
				links: nil,
			},
			args: args{
				uri:     "/test1",
				baseURL: "http://localhost",
			},
			want:    "http://localhost/test1",
			wantErr: false,
		}, {
			name: "error case",
			fields: fields{
				title: "Home page",
				links: nil,
			},
			args: args{
				uri:     "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require",
				baseURL: "",
			},
			want:    "localhost/test2",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &page{
				title: tt.fields.title,
				links: tt.fields.links,
			}
			got, err := p.formatURL(tt.args.uri, tt.args.baseURL)

			if (got != tt.want || err != nil) && !tt.wantErr {
				t.Errorf("formatURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_page_parseLinks(t *testing.T) {
	var doc *goquery.Document

	r, err := os.Open("../../mocks/crawler_crawl_filled.html")
	assert.Nil(t, err)

	doc, err = goquery.NewDocumentFromReader(bufio.NewReader(r))
	assert.Nil(t, err)

	type fields struct {
		title string
		links []string
	}

	type args struct {
		doc     *goquery.Document
		baseURL string
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUrls []string
		wantErr  bool
	}{
		{
			name: "normal case",
			fields: fields{
				title: "Home page",
				links: nil,
			},
			args: args{
				doc:     doc,
				baseURL: "",
			},
			wantUrls: nil,
			wantErr:  false,
		}, {
			name: "normal case",
			fields: fields{
				title: "Home page",
				links: nil,
			},
			args: args{
				doc:     doc,
				baseURL: "",
			},
			wantUrls: []string{"test"},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &page{
				title: tt.fields.title,
				links: tt.fields.links,
			}
			gotUrls, err := p.parseLinks(tt.args.doc, tt.args.baseURL)

			if (!reflect.DeepEqual(gotUrls, tt.wantUrls) || err != nil) && !tt.wantErr {
				t.Errorf("parseLinks() gotUrls = %v, want %v", gotUrls, tt.wantUrls)
			}
		})
	}
}

func Test_page_parseTitle(t *testing.T) {
	var doc *goquery.Document

	r, err := os.Open("../../mocks/crawler_crawl_filled.html")
	assert.Nil(t, err)

	doc, err = goquery.NewDocumentFromReader(bufio.NewReader(r))
	assert.Nil(t, err)

	type fields struct {
		title string
		links []string
	}

	type args struct {
		doc *goquery.Document
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "normal case",
			fields: fields{
				title: "Home page",
				links: nil,
			},
			args: args{
				doc: doc,
			},
			want:    "Home page",
			wantErr: false,
		},
		{
			name: "error case",
			fields: fields{
				title: "Home page",
				links: nil,
			},
			args: args{
				doc: doc,
			},
			want:    "Home page 2",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &page{
				title: tt.fields.title,
				links: tt.fields.links,
			}
			got := p.parseTitle(tt.args.doc)
			if got != tt.want && !tt.wantErr {
				t.Errorf("parseTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

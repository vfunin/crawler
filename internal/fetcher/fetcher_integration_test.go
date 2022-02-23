//go:build integration
// +build integration

package fetcher

import (
	"bufio"
	"context"
	"html/template"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/vfunin/crawler/internal/config"
	"github.com/vfunin/crawler/internal/parser"
)

func Test_fetcher_Fetch(t *testing.T) {
	cwv := context.WithValue(context.Background(), config.LoggerCtxKey, log.Logger)
	ctx, cancel := context.WithCancel(cwv)

	defer cancel()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		type tData struct {
			Title string
		}

		tmpl, err := template.ParseFiles("../../mocks/crawler_crawl.html")
		assert.Nil(t, err)
		err = tmpl.Execute(w, tData{
			Title: "Home page",
		})
		assert.Nil(t, err)
	})

	addr := "localhost:8181"
	url := "http://" + addr + "/"
	server := &http.Server{Addr: addr, Handler: nil}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			return
		}
	}()

	time.Sleep(1 * time.Second)

	f := New(time.Duration(100) * time.Second)
	p, err := f.Fetch(ctx, url)
	assert.Nil(t, err)

	var r *os.File

	r, err = os.Open("../../mocks/crawler_crawl_filled.html")
	assert.Nil(t, err)

	page := parser.New()

	page, err = page.Parse(url, bufio.NewReader(r))

	assert.Nil(t, err)
	assert.Equal(t, page, p)

	err = server.Shutdown(ctx)

	assert.Nil(t, err)
}

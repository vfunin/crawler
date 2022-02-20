package crawler

import (
	"context"
	"html/template"
	"net/http"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/vfunin/crawler/internal/config"
)

func TestNew(t *testing.T) {
	c := New(0, 0)
	assert.NotNil(t, c, "new request is nil")
	assert.Equal(t, c.GetCnt(), int64(1))
}

func TestCrawl(t *testing.T) {
	log.Level(zerolog.TraceLevel)
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

	addr := "localhost:8080"
	url := "http://" + addr + "/"
	server := &http.Server{Addr: addr, Handler: nil}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			return
		}
	}()

	time.Sleep(time.Second)

	c := New(0, 0)

	errCh := make(chan error)

	go c.Crawl(ctx, cancel, url, false, 0, errCh)

LOOP:
	for {
		select {
		case <-ctx.Done():
			cancel()
		case err := <-errCh:
			assert.Nil(t, err)
		case res := <-c.ResultCh():
			assert.Equal(t, Result{
				URL:   "http://localhost:8080/",
				Title: "Home page",
			}, res)
		default:
			if c.GetCnt() != 0 {
				break
			}

			break LOOP
		}
	}

	err := server.Shutdown(ctx)
	assert.Nil(t, err)
}

func TestDecCnt(t *testing.T) {
	c := New(0, 0)
	assert.Equal(t, c.GetCnt(), int64(1))
	c.DecCnt()
	assert.Equal(t, c.GetCnt(), int64(0))
}

func TestGetCnt(t *testing.T) {
	c := New(0, 0)
	assert.Equal(t, c.GetCnt(), int64(1))
	c.DecCnt()
	assert.Equal(t, c.GetCnt(), int64(0))
	c.IncCnt()
	assert.Equal(t, c.GetCnt(), int64(1))
}

func TestIncCnt(t *testing.T) {
	c := New(0, 0)
	assert.Equal(t, c.GetCnt(), int64(1))
	c.IncCnt()
	assert.Equal(t, c.GetCnt(), int64(2))
}

func TestIncMaxDepth(t *testing.T) {
	c := New(0, 0)
	assert.Equal(t, c.MaxDepth(), uint64(0))
	c.IncMaxDepth(uint64(1))
	assert.Equal(t, c.MaxDepth(), uint64(1))
}

func TestResultCh(t *testing.T) {
	c := New(0, 0)
	assert.NotEqual(t, c.ResultCh(), make(<-chan Result))
}

package printer

import (
	"context"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/vfunin/crawler/internal/config"
	"github.com/vfunin/crawler/internal/crawler"
	"github.com/vfunin/crawler/mocks"
)

func TestNew(t *testing.T) {
	cwv := context.WithValue(context.Background(), config.LoggerCtxKey, log.Logger)
	ctx, cancel := context.WithCancel(cwv)

	defer cancel()

	errCh := make(chan error)
	resCh := make(chan crawler.Result)
	cfg := &mocks.Configuration{}
	cfg.On("DepthIncStep").Return(2)
	cfg.On("MaxDepth").Return(1)
	cfg.On("NeedHelp").Return(false)
	cfg.On("Output").Return("")
	cfg.On("OutputToFile").Return(false)
	cfg.On("ShowHelp").Return(false)
	cfg.On("Timeout").Return(0)
	cfg.On("URL").Return("https://start.url")
	cfg.On("String").Return("")
	cfg.On("WithPanic").Return(false)

	c := &mocks.Crawler{}

	c.On("ResultCh").Return(resCh)
	c.On("GetCnt").Return(int64(0))
	p := New(ctx, cancel, cfg, c, errCh)
	assert.NotNil(t, p)
	p.Print()
}

func ExamplePrint() {
	cwv := context.WithValue(context.Background(), config.LoggerCtxKey, log.Logger)
	ctx, cancel := context.WithCancel(cwv)

	defer cancel()

	errCh := make(chan error)
	resCh := make(chan crawler.Result)
	cfg := &mocks.Configuration{}
	cfg.On("DepthIncStep").Return(2)
	cfg.On("MaxDepth").Return(1)
	cfg.On("NeedHelp").Return(false)
	cfg.On("Output").Return("")
	cfg.On("OutputToFile").Return(false)
	cfg.On("ShowHelp").Return(false)
	cfg.On("Timeout").Return(0)
	cfg.On("URL").Return("https://start.url")
	cfg.On("String").Return("")
	cfg.On("WithPanic").Return(false)

	c := &mocks.Crawler{}

	c.On("ResultCh").Return(resCh)
	c.On("GetCnt").Return(int64(0))

	go func() {
		resCh <- crawler.Result{
			URL:   "http://localhost",
			Title: "Test page",
		}
	}()
	time.Sleep(time.Second)

	p := New(ctx, cancel, cfg, c, errCh)

	p.Print()
	//Output:
	//http://localhost;Test page
}

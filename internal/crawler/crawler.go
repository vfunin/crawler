package crawler

import (
	"context"
	"sync"
	"time"

	"github.com/vfunin/crawler/internal/config"
	"github.com/vfunin/crawler/internal/fetcher"

	"github.com/pkg/errors"

	"github.com/rs/zerolog"
)

type Result struct {
	URL   string
	Title string
}

type Crawler interface {
	Crawl(ctx context.Context, cancel context.CancelFunc, url string, withPanic bool, depth uint64, errCh chan<- error)
	IncMaxDepth(step uint64)
	IncCnt()
	DecCnt()
	GetCnt() int64
	MaxDepth() uint64
	ResultCh() chan Result
}

type crawler struct {
	mu                sync.RWMutex
	result            chan Result
	visited           map[string]struct{}
	maxDepth          uint64
	connectionTimeout int
	cnt               int64
}

func New(depth uint64, connectionTimeout int) Crawler {
	return &crawler{
		mu:                sync.RWMutex{},
		result:            make(chan Result),
		visited:           make(map[string]struct{}),
		maxDepth:          depth,
		connectionTimeout: connectionTimeout,
		cnt:               1,
	}
}

func (c *crawler) IncMaxDepth(step uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.maxDepth += step
}

func (c *crawler) IncCnt() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cnt++
}

func (c *crawler) DecCnt() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cnt--
}

func (c *crawler) GetCnt() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.cnt
}

func (c *crawler) MaxDepth() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.maxDepth
}

func (c *crawler) setVisited(url string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.visited[url] = struct{}{}
}

func (c *crawler) ResultCh() chan Result {
	return c.result
}

func (c *crawler) isVisited(url string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	_, ok := c.visited[url]

	return ok
}

func (c *crawler) canGoDeeper(depth uint64) bool {
	return depth <= c.maxDepth
}

func (c *crawler) recoverAndCount(log zerolog.Logger) {
	c.DecCnt()

	if iErr := recover(); iErr != nil {
		err := errors.New("url: " + iErr.(string))

		log.Err(err).Msg("panic during link parsing")
	}
}

// Crawl - Scans the link for nested links and outputs them to the crawler.Result channel
func (c *crawler) Crawl(ctx context.Context, cancel context.CancelFunc, url string, withPanic bool, depth uint64, errCh chan<- error) {
	log := ctx.Value(config.LoggerCtxKey).(zerolog.Logger)

	defer c.recoverAndCount(log)

	if c.isVisited(url) {
		log.Debug().Msgf("url %s has been visited - skip", url)

		return
	}

	c.setVisited(url)

	select {
	case <-ctx.Done():
		return
	default:
		f := fetcher.New(time.Duration(c.connectionTimeout) * time.Second)

		page, err := f.Fetch(ctx, url)
		if err != nil {
			errCh <- err

			return
		}

		c.result <- Result{
			Title: page.Title(),
			URL:   url,
		}

		if !c.canGoDeeper(depth + 1) {
			log.Debug().Msgf("depth limit reached for url %s", url)

			return
		}

		for _, link := range page.Links() {
			c.IncCnt()

			go c.Crawl(ctx, cancel, link, false, depth+1, errCh)
		}

		if withPanic {
			panic(url)
		}
	}
}

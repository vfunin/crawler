package crawler

import (
	"context"
	"crawler/internal/config"
	"crawler/internal/fetcher"
	"crawler/internal/parser"
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

type Result struct {
	URL   string
	Title string
}

type Crawler struct {
	mu                sync.Mutex
	result            chan Result
	visited           map[string]struct{}
	maxDepth          uint64
	connectionTimeout int
	cnt               uint64
}

func New(depth uint64, connectionTimeout int) *Crawler {
	return &Crawler{
		mu:                sync.Mutex{},
		result:            make(chan Result),
		visited:           make(map[string]struct{}),
		maxDepth:          depth,
		connectionTimeout: connectionTimeout,
		cnt:               1,
	}
}

func (c *Crawler) IncMaxDepth(step uint64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.maxDepth += step
}

func (c *Crawler) IncCnt() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cnt++
}

func (c *Crawler) DecCnt() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cnt--
}

func (c *Crawler) GetCnt() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.cnt
}

func (c *Crawler) setVisited(url string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.visited[url] = struct{}{}
}

func (c *Crawler) ResultCh() <-chan Result {
	return c.result
}

func (c *Crawler) isVisited(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.visited[url]

	return ok
}

func (c *Crawler) canGoDeeper(depth uint64) bool {
	return depth <= c.maxDepth
}

func (c *Crawler) recoverAndCount(log zerolog.Logger) {
	c.DecCnt()

	if iErr := recover(); iErr != nil {
		err := fmt.Errorf("url: %v", iErr)

		log.Err(err).Msg("panic during link parsing")
	}
}

//Crawl - Scans the link for nested links and outputs them to the crawler.Result channel
func (c *Crawler) Crawl(ctx context.Context, cancel context.CancelFunc, url string, withPanic bool, depth uint64, errCh chan<- error) {
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
		page := c.fetchURLAndGetPage(ctx, url, errCh)

		c.result <- Result{
			Title: page.Title,
			URL:   url,
		}

		if !c.canGoDeeper(depth + 1) {
			log.Debug().Msgf("depth limit reached for url %s", url)

			return
		}

		for _, link := range page.Links {
			c.IncCnt()

			go c.Crawl(ctx, cancel, link, false, depth+1, errCh)
		}

		if withPanic {
			panic(url)
		}
	}
}

func (c *Crawler) fetchURLAndGetPage(ctx context.Context, url string, errCh chan<- error) *parser.Page {
	var (
		page *parser.Page
		err  error
	)

	f := fetcher.New(time.Duration(c.connectionTimeout) * time.Second)
	page, err = f.Fetch(ctx, url)

	if err != nil {
		errCh <- err

		return nil
	}

	return page
}

package fetcher

import (
	"context"
	"crawler/internal/parser"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Fetcher struct {
	timeout time.Duration
}

func New(timeout time.Duration) *Fetcher {
	return &Fetcher{timeout: timeout}
}

// Fetch - makes a request for a link and returns a parser.Page with title and a slice of links on the page
func (f Fetcher) Fetch(ctx context.Context, url string) (page *parser.Page, err error) {
	var (
		resp *http.Response
		req  *http.Request
	)

	select {
	case <-ctx.Done():
		return
	default:
		client := &http.Client{ //nolint:exhaustivestruct
			Timeout: f.timeout,
		}
		req, err = http.NewRequestWithContext(ctx, "GET", url, nil)

		if err != nil {
			return nil, errors.Wrap(err, "request")
		}

		resp, err = client.Do(req)

		if err != nil {
			return nil, errors.Wrap(err, "response")
		}
		defer resp.Body.Close()
		page, err = parser.Parse(url, resp.Body)

		if err != nil {
			return nil, errors.Wrap(err, "parsing url")
		}

		return
	}
}

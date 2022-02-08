package printer

import (
	"context"
	"crawler/internal/config"
	"crawler/internal/crawler"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type Printer struct {
	ctx     context.Context
	cancel  context.CancelFunc
	cfg     *config.Configuration
	crawler *crawler.Crawler
	errCh   chan<- error
}

func New(ctx context.Context, cancel context.CancelFunc, cfg *config.Configuration, crawler *crawler.Crawler, errCh chan<- error) *Printer {
	return &Printer{ctx: ctx, cancel: cancel, cfg: cfg, errCh: errCh, crawler: crawler}
}

// Print - outputs links to the console or saves to a file
func (p *Printer) Print() {
	var (
		err        error
		outputFile *os.File
		writer     *csv.Writer
	)

	if p.cfg.OutputToFile() {
		writer, outputFile = p.prepareOutputFile()
		if writer != nil {
			defer outputFile.Close()
			defer writer.Flush()
		}
	}

	for {
		select {
		case <-p.ctx.Done():
			return
		case msg := <-p.crawler.ResultCh():
			if p.cfg.OutputToFile() {
				err = writer.Write([]string{msg.URL, msg.Title})
				if err != nil {
					p.errCh <- errors.Wrap(err, "printer writing message")
					p.cancel()

					return
				}

				break
			}

			fmt.Println(msg.URL, ";", msg.Title)
		default:
			if p.crawler.GetCnt() == 0 {
				p.cancel()
			}
		}
	}
}

func (p *Printer) prepareOutputFile() (writer *csv.Writer, outputFile *os.File) {
	var err error

	if outputFile, err = os.Create(p.cfg.Output); err != nil {
		p.errCh <- errors.Wrap(err, "printer file creation")
		p.cancel()

		return
	}

	writer = csv.NewWriter(outputFile)
	writer.Comma = ';'

	if err = writer.Write([]string{"url", "title"}); err != nil {
		p.errCh <- errors.Wrap(err, "printer writing header")
		p.cancel()

		return
	}

	return
}

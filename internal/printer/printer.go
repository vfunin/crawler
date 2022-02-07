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

func Run(ctx context.Context, cfg *config.Configuration, resCh <-chan crawler.Result, errCh chan<- error) {
	var (
		err        error
		outputFile *os.File
		writer     *csv.Writer
	)

	if cfg.OutputToFile() {
		if outputFile, err = os.Create(cfg.Output); err != nil {
			errCh <- errors.Wrap(err, "printer file creation")

			return
		}
		defer outputFile.Close()
		writer = csv.NewWriter(outputFile)
		writer.Comma = ';'

		if err = writer.Write([]string{"url", "title"}); err != nil {
			errCh <- errors.Wrap(err, "printer writing header")

			return
		}
		defer writer.Flush()
	}

	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-resCh:
			if cfg.OutputToFile() {
				if err = writer.Write([]string{msg.URL, msg.Title}); err != nil {
					errCh <- errors.Wrap(err, "printer writing message")

					return
				}

				break
			}

			fmt.Println(msg.URL, ";", msg.Title)
		}
	}
}

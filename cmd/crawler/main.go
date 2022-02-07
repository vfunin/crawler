package main

import (
	"context"
	"crawler/internal/config"
	"crawler/internal/crawler"
	"crawler/internal/printer"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.Load()

	if cfg.NeedHelp {
		config.ShowHelp()
		os.Exit(0)
	}

	if err != nil {
		log.Fatal().Err(err).Msg("config error")
	}

	ctx, cancel := getContext()

	errCh := make(chan error)

	log.Debug().Msgf("starting with config: %s", cfg)

	c := crawler.New(cfg.MaxDepth, cfg.Timeout)

	log.Trace().Msg("start crawler goroutine")

	go c.Crawl(ctx, cancel, cfg.URL, cfg.WithPanic, 0, errCh)

	log.Trace().Msg("start printer goroutine")

	go printer.Run(ctx, &cfg, c.ResultCh(), errCh)

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT)

	incCh := make(chan os.Signal, 1)
	signal.Notify(incCh, syscall.SIGUSR1)

	log.Debug().Msgf("current process id: %d", os.Getpid())

	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("crawling done")

			return
		case err = <-errCh:
			log.Err(err).Msg("crawling error")
		case <-stopCh:
			log.Info().Msg("start graceful shutdown")
			cancel()
		case <-incCh:
			log.Info().Msgf("maximum depth increased by %d", cfg.DepthIncStep)
			c.IncMaxDepth(uint64(cfg.DepthIncStep))
		}
	}
}

func getContext() (ctx context.Context, cancel context.CancelFunc) {
	cwv := context.WithValue(context.Background(), config.LoggerCtxKey, log.Logger)

	ctx, cancel = context.WithCancel(cwv)

	return
}

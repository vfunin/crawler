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
	cfg := handleConfiguration()

	log.Debug().Msgf("starting with config: %s", cfg)
	log.Debug().Msgf("current process id: %d", os.Getpid())

	ctx, cancel := getContext()
	errCh := make(chan error)

	log.Trace().Msg("start crawler goroutine")

	c := crawler.New(cfg.MaxDepth(), cfg.Timeout())
	go c.Crawl(ctx, cancel, cfg.URL(), cfg.WithPanic(), 0, errCh)

	log.Trace().Msg("start printer goroutine")

	p := printer.New(ctx, cancel, cfg, c, errCh)
	go p.Print()

	listenChannels(ctx, cancel, errCh, cfg.DepthIncStep(), c)
}

func listenChannels(ctx context.Context, cancel context.CancelFunc, errCh <-chan error, depthStep int, c crawler.Crawler) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT)

	incCh := make(chan os.Signal, 1)
	signal.Notify(incCh, syscall.SIGUSR1)

	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("crawling done")

			return
		case err := <-errCh:
			log.Err(err).Msg("crawling error")
		case <-stopCh:
			log.Info().Msg("start graceful shutdown")
			cancel()
		case <-incCh:
			log.Info().Msgf("maximum depth increased by %d", depthStep)
			c.IncMaxDepth(uint64(depthStep))
		}
	}
}

func handleConfiguration() config.Configuration {
	cfg, err := config.New()

	if cfg.NeedHelp() {
		cfg.ShowHelp()
		os.Exit(0)
	}

	if err != nil {
		log.Fatal().Err(err).Msg("config error")
	}

	return cfg
}

func getContext() (ctx context.Context, cancel context.CancelFunc) {
	cwv := context.WithValue(context.Background(), config.LoggerCtxKey, log.Logger)

	ctx, cancel = context.WithCancel(cwv)

	return
}

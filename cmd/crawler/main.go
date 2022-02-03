package main

import (
	"context"
	"crawler/internal/config"
	"crawler/internal/crawler"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	cfg, err := config.Load()

	if err != nil {
		log.Fatal().Err(err).Msg("config parsing error")
	}

	if cfg.NeedHelp {
		config.ShowHelp()
		os.Exit(0)
	}

	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	if cfg.Timeout == 0 {
		ctx, cancel = context.WithCancel(context.Background())
	} else {
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	}

	go crawler.Run(ctx)

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT)

	incCh := make(chan os.Signal, 1)
	signal.Notify(incCh, syscall.SIGUSR1)

	log.Info().Msgf("current process id: %d", os.Getpid())

	for {
		select {
		case <-ctx.Done():
			log.Info().Msg("crawling done")

			return
		case <-stopCh:
			log.Info().Msg("start graceful shutdown")
			cancel()
		case <-incCh:
			log.Info().Msgf("maximum depth increased by %d", cfg.DepthIncStep)
			cancel()
		}
	}
}

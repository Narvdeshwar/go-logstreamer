package main

import (
	"context"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"github.com/Narvdeshwar/go-logstreamer/internal/config"
	"github.com/Narvdeshwar/go-logstreamer/internal/logger"
	"github.com/Narvdeshwar/go-logstreamer/internal/pipeline"
)

func main() {
	cfg := config.Load()

	log := logger.Init(cfg.Debug)

	log.Info().
		Int("workers", cfg.Workers).
		Int("buffer", cfg.Buffer).
		Strs("files", cfg.Files).
		Msg("logstreamer started")

	// Optional CPU profiling
	var stopProfile func()
	if cfg.Profile {
		f, err := os.Create("cpu.out")
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create cpu profile")
		}

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal().Err(err).Msg("failed to start cpu profile")
		}

		stopProfile = func() {
			pprof.StopCPUProfile()
			f.Close()
			log.Info().Msg("cpu profiling stopped")
		}
		defer stopProfile()
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		log.Info().
			Str("signal", sig.String()).
			Msg("shutdown signal received")
		cancel()
	}()

	p := pipeline.New(cfg, log)
	p.Run(ctx)

	log.Info().Msg("logstreamer exited cleanly")
}

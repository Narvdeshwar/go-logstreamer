package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Narvdeshwar/go-logstreamer/internal/config"
	"golang.org/x/text/message/pipeline"
)

func main() {
	cfg := config.Load()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		cancel()
	}()
	p := pipeline.New(cfg)
	p.Run(ctx)
}

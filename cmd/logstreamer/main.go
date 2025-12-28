package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Narvdeshwar/go-logstreamer/internal/config"
	"github.com/Narvdeshwar/go-logstreamer/internal/pipeline"
)

func main() {
	cfg := config.Load() // loading the file with [file_address,workers,buffer]
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		log.Println("Received Signal: ", sig)
		cancel()
	}()
	p := pipeline.New(cfg) // this will return the cfg file address so that single source of data can be maintain
	p.Run(ctx)
}

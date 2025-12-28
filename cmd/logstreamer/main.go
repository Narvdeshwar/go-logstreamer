package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"github.com/Narvdeshwar/go-logstreamer/internal/config"
	"github.com/Narvdeshwar/go-logstreamer/internal/pipeline"
)

func main() {
	cfg := config.Load() // loading the file with [file_address,workers,buffer]
	f, err := os.Create("cpu.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	// pprof.WriteHeapProfile(f)
	defer pprof.StopCPUProfile()

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

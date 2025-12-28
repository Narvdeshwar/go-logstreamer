package pipeline

import (
	"context"
	"sync"
	"time"

	"github.com/Narvdeshwar/go-logstreamer/internal/aggregator"
	"github.com/Narvdeshwar/go-logstreamer/internal/config"
	"github.com/Narvdeshwar/go-logstreamer/internal/output"
	"github.com/Narvdeshwar/go-logstreamer/internal/parser"
	"github.com/Narvdeshwar/go-logstreamer/internal/source"
	"github.com/Narvdeshwar/go-logstreamer/pkg/model"
	"github.com/rs/zerolog"
)

type Pipeline struct {
	cfg *config.Config
	log zerolog.Logger
}

func New(cfg *config.Config, log zerolog.Logger) *Pipeline {
	return &Pipeline{cfg: cfg, log: log}
}

func (p *Pipeline) Run(ctx context.Context) {
	rawChan := make(chan string, p.cfg.Buffer)
	parsedChan := make(chan model.LogEntry, p.cfg.Buffer)

	var srcWG sync.WaitGroup
	var workerWG sync.WaitGroup
	for _, file := range p.cfg.Files {
		srcWG.Add(1)
		go func(path string) {
			defer srcWG.Done()
			src := source.NewFileSource(path)
			_ = src.Start(ctx, rawChan)
		}(file)
	}
	go func() {
		srcWG.Wait()
		close(rawChan)
	}()
	prsr := parser.NewSimpleParser()
	for i := 0; i < p.cfg.Workers; i++ {
		workerWG.Add(1)
		go func() {
			defer workerWG.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case line, ok := <-rawChan:
					if !ok {
						return
					}
					entry, err := prsr.Parse(line)
					if err != nil {
						continue
					}
					select {
					case <-ctx.Done():
						return
					case parsedChan <- *entry:
					}
				}

			}
		}()
	}
	go func() {
		workerWG.Wait()
		close(parsedChan)
	}()
	agg := aggregator.NewAggregator()
	start := time.Now()
	aggDone := make(chan struct{})
	go func() {
		agg.Run(ctx, parsedChan)
		close(aggDone)
	}()
	<-aggDone
	agg.PrintSummary()
	if p.cfg.Output != "" {
		err := output.WriteJSON(p.cfg.Output, agg.Summary())
		if err != nil {
			p.log.Error().
				Err(err).
				Str("output", p.cfg.Output).
				Msg("failed to write JSON summary")
		} else {
			p.log.Info().
				Str("output", p.cfg.Output).
				Msg("JSON summary written")
		}
	}
	elapsed := time.Since(start)
	summary := agg.Summary()
	elapsedSec := elapsed.Seconds()

	var throughput float64
	if elapsedSec > 0 {
		throughput = float64(summary.TotalLines) / elapsedSec
	}

	p.log.Info().
		Dur("elapsed", elapsed).
		Float64("throughput", throughput).
		Msg("pipeline completed")

}

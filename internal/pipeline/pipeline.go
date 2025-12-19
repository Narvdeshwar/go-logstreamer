package pipeline

import (
	"context"
	"sync"

	"github.com/Narvdeshwar/go-logstreamer/internal/config"
	"github.com/Narvdeshwar/go-logstreamer/internal/source"
)

type Pipeline struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Pipeline {
	return &Pipeline{cfg: cfg}
}

func (p *Pipeline) Run(ctx context.Context) {
	rowChan := make(chan string, p.cfg.Buffer)

	var wg sync.WaitGroup
	for _, file := range p.cfg.Files {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			src := source.NewFileSource(path)
			_ = src.Start(ctx, rowChan)
		}(file)
	}
	go func() {
		wg.Wait()
		close(rowChan)
	}()
	for line := range rowChan {
		_ = line
	}
	<-ctx.Done()
}

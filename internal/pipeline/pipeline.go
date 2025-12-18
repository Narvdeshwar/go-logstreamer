package pipeline

import (
	"context"

	"github.com/Narvdeshwar/go-logstreamer/internal/config"
)

type Pipeline struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Pipeline {
	return &Pipeline{cfg: cfg}
}

func (p *Pipeline) Run(ctx context.Context) {
	
	<-ctx.Done()
}

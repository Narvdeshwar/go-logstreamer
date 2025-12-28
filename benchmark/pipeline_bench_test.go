package logstreamer

import (
	"context"
	"runtime"
	"testing"

	"github.com/Narvdeshwar/go-logstreamer/internal/config"
	"github.com/Narvdeshwar/go-logstreamer/internal/logger"
	"github.com/Narvdeshwar/go-logstreamer/internal/pipeline"
)

func BenchmarkPipeline(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping heavy benchmark")
	}

	cfg := &config.Config{
		Files:   []string{"../big.log"},
		Workers: runtime.NumCPU(),
		Buffer:  10000,
		Debug:   false,
	}
	log := logger.Init(false)

	ctx := context.Background()
	p := pipeline.New(cfg, log)

	b.ResetTimer()
	p.Run(ctx)
}

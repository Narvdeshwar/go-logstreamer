package benchmark

import (
	"context"
	"runtime"
	"testing"

	"github.com/Narvdeshwar/go-logstreamer/internal/config"
	"github.com/Narvdeshwar/go-logstreamer/internal/pipeline"
)

func BenchmarkPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cfg := &config.Config{
			Files:   []string{"../big.log"},
			Workers: runtime.NumCPU(),
			Buffer:  10000,
		}
		ctx := context.Background()
		p := pipeline.New(cfg)
		p.Run(ctx)
	}
}

package aggregator

import (
	"context"

	"github.com/Narvdeshwar/go-logstreamer/pkg/model"
)

type Aggregator struct {
	totalLines   int
	levelCount   map[string]int
	serviceCount map[string]int
}

func NewAggregator() *Aggregator {
	return &Aggregator{
		levelCount:   make(map[string]int),
		serviceCount: make(map[string]int),
	}
}

func (a *Aggregator) Run(ctx context.Context, in <-chan model.LogEntry) {
	for {
		select {
		case <-ctx.Done():
			return
		case entry, ok := <-in:
			if !ok {
				return
			}
			a.totalLines++
			a.levelCount[entry.Level]++
			a.serviceCount[entry.Service]++
		}
	}
}

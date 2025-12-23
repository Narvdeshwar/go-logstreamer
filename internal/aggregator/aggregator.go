package aggregator

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

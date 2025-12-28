package aggregator

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

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

func (a *Aggregator) Summary() model.Summary {
	return model.Summary{
		TotalLines:   a.totalLines,
		LevelCount:   a.levelCount,
		ServiceCount: a.serviceCount,
	}
}

func (a *Aggregator) Run(ctx context.Context, in <-chan model.LogEntry) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Aggregator Shutting Down...")
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

func (a *Aggregator) PrintSummary() {
	summary := a.Summary()
	fmt.Println()
	printLine("=", 30)
	fmt.Println("LOG STREAMER SUMMARY")
	printLine("=", 30)
	fmt.Printf("➤  Total Lines Processed: %d\n\n", summary.TotalLines)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	printSection := func(title string, data map[string]int) {
		fmt.Fprintln(w, title+"\tCOUNT")
		fmt.Fprintln(w, strings.Repeat("-", len(title))+"\t-----")
		var keys []string
		for k := range data {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Fprintf(w, "%s\t%d\n", k, data[k])
		}
		fmt.Fprintln(w, "\t")
	}
	printSection("LOG LEVEL", summary.LevelCount)
	printSection("SERVICE", summary.ServiceCount)
	w.Flush()
	printLine("=", 30)
	fmt.Println()
}
func printLine(char string, count int) {
	fmt.Println(strings.Repeat(char, count))
}

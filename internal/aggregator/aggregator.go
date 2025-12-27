package aggregator

import (
	"context"
	"fmt"
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

func (a *Aggregator) PrintSummary() {
	// Header Design
	fmt.Println()
	printLine("=", 30)
	fmt.Println("LOG STREAMER SUMMARY")
	printLine("=", 30)

	// 1. Total Lines
	fmt.Printf("➤  Total Lines Processed: %d\n\n", a.totalLines)

	// Writer setup for alignment (minwidth, tabwidth, padding, padchar, flags)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)

	// 2. Helper function to print sorted maps (Kyuki Map random hota hai)
	printSection := func(title string, data map[string]int) {
		fmt.Fprintln(w, title+"\tCOUNT")
		fmt.Fprintln(w, strings.Repeat("-", len(title))+"\t-----")

		// Keys ko sort kar rahe hain
		var keys []string
		for k := range data {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Fprintf(w, "%s\t%d\n", k, data[k])
		}
		fmt.Fprintln(w, "\t") // Empty line spacing
	}

	// 3. Print Sections
	printSection("LOG LEVEL", a.levelCount)
	printSection("SERVICE", a.serviceCount)

	// Buffer flush karna zaroori hai
	w.Flush()
	printLine("=", 30)
	fmt.Println()
}

// Helper utility for lines
func printLine(char string, count int) {
	fmt.Println(strings.Repeat(char, count))
}

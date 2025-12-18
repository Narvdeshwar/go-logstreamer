package config

import (
	"flag"
	"strings"
)

type Config struct {
	Files   []string
	Workers int
	Buffer  int
	Output  string
}

func Load() *Config {
	files := flag.String("files", "", "comma seperated log files")
	workers := flag.Int("workers", 8, "Number of parser workers")
	buffer := flag.Int("buffer", 10000, "Channel buffer size")
	output := flag.String("	output", "", "Output file (optional)")

	flag.Parse()

	return &Config{
		Files:   splitComma(*files),
		Workers: *workers,
		Buffer:  *buffer,
		Output:  *output,
	}

}
func splitComma(s string) []string {
	if s == " " {
		return nil
	}
	return strings.Split(s, ",")

}

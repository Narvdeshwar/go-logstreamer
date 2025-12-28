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
	Debug   bool
	Profile bool
}

func Load() *Config {
	files := flag.String("files", "", "comma separated log files")
	workers := flag.Int("workers", 8, "Number of parser workers")
	buffer := flag.Int("buffer", 10000, "Channel buffer size")
	output := flag.String("output", "", "Output file (optional)")
	debug := flag.Bool("debug", false, "Enable debug logging")
	profile := flag.Bool("profile", false, "Enable CPU profiling")

	flag.Parse()

	return &Config{
		Files:   splitComma(*files),
		Workers: *workers,
		Buffer:  *buffer,
		Output:  *output,
		Debug:   *debug,
		Profile: *profile,
	}
}

func splitComma(s string) []string {
	if s == " " {
		return nil
	}
	return strings.Split(s, ",")

}

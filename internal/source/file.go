package source

import (
	"bufio"
	"context"
	"os"
)

type FileSource struct {
	Path string
}

func NewFileSource(path string) *FileSource {
	return &FileSource{Path: path}
}

func (f *FileSource) Start(ctx context.Context, out chan<- string) error {
	file, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case out <- scanner.Text():
			// line successfully sent
		}

	}
	return scanner.Err()
}

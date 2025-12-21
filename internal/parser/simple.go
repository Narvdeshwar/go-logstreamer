package parser

import (
	"errors"
	"strings"
	"time"

	"github.com/Narvdeshwar/go-logstreamer/pkg/model"
)

type SimpleParser struct{}

func NewSimpleParser() *SimpleParser {
	return &SimpleParser{}
}

func (p *SimpleParser) Parse(line string) (*model.LogEntry, error) {
	parts := strings.SplitN(line, " ", 4)
	if len(parts) < 4 {
		return nil, errors.New("Invalid log format")
	}
	ts, err := time.Parse("2025-01-02", parts[0])
	if err != nil {
		return nil, errors.New("Time format is incorrect!")
	}
	return &model.LogEntry{
		TimeStamp: ts,
		Level:     parts[1],
		Service:   parts[2],
	}, nil
}

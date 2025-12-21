package parser

import "github.com/Narvdeshwar/go-logstreamer/pkg/model"

type Parser interface {
	Parse(line string) (*model.LogEntry, error)
}

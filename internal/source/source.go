package source

import "context"

type LogSource interface {
	Start(ctx context.Context, out chan<- string) error
}

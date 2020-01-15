package concurrency

import (
	"context"
	"time"
)

// IConcurrent interface for concurrency package
type IConcurrent interface {
	Parallelize(functions ...func()) error
	ParallelizeTimeout(timeout time.Duration, functions ...func()) error
	ParallelizeContext(ctx context.Context, functions ...func()) error
}

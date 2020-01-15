package concurrency

import (
	"context"
	"time"

	"github.com/shomali11/parallelizer"
)

// Concurrent manage all concurrency function
type Concurrent struct{}

// Parallelize parallelizes the function calls
func (c *Concurrent) Parallelize(functions ...func()) error {
	return c.ParallelizeTimeout(0, functions...)
}

// ParallelizeTimeout parallelizes the function calls with a timeout
func (c *Concurrent) ParallelizeTimeout(timeout time.Duration, functions ...func()) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return c.ParallelizeContext(ctx, functions...)
}

// ParallelizeContext parallelizes the function calls with a context
func (c *Concurrent) ParallelizeContext(ctx context.Context, functions ...func()) error {
	group := parallelizer.NewGroup()
	for _, function := range functions {
		group.Add(function)
	}

	return group.Wait(parallelizer.WithContext(ctx))
}

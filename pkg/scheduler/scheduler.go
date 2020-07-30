package scheduler

import (
	"context"
	"sync"
	"time"
)

type JobContext func(ctx context.Context)

type Scheduler struct {
	wg              *sync.WaitGroup
	queueCancelFunc []context.CancelFunc
}

// Init scheduler worker
func NewScheduler() *Scheduler {
	return &Scheduler{
		wg:              new(sync.WaitGroup),
		queueCancelFunc: make([]context.CancelFunc, 0),
	}
}

// Add starts goroutine which job with interval delay
func (s *Scheduler) Add(ctx context.Context, j JobContext, interval time.Duration) {
	ctx, cancel := context.WithCancel(ctx)
	s.queueCancelFunc = append(s.queueCancelFunc, cancel)

	s.wg.Add(1)
	go s.process(ctx, j, interval)
}

// Stop all running jobs
func (s *Scheduler) Stop() {
	for _, cancel := range s.queueCancelFunc {
		cancel()
	}
	s.wg.Wait()
}

func (s *Scheduler) process(ctx context.Context, j JobContext, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			j(ctx)
		case <-ctx.Done():
			s.wg.Done()
			return
		}
	}
}

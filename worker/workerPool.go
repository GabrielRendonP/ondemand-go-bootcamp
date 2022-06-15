package worker

import (
	"context"
	"fmt"
	"sync"
)

var maxWorkers int = 10

type Worker struct {
	id           int
	workCapacity int
}

// Sends job to results channel and shutdown worker on capacity depleted
func (w *Worker) Work(wg *sync.WaitGroup, jobs <-chan []string, results chan<- []string, cancel context.CancelFunc) {
	for job := range jobs {
		// Worker shutdown on capacity depleted
		if w.workCapacity <= 0 {
			results <- job
			cancel()
			break
		}
		w.workCapacity--
		results <- job
	}
	wg.Done()
}

// Starts a worker pool with predefined max number of workers.
func (w *Worker) WorkerPool(ipw int, jobs chan []string, results chan<- []string, ctx context.Context, cancel context.CancelFunc) {
	var wg sync.WaitGroup
	for i := 0; i < maxWorkers; i++ {
		var worker Worker
		worker.workCapacity = ipw
		worker.id = i
		wg.Add(1)
		go func(ctx context.Context) {
			select {
			case <-ctx.Done():
				fmt.Println("Terminate")
				return
			default:
				worker.Work(&wg, jobs, results, cancel)
			}
		}(ctx)
	}
}

// CreateJobs creates jobs based on list and max items
func (w *Worker) CreateJobs(jobs chan []string, list [][]string, maxItems int) {
	for i := 0; i < maxItems; i++ {
		jobs <- list[i]
	}
	close(jobs)
}

// Checks if total workers will be able to complete te amount of jobs based on items per worker
func CanBeDone(ipw int, totalJobs int) bool {
	totalWork := ipw * maxWorkers
	return totalWork >= totalJobs
}

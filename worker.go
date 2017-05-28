package gfp

import (
	"log"
)

type jobRequest struct {
	User *UserNode
}

type worker struct {
	ID          int
	Job         chan jobRequest
	WorkerQueue chan chan jobRequest
}

// newWorker returns a new worker.
func newWorker(id int, workerQueue chan chan jobRequest) *worker {
	return &worker{
		ID:          id,
		Job:         make(chan jobRequest),
		WorkerQueue: workerQueue,
	}
}

// start initializes worker w by adding each worker to the worker queue.
func (w *worker) start() {
	go func() {
		for {
			w.WorkerQueue <- w.Job

			select {
			case job := <-w.Job:
				log.Printf("worker %d: %s", w.ID, job.User.Login)
				job.User.run()
			case <-done:
				return
			}
		}
	}()
}

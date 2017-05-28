package gfp

// startDispatcher initializes n workers and starts the dispatch process
// for incoming work requests.
func startDispatcher(n int) {
	workerQueue := make(chan chan jobRequest, n)

	for i := 0; i < n; i++ {
		worker := newWorker(i+1, workerQueue)
		worker.start()
	}

	go func() {
		for {
			select {
			case job := <-jobQueue:
				go func() {
					worker := <-workerQueue
					worker <- job
				}()
			}
		}
	}()
}

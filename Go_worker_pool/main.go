package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID int
}

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {

	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d \n", id, job.ID)
	}
}

func main() {
	jobs := make(chan Job, 10)

	var wg sync.WaitGroup

	numWorker := 3

	for w := 1; w <= numWorker; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- Job{ID: j}
	}

	close(jobs)
	wg.Wait()
	fmt.Printf("All jobs done")
}

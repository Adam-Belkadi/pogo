package scanner

import (
	"sync"
)

// Task definition
type Task struct {
	ID int
}

// Way to process the task
func (t Task) Process(evaluatedTarget string) {
	scan_port(evaluatedTarget, t.ID)
}

// Worker pool definition
type WorkerPool struct {
	Tasks []Task
	concurrency int
	tasksChan chan Task
	wg sync.WaitGroup
}

// Functions to execute the pool
func (wp *WorkerPool) worker(evaluated_targer string) {
	for task := range wp.tasksChan {
		task.Process(evaluated_targer)
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run(evaluated_target string) {
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	// Start workers
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker(evaluated_target)
	}

	// Send tasks to the task channel
	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}
	close(wp.tasksChan)

	// Wait for all tasks to finish
	wp.wg.Wait()	

}
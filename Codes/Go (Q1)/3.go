package main

import (
	"fmt"
	"sync"
)

// Define the worker function
func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Called once when the worker finishes

	for task := range tasks {
		result := task * task
		fmt.Printf("Worker %d processed task %d (result: %d)\n", id, task, result)
	}
}

func main() {
	var wg sync.WaitGroup
	tasks := make(chan int, 6) // Buffered channel with 6 tasks

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, tasks, &wg)
	}

	// Send 6 tasks (numbers to square)
	for i := 1; i <= 6; i++ {
		tasks <- i
	}
	close(tasks) // Close the channel so workers stop when all tasks are read

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All tasks completed.")
}

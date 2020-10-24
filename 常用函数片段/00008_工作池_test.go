package command

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func worker1(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func Test_woork_pool(t *testing.T) {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker1(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
ForEnd:
	for {
		select {
		case value := <-results:
			fmt.Println(value)
		case <-time.Tick(time.Second * 3):
			fmt.Println("end")
			break ForEnd
		}
	}
}

func worker2(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func Test_woork_poo2(t *testing.T) {

	jobs := make(chan int, 100)
	results := make(chan int, 100)
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	wg.Add(5)
	wg.Wait()
	close(jobs)
	close(results)

	for value := range results {
		fmt.Println(value)
	}
}

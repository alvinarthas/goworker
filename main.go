package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// worker than make squares
func sqrWorker(tasks <-chan string, results chan<- string, instance int) {

	for num := range tasks {
		fmt.Printf("[worker %v] Sending result by worker %v\n", instance, instance)
		results <- num
	}

}

func chunkConditions(logs []int, perData int) ([]string, int) {

	var (
		condition string
		workers   = 0
		devided   []string
	)

	for i := 1; i <= len(logs); i++ {

		condition += fmt.Sprintf(`Here Number: %d, `, logs[i-1])

		if i%perData == 0 {
			condition = strings.TrimSuffix(condition, ", ")
			devided = append(devided, condition)
			workers++
			condition = ""
		} else if i == len(logs) {
			condition = strings.TrimSuffix(condition, ", ")
			devided = append(devided, condition)
			workers++
		}

	}

	return devided, workers

}

func createRandomArrays(n int) []int {

	//Provide seed
	rand.Seed(time.Now().Unix())

	return rand.Perm(n)

}

func main() {

	var (
		perData = 1000
		start   = time.Now()
	)

	logs := createRandomArrays(250000)

	chunkSize := (len(logs) + perData - 1) / perData
	fmt.Println("chunk: ", chunkSize)
	chunkArray, workers := chunkConditions(logs, perData)
	fmt.Println("chunkArray: ", len(chunkArray))
	panic("here")
	tasks := make(chan string)   // sender
	results := make(chan string) // receiver

	// launching workers goroutines
	// for i := 0; i < workers; i++ {
	// 	go sqrWorker(tasks, results, i)
	// }

	// // passing tasks
	// for _, v := range chunkArray {
	// 	tasks <- v // read da
	// }

	fmt.Println("[main] Wrote tasks")

	// closing tasks
	close(tasks)

	// receving results from all workers
	for i := 0; i < workers; i++ {
		result := <-results // non-blocking because buffer is non-empty

		fmt.Printf(`---------- RESULT %d -----------`, i)
		fmt.Println()
		fmt.Printf("%s\n", result)
		fmt.Println("---------------------")
	}

	fmt.Println("[main] main() stopped")
	fmt.Println(time.Since(start))
}

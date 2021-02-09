package channelworker

import "fmt"

func squares(c chan int) {

	for i := 0; i <= 9; i++ {
		c <- i * 1
	}

	close(c)

}

func main() {

	fmt.Println("Program Started")

	c := make(chan int)

	go squares(c) // Start GoRoutine

	// Periodic block/unblock of main goroutine until channel closes

	for {

		val, ok := <-c

		if ok == false {

			fmt.Println(val, ok, "<-- loop broke")
			break

		} else {

			fmt.Println(val, ok)

		}

	}

	fmt.Println("Program Stopped")

}

package channelworker

import "fmt"

func squares2(c chan int) {

	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}

}

func main2() {

	fmt.Println("Program Started")

	c := make(chan int, 3)

	go squares(c) // Start GoRoutine

	c <- 1
	c <- 2
	c <- 3

	fmt.Println("Program Stopped")

}

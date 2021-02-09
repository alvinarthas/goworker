package basicgoroutine

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func channeling() {
	fmt.Println("Program Started")

	c := make(chan string)

	go greet(c)

	c <- "John"

	fmt.Println("Program Stopped")
}

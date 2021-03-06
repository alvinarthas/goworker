package channelworker

import "fmt"

func sender(c chan int) {

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	close(c)
}

func main4() {

	c := make(chan int, 3)

	go sender(c)

	fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))

	// read values from c (blocked here)
	for val := range c {
		fmt.Printf("Length of channel c after value '%v' read is %v\n", val, len(c))
	}

}

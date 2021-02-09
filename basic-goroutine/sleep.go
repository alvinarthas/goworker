package basicgoroutine

import (
	"fmt"
	"time"
)

func sleep() {
	go count("sheep")
	count("fish")
}

func count(thing string) {

	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

}

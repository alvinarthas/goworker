package basicgoroutine

import (
	"fmt"
	"sync"
	"time"
)

func waitgrouptest() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		counting("sheep")
		wg.Done()
	}()

	wg.Wait()
}

func counting(thing string) {

	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

}

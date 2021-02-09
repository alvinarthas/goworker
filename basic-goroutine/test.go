package basicgoroutine

import (
	"fmt"
	"strings"
)

var (
	logs      = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
	numCPU    = 3
	condition string
	divided   []string
)

func main3() {

	chunkSize := (len(logs) + numCPU - 1) / numCPU
	fmt.Println(chunkSize)

	for i := 1; i <= len(logs); i++ {

		condition += fmt.Sprintf(`Here Number: %d, `, logs[i-1])

		if i%chunkSize == 0 {
			condition = strings.TrimSuffix(condition, ", ")
			divided = append(divided, condition)
			fmt.Println(condition)
			condition = ""
		} else if i == len(logs) {
			condition = strings.TrimSuffix(condition, ", ")
			divided = append(divided, condition)
			fmt.Println(condition)
			fmt.Println("End")
		}

	}
	fmt.Println("-------------------")
	fmt.Printf("%#v\n", divided[2])
}

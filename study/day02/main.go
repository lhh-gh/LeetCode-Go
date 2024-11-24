package main

import (
	"fmt"
	"time"
)

func asyncPrint() {

	fmt.Println("go")
	time.Sleep(111)
}
func main() {

	for i := 0; i < 100; i++ {

		go func(i int) {

			fmt.Println(i)
		}(i)
	}
	time.Sleep(10 * time.Second)

}

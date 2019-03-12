package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Outside a goroutine")
	go func() {
		fmt.Println("Inside a goroutime")
	} ()
	fmt.Println("Outside again")

	time.Sleep(1 * time.Millisecond)
	//runtime.Gosched()
}

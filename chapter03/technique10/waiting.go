package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {
	for i :=1; i <= 5; i++ {
		waitGroup.Add(1)
		go waitForAWhile(i)
	}
	waitGroup.Wait()
	fmt.Println("All done!")
}

func waitForAWhile(seconds int) {
	fmt.Printf("About to sleep for %v seconds\n", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("Finished sleeping for %v seconds\n", seconds)
	waitGroup.Done()
}

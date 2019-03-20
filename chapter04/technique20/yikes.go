package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Trapped panic [%v] (%T) \n", r, r)
		}
	}()

	yikes()
	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func yikes() {
	time.Sleep(2 * time.Second)
	panic(errors.New("Something bad happened"))
}

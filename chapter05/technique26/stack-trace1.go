package main

import (
	"fmt"
	"runtime"
)

func main() {
	foo()
}

func foo() {
	bar()
}

func bar() {
	//debug.PrintStack()
	buffer := make([]byte, 4096)
	runtime.Stack(buffer, false)
	fmt.Printf("Trace:\n\n%s\n", string(buffer))
}

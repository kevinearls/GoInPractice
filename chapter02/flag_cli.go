package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish")
	flag.BoolVar(&spanish, "s", false, "Use Spanish")
}

func main() {
	flag.Parse()

	if spanish {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}

	flag.Usage()
}

package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./log.txt")
	defer logfile.Close()

	logger := log.New(logfile, "BLAH ", log.LstdFlags | log.Lshortfile)

	logger.Println("This is a regular message")
	logger.Fatalln("Puked")
}

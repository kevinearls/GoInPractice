package main

import (
	"log"
	"net"
)

func main() {
	connection, error := net.Dial("tcp", "localhost:1902")
	if error != nil {
		panic("Failed to connect to localhost:1902")
	}

	defer connection.Close()

	loggerFlags := log.Ldate | log.Lshortfile
	logger := log.New(connection, "FUD ", loggerFlags)

	logger.Println("This is a log message")
	logger.Fatalln("Oh noes")
}

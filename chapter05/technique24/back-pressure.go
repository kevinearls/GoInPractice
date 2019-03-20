package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	timeout := 30 * time.Second
	connection, error := net.DialTimeout("udp", "localhost:1902", timeout)
	if error != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer connection.Close()

	loggerFlags := log.Ldate | log.Lshortfile
	logger := log.New(connection, "UDP ", loggerFlags)

	logger.Println(">>>> Here is a message ")

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
	//logger.Panicln("Dead")
}

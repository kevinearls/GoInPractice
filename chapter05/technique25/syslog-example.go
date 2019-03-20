package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
	flags := log.Ldate | log.Lshortfile
	logger, error := syslog.NewLogger(priority, flags)
	if error != nil {
		fmt.Printf("Can't attach to syslog: %s\n", error)
		return
	}

	logger.Println("Here is a test log message")
}

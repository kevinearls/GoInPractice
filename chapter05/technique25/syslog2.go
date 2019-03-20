package main

import "log/syslog"

func main() {
	logger, error := syslog.New(syslog.LOG_LOCAL3, "FRED")
	if error != nil {
		panic("Can't attach to syslog")
	}
	defer logger.Close()

	logger.Debug("This is a debug message")
	logger.Notice("This is a notice message")
	logger.Warning("THis is a warning")
	logger.Alert("Red Alert")
}

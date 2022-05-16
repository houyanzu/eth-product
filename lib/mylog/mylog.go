package mylog

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	myLog   *log.Logger
	logFile *os.File
)

func Init(logFileName string) error {
	var err error
	logFile, err = os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	myLog = log.New(logFile, "[debug]", log.Llongfile)
	return nil
}

func Write(v ...interface{}) {
	now := time.Now()
	_ = myLog.Output(2, fmt.Sprintln(now, v))
}

func Close() {
	_ = logFile.Close()
}

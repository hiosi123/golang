package logging

import (
	"log"
	"os"
)

func Logging() {
	// log.Println("sup ninjas!")
	// log.SetFlags(log.Ldate | log.Lshortfile)
	// log.Fatal("un-oh")

	// file, _ := os.Create("file.log")
	// log.SetOutput(file)
	// log.Println("hello world")
	// file.Close()

	// log.SetOutput(os.Stdout)
	// log.Println("Printing into standard out again")

	// Common loggers for production

	flags := log.LstdFlags | log.Lshortfile
	infoLogger := log.New(os.Stdout, "INFO: ", flags)
	warnLogger := log.New(os.Stdout, "WARN: ", flags)
	errorLogger := log.New(os.Stdout, "ERROR: ", flags)

	// infoLogger.Println("This is an info log")
	// warnLogger.Println("This is a warning log")
	// errorLogger.Println("This is an error log")

	//you can also aggregate all three into one
	al := aggregatedLogger{
		infoLogger:  infoLogger,
		warnLogger:  warnLogger,
		errorLogger: errorLogger,
	}

	al.info("this is an info log")
	al.warn("this is an warn log")
	al.error("this is an error log")
}

type aggregatedLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func (l *aggregatedLogger) info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

func (l *aggregatedLogger) warn(v ...interface{}) {
	l.warnLogger.Println(v...)
}

func (l *aggregatedLogger) error(v ...interface{}) {
	l.errorLogger.Println(v...)
}

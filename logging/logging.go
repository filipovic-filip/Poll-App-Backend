package logging

import (
	"log"
	"os"
)

var (
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate)
	errLogger = log.New(os.Stdout, "ERROR: ", log.Ltime|log.Ldate)
)

func Info(msg string) {
	infoLogger.Println(msg)
}

func Err(err error) {
	errLogger.Println(err)
}

func ErrMsg(msg string, err error) {
	errLogger.Println(msg + ", Error: " + err.Error())
}

func Panic(err error) {
	errLogger.Println("Panic! " + err.Error())
}

func PanicMsg(msg string, err error) {
	errLogger.Println("Panic!" + msg + ", Error: " + err.Error())
}

func Fatal(err error) {
	errLogger.Fatal("FATAL! " + err.Error())
}

func FatalMsg(msg string, err error) {
	errLogger.Fatal("FATAL! " + msg + ", Error: " + err.Error())
}
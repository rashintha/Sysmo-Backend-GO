package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"

var errorLabel = "[ERROR] "
var warnLabel = "[WARN] "

func init() {
	newPath := filepath.Join(".", "log")
	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		log.Println(colorRed, err.Error(), colorReset)
	}
}

func Errorln(v string) {
	log.Println(colorRed, v, colorReset)
	logFileWriteln(errorLabel + v)
}

func Errorf(format string, v ...interface{}) {
	log.Printf(colorRed, v, colorReset)
	logFileWritef(format, v)
}

func ErrorFatal(v string) {
	logFileWriteln(errorLabel + v)
	log.Fatal(colorRed, v, colorReset)
}

func ErrorFatalf(format string, v ...interface{}) {
	logFileWritef(format, v)
	log.Fatalf(colorRed, v, colorReset)
}

func Defaultln(v string) {
	log.Println(colorGreen, v, colorReset)
	logFileWriteln(v)
}

func Defaultf(format string, v ...interface{}) {
	log.Println(colorGreen, v, colorReset)
	logFileWritef(format, v)
}

func Warningln(v string) {
	log.Println(colorYellow, v, colorReset)
	logFileWriteln(warnLabel + v)
}

func Warningf(format string, v ...interface{}) {
	log.Println(colorYellow, v, colorReset)
	logFileWritef(format, v)
}

func logFileWriteln(v string) {
	f, err := os.OpenFile("log/log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(colorRed, err.Error(), colorReset)
	}

	writeString := time.Now().Local().String() + fmt.Sprintln("", v)

	if _, wErr := f.Write([]byte(writeString)); wErr != nil {
		log.Println(colorRed, wErr.Error(), colorReset)
	}

	if cErr := f.Close(); cErr != nil {
		log.Println(colorRed, cErr.Error(), colorReset)
	}
}

func logFileWritef(format string, v ...interface{}) {
	f, err := os.OpenFile("log/log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(colorRed, err.Error(), colorReset)
	}

	writeString := time.Now().Local().String() + fmt.Sprintf(format, v, "")

	if _, wErr := f.Write([]byte(writeString)); wErr != nil {
		log.Println(colorRed, wErr.Error(), colorReset)
	}

	if cErr := f.Close(); cErr != nil {
		log.Println(colorRed, cErr.Error(), colorReset)
	}
}

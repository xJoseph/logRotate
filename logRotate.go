package logRotate

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type LoggerStruct struct {
	Level   string `json:"level"`
	Date    string `json:"date"`
	Message string `json:"message"`
}

var Logger = LoggerStruct{"", "", ""}

func (Logger *LoggerStruct) Info(msg string) {
	create(msg, "info")
}

func (Logger *LoggerStruct) Warn(msg string) {
	create(msg, "warn")
}

func (Logger *LoggerStruct) Error(msg string) {
	create(msg, "error")
}

func (Logger *LoggerStruct) Bug(msg string) {
	create(msg, "bug")
}

func newLog(msg string, level string) string {
	timeNow := time.Now().UTC().String()
	Logger.Message = msg
	Logger.Level = level
	Logger.Date = timeNow[:strings.IndexByte(timeNow, '.')]
	loggerbytes, err := json.Marshal(Logger)

	if err != nil {
		panic(err)
	}
	loggerString := string(loggerbytes) + "\n"

	return loggerString
}

func create(msg string, level string) {
	timeNow := time.Now().Format(time.RFC3339)
	fileName := os.Getenv("APPDATA") + os.Getenv("APP_LOGS_PATH") + os.Getenv("APP_NAME") + "\\" + timeNow[:strings.IndexByte(timeNow, 'T')] + ".log"
	newLog := newLog(msg, level)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		os.MkdirAll(os.Getenv("APPDATA")+os.Getenv("APP_LOGS_PATH")+os.Getenv("APP_NAME"), 0755)
	}
	fileOpened, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if _, err := fileOpened.WriteString(newLog); err != nil {
		fmt.Println("Erro ao escrever logs da aplicação no arquivo de logs ", fileName)
	}
}

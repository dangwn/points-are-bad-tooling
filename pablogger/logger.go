package pablogger

import (
	"log"
	"os"
)

type PABLogger struct {
	DebugLogger *log.Logger
	InfoLogger *log.Logger
	WarningLogger *log.Logger
	ErrorLogger *log.Logger
}

func customLoggerComponent(level string) *log.Logger {
	return log.New(os.Stdout, " [" + level + "] ", log.Ldate | log.Ltime)
}

func NewLogger() *PABLogger {
	return &PABLogger{
		DebugLogger: customLoggerComponent("DEBUG"),
		InfoLogger: customLoggerComponent("INFO"),
		WarningLogger: customLoggerComponent("WARNING"),
		ErrorLogger: customLoggerComponent("ERROR"),
	}
}

func (p *PABLogger) Debug(msg any) {
	p.DebugLogger.Println(msg)
}

func (p *PABLogger) Error(msg any) {
	p.ErrorLogger.Println(msg)
}

func (p *PABLogger) Info(msg any) {
	p.InfoLogger.Println(msg)
}

func (p *PABLogger) Warning(msg any) {
	p.WarningLogger.Println(msg)
}

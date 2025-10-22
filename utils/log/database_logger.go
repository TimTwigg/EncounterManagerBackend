package logger

import (
	"context"
	"log"
	"time"
	"os"

	tracelog "github.com/jackc/pgx/v5/tracelog"
)

var DATABASE_LOGGER_COUNT int = 0

func GetLogFilePath() string {
	timeStamp := time.Now().Format("2006-01-02")
	filePath := "logs/" + timeStamp + ".log"
	
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	f.Close()

	return filePath
}

type DatabaseLogger struct {
	Logger *log.Logger
}

func (cl *DatabaseLogger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
	DATABASE_LOGGER_COUNT++
	cl.Logger.Printf("[%s] %s - %+v\n", level, msg, data)
}

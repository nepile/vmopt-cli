package logger

import (
	"fmt"
	"os"
	"time"
)

func LogEvent(message string) {
	f, err := os.OpenFile("vmopt.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Log Error:", err)
	}
	defer f.Close()

	logLine := fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
	f.WriteString(logLine)
}

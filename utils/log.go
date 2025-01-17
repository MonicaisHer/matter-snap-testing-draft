package utils

import (
	"os"
	"strings"
	"testing"
	"time"
)

func logFileName(t *testing.T, label string) string {
	fileName := label + ".log"
	if t != nil {
		fileName = strings.ReplaceAll(t.Name(), "/", "-") + "-" + fileName
	}
	logDirectory := "logs"

	err := os.MkdirAll(logDirectory, 0777)
	if err != nil {
		t.Fatalf("Can't create log directory")
	}

	return logDirectory + "/" + fileName
}

func WriteLogFile(t *testing.T, label string, content string) error {
	return os.WriteFile(
		logFileName(t, label),
		[]byte(content),
		0644,
	)
}

func WaitForLogMessage(t *testing.T, snap, expectedLog string, since time.Time) {
	const maxRetry = 10

	for i := 1; i <= maxRetry; i++ {
		time.Sleep(1 * time.Second)
		t.Logf("Retry %d/%d: Waiting for expected content in logs: %s", i, maxRetry, expectedLog)

		logs := SnapLogs(t, since, snap)
		if strings.Contains(logs, expectedLog) {
			t.Logf("Found expected content in logs: %s", expectedLog)
			return
		}
	}

	t.Fatalf("Time out: reached max %d retries.", maxRetry)
}

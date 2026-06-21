package logging

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	logFile *os.File
	mu      sync.Mutex
)

// now は「現在時刻を RFC3339 + ミリ秒」で返す
func now() string {
	t := time.Now()
	return t.Format("2006-01-02T15:04:05.000Z07:00")
}

// InitLogger は log ディレクトリを作り、ミリ秒付きのログファイルを開く
func InitLogger() error {
	// log ディレクトリ作成
	if err := os.MkdirAll("log", 0755); err != nil {
		return err
	}

	// ログファイル名作成
	t := time.Now()
	filename := fmt.Sprintf(
		"log/%04d%02d%02d%02d%02d%02d%03d.log",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/1e6, // ミリ秒
	)

	// ファイルを開く
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	logFile = f
	return nil
}

// CloseLogger はログファイルを閉じる
func CloseLogger() {
	if logFile != nil {
		logFile.Close()
		logFile = nil
	}
}

// writeLogはログの共通書き込み処理
func writeLog(level string, format string, a ...interface{}) {
	if logFile == nil {
		return
	}
	mu.Lock()
	defer mu.Unlock()

	ts := now()
	msg := fmt.Sprintf(format, a...)

	fmt.Fprintf(logFile, "%s %s %s\n", ts, level, msg)
	logFile.Sync() // 即時ディスク書き込み
}

// ログレベル別の関数
func Info(format string, a ...interface{})  { writeLog("INFO", format, a...) }
func Debug(format string, a ...interface{}) { writeLog("DEBUG", format, a...) }
func Warn(format string, a ...interface{})  { writeLog("WARN", format, a...) }
func Error(format string, a ...interface{}) { writeLog("ERROR", format, a...) }

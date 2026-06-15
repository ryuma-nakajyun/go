package logging

import (
    "fmt"
    "os"
    "time"
)

var logFile *os.File

// now は RFC3339 + ミリ秒のタイムスタンプを返す
func now() string {
    t := time.Now()
    return t.Format("2006-01-02T15:04:05.000Z07:00")
}

// InitLogger は log ディレクトリを作成し、ミリ秒付きのファイル名でログを開く
func InitLogger() error {
    // log ディレクトリ作成
    if err := os.MkdirAll("log", 0755); err != nil {
        return err
    }

    // ファイル名 yyyymmddhhmmsszzz.log
    t := time.Now()
    filename := fmt.Sprintf("log/%04d%02d%02d%02d%02d%02d%03d.log",
        t.Year(), 
        t.Month(),
        t.Day(),
        t.Hour(), 
        t.Minute(),
        t.Second(),
        t.Nanosecond()/1e6,
    )

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

// writeLog は共通ログ書き込み処理
func writeLog(level string, format string, a ...interface{}) {
    if logFile == nil {
        return
    }

    ts := now()
    msg := fmt.Sprintf(format, a...)
    fmt.Fprintf(logFile, "%s %s %s\n", ts, level, msg)
    logFile.Sync()
}

func Info(format string, a ...interface{}) {
    writeLog("INFO", format, a...)
}

func Debug(format string, a ...interface{}) {
    writeLog("DEBUG", format, a...)
}

func Warn(format string, a ...interface{}) {
    writeLog("WARN", format, a...)
}

func Error(format string, a ...interface{}) {
    writeLog("ERROR", format, a...)
}

package common

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

var (
	DebugLogger *log.Logger
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrLogger   *log.Logger
	LogLevel    = os.Getenv("SRVEXEC_LOG_LEVEL")
)

func init() {
	prefix := func(name string) string {
		return "[" + name + "] "
	}

	flags := log.Ldate | log.Ltime | log.Lshortfile

	DebugLogger = log.New(os.Stdout, prefix("DEBUG"), flags)
	InfoLogger = log.New(os.Stdout, prefix("INFO"), flags)
	WarnLogger = log.New(os.Stdout, prefix("WARN"), flags)
	ErrLogger = log.New(os.Stderr, prefix("ERROR"), flags)

	InfoLogger.Println("Logger initialized")
	if LogLevel == "" {
		LogLevel = "info"
	}
	InfoLogger.Println("Log level:", LogLevel)
}

func LogDebug(msg string) {
	if LogLevel == "debug" {
		DebugLogger.Println(msg)
	}
}

func LogInfo(msg string) {
	if LogLevel == "debug" || LogLevel == "info" {
		InfoLogger.Println(msg)
	}
}

func LogWarn(msg string) {
	if LogLevel == "debug" || LogLevel == "info" || LogLevel == "warn" {
		WarnLogger.Println(msg)
	}
}

func LogError(msg string) {
	if LogLevel == "debug" || LogLevel == "info" || LogLevel == "warn" || LogLevel == "error" {
		ErrLogger.Println(msg)
	}
}

func LogFatal(msg string) {
	ErrLogger.Fatalln(msg)
}

// Entoure un string multiligne au format markdown
// Pas de retour à la ligne à la fin.
func WrapMultiline(multiline string, lang string) string {
	if LogLevel == "debug" {
		return fmt.Sprint("\n```" + lang + "\n" + multiline + "\n```")
	} else {
		return base64.StdEncoding.EncodeToString([]byte("\n```" + lang + "\n" + multiline + "\n```"))
	}
}

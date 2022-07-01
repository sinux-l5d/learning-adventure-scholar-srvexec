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

	if LogLevel == "" {
		LogLevel = "info"
	}
	InfoLogger.Printf("initialized with level " + LogLevel)
}

func LogDebug(a ...any) {
	if LogLevel == "debug" {
		DebugLogger.Printf(a[0].(string)+"\n", a[1:]...)
	}
}

func LogInfo(a ...any) {
	if LogLevel == "debug" || LogLevel == "info" {
		InfoLogger.Printf(a[0].(string)+"\n", a[1:]...)
	}
}

func LogWarn(a ...any) {
	if LogLevel == "debug" || LogLevel == "info" || LogLevel == "warn" {
		WarnLogger.Printf(a[0].(string)+"\n", a[1:]...)
	}
}

func LogError(a ...any) {
	if LogLevel == "debug" || LogLevel == "info" || LogLevel == "warn" || LogLevel == "error" {
		ErrLogger.Printf(a[0].(string)+"\n", a[1:]...)
	}
}

func LogFatal(a ...any) {
	ErrLogger.Fatalf(a[0].(string)+"\n", a[1:]...)
}

// Entoure un string multiligne au format markdown
// Si debug : renvoie sous forme de base64, sans "\n"
// sinon : renvoie sous forme de string, avec "\n" au début
func WrapMultiline(multiline string, lang string) string {
	if LogLevel == "debug" {
		return fmt.Sprint("\n```" + lang + "\n" + multiline + "\n```")
	} else {
		return base64.StdEncoding.EncodeToString([]byte("```" + lang + "\n" + multiline + "\n```"))
	}
}

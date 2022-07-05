package common

// Ce fichier est nommé avec un prefix "1_" pour être chargé et initialisé avant les autres fichiers, qui en ont besoin

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
	LogLevel    = "info" // Default value, real value loaded in env.go
)

func init() {
	prefix := func(name string) string {
		return "[" + name + "] "
	}

	flags := log.Ldate | log.Ltime

	DebugLogger /**/ = log.New(os.Stdout, prefix(Colorize("DEBUG", Gray)), flags)
	InfoLogger /* */ = log.New(os.Stdout, prefix(Colorize(" INFO", Blue)), flags)
	WarnLogger /* */ = log.New(os.Stdout, prefix(Colorize(" WARN", Yellow)), flags)
	ErrLogger /*  */ = log.New(os.Stderr, prefix(Colorize("ERROR", Red)), flags)

	InfoLogger.Printf("loggers initialized with level " + LogLevel)
}

func LogDebug(format string, a ...any) {
	if LogLevel == "debug" {
		DebugLogger.Printf(format+"\n", a...)
	}
}

func LogInfo(format string, a ...any) {
	if LogLevel == "debug" || LogLevel == "info" {
		InfoLogger.Printf(format+"\n", a...)
	}
}

func LogWarn(format string, a ...any) {
	if LogLevel == "debug" || LogLevel == "info" || LogLevel == "warn" {
		WarnLogger.Printf(format+"\n", a...)
	}
}

func LogError(format string, a ...any) {
	if LogLevel == "debug" || LogLevel == "info" || LogLevel == "warn" || LogLevel == "error" {
		ErrLogger.Printf(format+"\n", a...)
	}
}

func LogFatal(format string, a ...any) {
	ErrLogger.Fatalf(format+"\n", a...)
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

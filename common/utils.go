package common

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"
)

// Hash un string avec sha1 et le temps actuelle
func Hash(s string) string {
	timeBytes := []byte(fmt.Sprintf("%d", time.Now().Unix()))
	toHash := append([]byte(s), timeBytes...)

	return fmt.Sprintf("%x", sha1.Sum(toHash))
}

// Retourne un context avec un timeout
func GetTimeoutCtx(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}

// Retourne la durée du timeout défini dans l'environnement,
// ou la durée par défaut (5 secondes)
func GetTimeoutEnv() time.Duration {
	tstr := Config.Get("TIMEOUT")

	// convert to time.Duration
	t, err := time.ParseDuration(tstr)

	// Default timeout
	if err != nil {
		t = time.Second * 5
		LogWarn("Timeout not set, using default value '" + t.String() + "'")
	}

	return t
}

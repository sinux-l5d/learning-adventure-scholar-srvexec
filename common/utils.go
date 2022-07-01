package common

import (
	"context"
	"crypto/sha1"
	"fmt"
	"os"
	"time"
)

func Hash(s string) string {
	timeBytes := []byte(fmt.Sprintf("%d", time.Now().Unix()))
	toHash := append([]byte(s), timeBytes...)

	return fmt.Sprintf("%x", sha1.Sum(toHash))
}

func GetTimeoutCtx(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}

func GetTimeoutEnv() time.Duration {
	tstr := os.Getenv("SRVEXEC_TIMEOUT")

	// convert to time.Duration
	t, err := time.ParseDuration(tstr)

	// Default timeout
	if err != nil {
		t = time.Second * 5
		LogWarn("Timeout not set, using default value '" + t.String() + "'")
	}

	return t
}

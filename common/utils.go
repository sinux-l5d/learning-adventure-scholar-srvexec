package common

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func Hash(s string) string {
	timeBytes := []byte(fmt.Sprintf("%d", time.Now().Unix()))
	toHash := append([]byte(s), timeBytes...)

	return fmt.Sprintf("%x", sha1.Sum(toHash))
}

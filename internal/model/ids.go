package model

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"
)

func StableID(parts ...string) string {
	h := sha1.New()
	for _, p := range parts {
		h.Write([]byte(p))
		h.Write([]byte{0})
	}
	return hex.EncodeToString(h.Sum(nil))[:16]
}
func TimeID(prefix string, at time.Time) string { return fmt.Sprintf("%s-%d", prefix, at.UnixNano()) }

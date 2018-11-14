package common

import (
	"crypto/md5"
	"fmt"
)

// GetMD5FromString returns md5 sum from string
func GetMD5FromString(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

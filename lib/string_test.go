package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrToUint8(t *testing.T) {
	r := GetMD5FromString("test string")
	assert.Equal(t, "6f8db599de986fab7a21625b7916589c", r)
}

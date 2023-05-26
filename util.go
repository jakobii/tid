package tri

import (
	"crypto/rand"
	"time"
)

func random(size int) []byte {
	b := make([]byte, size)
	rand.Read(b)
	return b
}

func unix() uint32 {
	return uint32(time.Now().Unix())
}

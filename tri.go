// Package tri creates tiny time based ids with specified length of randomness.
//
//	tri = b64(bin(time()) + rand(N))
package tri

import (
	"fmt"

	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
)

// UID return an id with a similar amount of collision safety as a uuid.
func UID() string {
	return New(12).String() //  bin(tri.seconds) + rand(12) = 16
}

// Time returns and id based only on time.
func Time() string {
	return New(0).String()
}

// Tri is a unique id based on unix epoch time + random value.
type Tri struct {
	seconds uint32
	rand    []byte
}

// new generates a new tri with a given amount of randomness.
func New(rand int) Tri {
	t := Tri{
		seconds: unix(),
	}
	if rand > 0 {
		t.rand = random(rand)
	}
	return t
}

func (t *Tri) decodeBytes(b []byte) error {
	if len(b) < 4 {
		return fmt.Errorf("invalid length: expected a slice of length 4")
	}
	t.seconds = binary.LittleEndian.Uint32(b[:4])
	if len(b) > 4 {
		t.rand = b[4:]
	}
	return nil
}

func (t Tri) encodeBytes() []byte {
	b := make([]byte, 4, 4+len(t.rand))
	binary.LittleEndian.PutUint32(b, t.seconds)
	return append(b, t.rand...)
}

// UB64 returns a url base64 encoded string of Tri.
func (t Tri) UB64() string {
	return base64.RawURLEncoding.EncodeToString(t.encodeBytes())
}

// Hex returns a hex encoded string of Tri.
func (t Tri) Hex() string {
	return hex.EncodeToString(t.encodeBytes())
}

// String returns the base64 encoded version of tri.
func (t Tri) String() string {
	return t.UB64()
}

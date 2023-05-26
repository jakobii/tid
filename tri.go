// Package tri creates tiny time based ids with specified length of randomness.
//
//	tri = b64(bin(time()) + rand(N))
package tri

import (
	"fmt"

	"encoding/base64"
	"encoding/binary"
)

func New(rand int) string {
	return new(rand).String()
}

// NewUUID return a TRI with a similar amount of collision safety as a uuid.
func NewUUID() string {
	return new(12).String() //  bin(tri.seconds) + rand(12) = 16
}

// Time returns and id based only on time.
func Time() string {
	return new(0).String()
}

// tri is a unique id based on unix epoch time + random value.
type tri struct {
	seconds uint32
	rand    []byte
}

// parse a tri from a b64 encoded string.
func parse(v string) (*tri, error) {
	b, err := base64.RawURLEncoding.DecodeString(v)
	if err != nil {
		return nil, fmt.Errorf("invalid format: expected base64 RawURLEncoding: %w", err)
	}
	return decodeBytes(b)
}

// new generates a new tri with a given amount of randomness.
func new(rand int) *tri {
	t := &tri{
		seconds: unix(),
	}
	if rand > 0 {
		t.rand = random(rand)
	}
	return t
}

func decodeBytes(b []byte) (*tri, error) {
	if len(b) < 4 {
		return nil, fmt.Errorf("invalid length: expected a slice of length 4")
	}
	var tri tri
	tri.seconds = binary.LittleEndian.Uint32(b[:4])
	if len(b) > 4 {
		tri.rand = b[4:]
	}
	return &tri, nil
}

func (t *tri) encodeBytes() []byte {
	b := make([]byte, 4, 4+len(t.rand))

	binary.LittleEndian.PutUint32(b, t.seconds)

	b = append(b, t.rand...)

	return b
}

func (t *tri) b64() string {
	return base64.RawURLEncoding.EncodeToString(t.encodeBytes())
}

// String returns the base64 encoded version of tri.
func (t *tri) String() string {
	return t.b64()
}

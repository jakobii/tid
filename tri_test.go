package tri

import (
	"fmt"
	"testing"
)

func expect[T comparable](want T, got T) error {
	if want != got {
		return fmt.Errorf("want: %v, got: %v", want, got)
	}
	return nil
}

func Test_tri_String(t *testing.T) {
	id := &tri{
		seconds: 1685084038, // unix epoch
		rand:    []byte("💩"),
	}

	err := expect("hldwZPCfkqk", id.String())
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewString_zero(t *testing.T) {
	id := new(0)
	err := expect(0, len(id.rand))
	if err != nil {
		t.Fatal(err)
	}
}

func Test_tri_encode_decode_bytes(t *testing.T) {
	id := new(2)

	b := id.encodeBytes()

	id2, err := decodeBytes(b)
	if err != nil {
		t.Fatal(err)
	}

	err = expect(id.String(), id2.String())
	if err != nil {
		t.Fatal(err)
	}
}

func Benchmark_tri_encodeBytes2(b *testing.B) {
	id := new(12)
	for n := 0; n < b.N; n++ {
		id.encodeBytes()
	}
}

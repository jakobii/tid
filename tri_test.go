package tri

import (
	"fmt"
	"testing"
)

func assertEq[T comparable](want T, got T) error {
	if want != got {
		return fmt.Errorf("want: %v, got: %v", want, got)
	}
	return nil
}

func Test_tri_String(t *testing.T) {
	id := Tri{
		seconds: 1685084038, // unix epoch
		rand:    nil,
	}

	err := assertEq("hldwZA", id.String())
	if err != nil {
		t.Fatal(err)
	}

	id.rand = []byte("💩")

	err = assertEq("hldwZPCfkqk", id.String())
	if err != nil {
		t.Fatal(err)
	}

}

func Test_new_zero(t *testing.T) {
	id := New(0)
	err := assertEq(0, len(id.rand))
	if err != nil {
		t.Fatal(err)
	}
}

func Test_tri_encode_decode_bytes(t *testing.T) {
	id := New(2)

	b := id.encodeBytes()

	var id2 Tri
	if err := id2.decodeBytes(b); err != nil {
		t.Fatal(err)
	}

	if err := assertEq(id.String(), id2.String()); err != nil {
		t.Fatal(err)
	}
}

func Benchmark_tri_encodeBytes2(b *testing.B) {
	id := New(12)
	for n := 0; n < b.N; n++ {
		id.encodeBytes()
	}
}

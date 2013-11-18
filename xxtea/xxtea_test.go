package xxtea

import (
	"math/rand"
	"testing"
)

func isEqual(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestAll(t *testing.T) {
	Nkeys := 8
	Blen := 256

	keys := make([][4]uint32, Nkeys)
	in := make([]uint32, Blen)
	out := make([]uint32, Blen)

	for _, key := range keys {
		for i := range key {
			key[i] = rand.Uint32()
		}
	}
	for i := range in {
		in[i] = rand.Uint32()
	}

	for _, key := range keys {
		copy(out, in)
		Encrypt(out, key)
		if isEqual(in, out) {
			t.Fatal("identity encription!")
		}
		refDecrypt(out, key)
		if !isEqual(in, out) {
			t.Fatal("encryption fail")
		}
		refEncrypt(out, key)
		if isEqual(in, out) {
			t.Fatal("identity encription!")
		}
		Decrypt(out, key)
		if !isEqual(in, out) {
			t.Fatal("decription fail")
		}
	}
}

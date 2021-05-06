package cuid

import (
	"crypto/rand"
	"reflect"
	"testing"
)

func Test_CUIDCryptoType(t *testing.T) {
	c, e := NewCrypto(rand.Reader)
	if e != nil {
		t.Errorf("Failed to generate crypto cuid: %v", e)
	}
	if reflect.TypeOf(c).Name() != "string" {
		t.Error("Incorrect CUID type")
	}
}

func Test_CUIDCryptoFormat(t *testing.T) {
	c, e := NewCrypto(rand.Reader)
	if e != nil {
		t.Errorf("Failed to generate crypto cuid: %v", e)
	}
	if err := IsCuid(c); err != nil {
		t.Error("Incorrect format")
	}
}

func Test_CUIDCryptoCollisions(t *testing.T) {
	ids := map[string]bool{}
	for i := 0; i < 600000; i++ {
		id, e := NewCrypto(rand.Reader)
		if e != nil {
			t.Errorf("Failed to generate crypto cuid, at iteration %d: %v", i, e)
		}
		if ids[id] == true {
			t.Errorf("Collision detected, at iteration %d", i)
		}
		ids[id] = true
	}
}

func Benchmark_CryptoCUIDGeneration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCrypto(rand.Reader)
	}
}

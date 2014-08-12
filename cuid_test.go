package cuid

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func Test_CUIDType(t *testing.T) {
	c := New()
	if reflect.TypeOf(c).Name() != "string" {
		t.Error("Incorrect CUID type")
	}
}

func Test_CUIDFormat(t *testing.T) {
	c := New()
	format := regexp.MustCompile(fmt.Sprintf("c[0-9a-z]{%d}", 6*BLOCK_SIZE))
	if !format.MatchString(c) {
		t.Error("Incorrect format")
	}
}

func Test_CUIDCollisions(t *testing.T) {
	ids := map[string]bool{}
	for i := 0; i < 600000; i++ {
		id := New()
		if ids[id] == true {
			t.Errorf("Collision detected, at iteration %d", i)
		}
		ids[id] = true
	}
}

func Test_SlugType(t *testing.T) {
	s := NewSlug()
	if reflect.TypeOf(s).Name() != "string" {
		t.Error("Incorrect CUID type")
	}
}

func Test_SlugFormat(t *testing.T) {

}

func Benchmark_CUIDGeneration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

// +build parallel

package mymath_synced_test

import (
	"testing"

	"go-rr.org/mymath_synced"
)

func TestCubeSynced(t *testing.T) {
	if mymath_synced.CubeSynced(4) != 64 {
		t.FailNow()
	}
}

func BenchmarkCubeSynced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mymath_synced.CubeSynced(4)
	}
}

func BenchmarkCubeSyncedParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mymath_synced.CubeSynced(4)
		}
	})
}

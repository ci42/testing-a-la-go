// +build benchmarks

package mymath_test

import (
	"testing"

	"go-rr.org/mymath"
)

func BenchmarkCube64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mymath.CubeInt64(6)
	}
}

// run with lock?
func BenchmarkCube64Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mymath.CubeInt64(6)
		}
	})
}

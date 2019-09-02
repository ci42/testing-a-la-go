// +build internal

package mymath

import (
	"testing"
)

func TestCubeInternal(t *testing.T) {
	if cubeInternal(4) != 64 {
		t.FailNow()
	}
}

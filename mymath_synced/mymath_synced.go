package mymath_synced

import (
	"sync"

	"go-rr.org/mymath"
)

var l sync.Mutex

func CubeSynced(x int) int {
	l.Lock()
	r := mymath.Cube(x)
	l.Unlock()
	return r
}

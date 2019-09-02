// +build propertytests

package mymath_test

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	"go-rr.org/mymath"
)

func int16Generator(values []reflect.Value, r *rand.Rand) {
	for i := range values {
		values[i] = reflect.ValueOf(int16(0))
		// values[i] = reflect.ValueOf(int16(r.Int31()))
	}
}

var retainSignConfig = &quick.Config{
	MaxCount:      0,
	MaxCountScale: 0,
	Rand:          nil,
	Values:        int16Generator,
}

func float32Generator(values []reflect.Value, r *rand.Rand) {
	for i := range values {
		values[i] = reflect.ValueOf(float32(0.0))
		// values[i] = reflect.ValueOf(float32(r.NormFloat64()))
	}
}

var sumOfThreeConfig = &quick.Config{
	MaxCount:      0,
	MaxCountScale: 0,
	Rand:          nil,
	Values:        float32Generator,
}

func TestCubeRetainsSign(t *testing.T) {
	property := func(x int16) bool {
		result := mymath.CubeInt16(x)
		if x < 0 {
			return result < 0
		}
		return result >= 0
	}

	if err := quick.Check(property, retainSignConfig); err != nil {
		t.Fatalf("%s", err)
	}
}

func TestAdditionOfThreeSummandsIsCommutative(t *testing.T) {
	f1 := func(x, y, z float32) float32 {
		return mymath.AdditionOffThreeSummands(x, y, z)
	}

	f2 := func(x, y, z float32) float32 {
		return mymath.AdditionOffThreeSummands(z, y, x)
	}

	if err := quick.CheckEqual(f1, f2, sumOfThreeConfig); err != nil {
		t.Fatalf("%s", err)
	}
}

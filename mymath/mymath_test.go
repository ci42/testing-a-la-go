package mymath_test

import (
	"testing"

	"go-rr.org/mymath"
)

func TestCubeInt8(t *testing.T) {
	if mymath.Cube(4) != 64 {
		t.FailNow()
	}
}

// func TestCubeSkip(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("Solving Rubiks Cube takes too long, skipping...")
// 	}
// }

// func TestCubeSubtests(t *testing.T) {
// 	t.Run("2**3", func(t *testing.T) {
// 		if mymath.Cube(2) != 8 {
// 			t.FailNow()
// 		}
// 	})
// 	t.Run("3**3", func(t *testing.T) {
// 		if mymath.Cube(3) != 27 {
// 			t.FailNow()
// 		}
// 	})
// }

// func TestCubeParallel(t *testing.T) {
// 	t.Run("4**3", func(t *testing.T) {
// 		t.Parallel()
// 		if mymath.Cube(4) != 64 {
// 			t.FailNow()
// 		}
// 	})
// 	t.Run("5**3", func(t *testing.T) {
// 		t.Parallel()
// 		if mymath.Cube(5) != 125 {
// 			t.FailNow()
// 		}
// 	})

// }

// func TestCubeTableDriven(t *testing.T) {

// 	testcases := []struct {
// 		name   string
// 		in     int64
// 		result int64
// 	}{
// 		{name: "3**3", in: 3, result: 27},
// 		{name: "4**3", in: 4, result: 64},
// 		{name: "5**3", in: 5, result: 125},
// 	}

// 	for _, testcase := range testcases {
// 		t.Run(testcase.name, func(t *testing.T) {

// 			testcase := testcase
// 			t.Parallel()

// 			if r := mymath.CubeInt64(testcase.in); r != testcase.result {
// 				t.Fatalf("Actual r: %d, Wanted: %d", r, testcase.result)
// 			}
// 		})
// 	}
// }

// func TestMain(m *testing.M) {

// 	result := 0

// 	flag.Parse()

// 	// run (expensive) setup here

// 	// acutally run the tests
// 	result = m.Run()

// 	// run teardown here

// 	os.Exit(result)
// }

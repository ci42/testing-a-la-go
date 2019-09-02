package mymath

func CubeInt64(x int64) int64 {
	return x * x * x
}

func CubeInt32(x int32) int32 {
	return x * x * x
}

func CubeInt16(x int16) int16 {
	return x * x * x
}

func CubeInt8(x int8) int8 {
	return x * x * x
}

func Cube(x int) int {
	return cubeInternal(x)
}

func cubeInternal(x int) int {
	// could memorize results
	return x * x * x
}

func AdditionOffThreeSummands(x, y, z float32) float32 {
	return x + y + z
}

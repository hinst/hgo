package hgo

import "strconv"

func Float32ToStr(x float32) string {
	return strconv.FormatFloat(float64(x), 'f', -1, 32)
}

func Float64ToIntStr(x float64) string {
	return strconv.FormatFloat(x, 'f', 0, 32)
}

func Float64ToIntStr0(x float64, length int) (result string) {
	result = strconv.FormatFloat(x, 'f', 0, 32)
	for len(result) < length {
		result = "0" + result
	}
	return
}

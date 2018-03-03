package hgo

func LockIntBetween(a, x, b int) int {
	if b <= x {
		x = b - 1
	}
	if x < a {
		x = a
	}
	return x
}

func CheckIntInRange(a, x, b int) bool {
	return a <= x && x < b
}

func BoolToInt(x bool) int {
	if x {
		return 1
	} else {
		return 0
	}
}

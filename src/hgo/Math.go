package hgo

func LockIntBetween(a, x, b int) int {
	if b <= x {
		x = b - 1
	}
	if x < a {
		x = 0
	}
	return x
}

func BoolToInt(x bool) int {
	if x {
		return 1
	} else {
		return 0
	}
}

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

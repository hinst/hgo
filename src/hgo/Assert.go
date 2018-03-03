package hgo

func AssertResult(e error) {
	if e != nil {
		panic(e)
	}
}

func Assert(condition bool) {
	if !condition {
		panic("Assert")
	}
}

func AssertWT(condition bool, message func() string) {
	if !condition {
		panic(message())
	}
}

func AssertIntEquals(a, b int) {
	if a != b {
		panic(IntToStr(a) + " != " + IntToStr(b))
	}
}

func AssertIntInRange(a, x, b int) {
	AssertWT(a <= x && x < b,
		func() string { return "Index out of range: " + IntToStr(a) + "," + IntToStr(x) + "," + IntToStr(b) })
}

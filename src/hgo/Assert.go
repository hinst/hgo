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

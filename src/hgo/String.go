package hgo

import "strconv"

func BoolToStr(x bool) string {
	if x {
		return "true"
	} else {
		return "false"
	}
}

func Int64ToStr(x int64) string {
	return strconv.FormatInt(x, 10)
}

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func StrToInt0(text string) (result int) {
	var x, parseResult = strconv.ParseInt(text, 10, 32)
	if parseResult == nil {
		result = int(x)
	}
	return
}

func StrToInt(text string) int {
	var result, parseResult = strconv.ParseInt(text, 10, 32)
	AssertResult(parseResult)
	return int(result)
}

func Add0ToStr(s string, length int) string {
	if len(s) < length {
		var prefix = make([]byte, length-len(s))
		for i := range prefix {
			prefix[i] = '0'
		}
		s = string(prefix) + s
	}
	return s
}

package hgo

import "runtime"

func GetCallerName(level int) string {
	callers := make([]uintptr, 16)
	runtime.Callers(1+level, callers)
	f := runtime.FuncForPC(callers[0])
	return f.Name()
}

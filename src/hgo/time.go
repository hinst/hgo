package hgo

import (
	"time"
)

func FormatFullTime(time time.Time) string {
	var toS = func(i int) string {
		return IntToStr(i)
	}
	var toS2 = func(i int) string {
		return Add0ToStr(IntToStr(i), 2)
	}
	var toS3 = func(i int) string {
		return Add0ToStr(IntToStr(i), 3)
	}
	return "" + toS(time.Year()) +
		"." + toS2(int(time.Month())) +
		"." + toS2(time.Day()) +
		"-" + toS2(time.Hour()) +
		":" + toS2(time.Minute()) +
		":" + toS2(time.Second()) +
		"." + toS3(time.Nanosecond()/1000/1000)
}

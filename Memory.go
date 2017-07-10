package hgo

import (
	"runtime/debug"
	"time"
)

func ReleaseSystemMemory() {
	debug.FreeOSMemory()
}

func EnablePeriodicReleaseMemory(interval time.Duration) {
	go func() {
		time.Sleep(interval)
		ReleaseSystemMemory()
	}()
}

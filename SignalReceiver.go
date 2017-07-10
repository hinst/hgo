package hgo

import (
	"os"
	"os/signal"
	"syscall"
)

func InstallShutdownReceiver(receiver func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for currentSignal := range c {
			var _ = currentSignal
			receiver()
		}
	}()
}

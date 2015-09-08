package base

import (
	"os"
	"os/signal"

	"syscall"
)

func Wait(funcs func()) {
	for {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGHUP)
		<-ch
		if funcs != nil {
			funcs()
		}
		NOTICE("exit")
		os.Exit(1)
	}
}

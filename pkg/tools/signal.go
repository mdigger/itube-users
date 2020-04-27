package tools

import (
	"os"
	"os/signal"
	"syscall"
)

// WaitSignal ждет получения сигнала от операционной системы.
// По умолчанию ждет сигналов SIGTERM или SIGINT.
func WaitSignal(signals ...os.Signal) os.Signal {
	if len(signals) == 0 {
		signals = []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	}
	var signalChan = make(chan os.Signal, 1)
	signal.Notify(signalChan, signals...)
	var sig = <-signalChan
	return sig
}

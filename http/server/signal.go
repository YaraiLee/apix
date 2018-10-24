/*
 * Revision History:
 *     Initial: 2018/5/26        ShiChao
 */

package server

import (
	"os/signal"
	"runtime"
	"syscall"
)

func (ep *Entrypoint) configureSignals() {
	if runtime.GOOS == "windows" {
		signal.Notify(ep.signals, syscall.SIGINT, syscall.SIGTERM)
	} else {
		signal.Notify(ep.signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	}
}

func (ep *Entrypoint) listenSignals() {
	for {
		sig := <-ep.signals

		switch sig {
		case syscall.SIGUSR1:
		default:
			ep.Stop()
			return
		}
	}
}

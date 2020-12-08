package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func serveSignal(stop chan struct{}) error {
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stop:
		close(quit)
	case <-quit:
		close(stop)
	}

	signal.Reset(syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("signal return")
	return nil
}

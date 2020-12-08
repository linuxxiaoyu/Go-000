package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"

	"golang.org/x/sync/errgroup"
)

func main() {
	stop := make(chan struct{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		return serveApp(stop)
	})

	g.Go(func() error {
		return serveSignal(stop)
	})

	if err := g.Wait(); err != nil {
		fmt.Println("g.Wait err:", err)
	}
}

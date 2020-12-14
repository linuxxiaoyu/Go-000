package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func realMain() {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(c)

	g.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			close(quit)
			return ctx.Err()
		case <-quit:
			cancel()
		}
		return nil
	})

	g.Go(func() error {
		s := http.Server{
			Addr:    ":8080",
			Handler: nil,
		}

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(1 * time.Second)
			fmt.Fprintf(w, "Hello World!")
		})

		http.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
			cancel()
			fmt.Fprintf(w, "Close!")
		})

		done := make(chan error)
		go func(ctx context.Context) {
			select {
			case <-ctx.Done():
				fmt.Println("shutdown")
				ctx2, cancle := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancle()

				s.Shutdown(ctx2)
				done <- ctx.Err()
			}
		}(ctx)

		err := s.ListenAndServe()
		if err != nil {
			return <-done
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("g.Wait err:", err)
	}
}

func main() {
	realMain()
}

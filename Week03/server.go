package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func serveApp(stop chan struct{}) error {
	s := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// time.Sleep(2 * time.Second)
		fmt.Fprintf(w, "Hello World!")
	})

	http.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		close(stop)
		fmt.Fprintf(w, "Close!")
	})

	done := make(chan error)
	go func() {
		<-stop
		fmt.Println("shutdown")
		ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancle()

		done <- s.Shutdown(ctx)
	}()

	err := s.ListenAndServe()
	if err != nil {
		return <-done
	}

	return nil
}

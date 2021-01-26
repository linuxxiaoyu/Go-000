package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serve(ctx)
}

func serve(ctx context.Context) error {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println(errors.Wrap(err, "accept error"))
				continue
			}
			go handleConn(ctx, conn)
		}
	}
}

func handleConn(ctx context.Context, conn net.Conn) {
	fmt.Println(conn.RemoteAddr(), "connected")

	defer conn.Close()
	ch := make(chan []byte)
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		return read(ctx, conn, ch)
	})

	g.Go(func() error {
		return write(ctx, conn, ch)
	})

	g.Wait()
	fmt.Println(conn.RemoteAddr(), "closed")
}

func read(ctx context.Context, conn net.Conn, ch chan []byte) error {
	rd := bufio.NewReader(conn)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			line, _, err := rd.ReadLine()
			if err != nil {
				close(ch)
				return err
			}
			ch <- line
		}

	}
}

func write(ctx context.Context, conn net.Conn, ch chan []byte) error {
	wr := bufio.NewWriter(conn)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			line, ok := <-ch
			if !ok {
				return nil
			}

			if len(line) <= 0 {
				continue
			}
			wr.WriteString("Hello ")
			wr.Write(line)
			wr.WriteString("\n")
			wr.Flush()
		}
	}
}

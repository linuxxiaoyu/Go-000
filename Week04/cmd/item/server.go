package main

import (
	"context"
	"net"
	"time"

	"github.com/linuxxiaoyu/Go-000/Week04/api"
	"github.com/linuxxiaoyu/Go-000/Week04/internal/service"
	"github.com/linuxxiaoyu/Go-000/Week04/pkg/app"

	"google.golang.org/grpc"
)

func NewServer(port string) *app.App {
	srv := grpc.NewServer()
	api.RegisterItemServiceServer(srv, &service.ItemService{})
	appServer := app.New(
		app.StartTimeout(3*time.Second),
		app.StopTimeout(3*time.Second),
	)
	appServer.Append(app.Hook{
		OnStart: func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				lis, err := net.Listen("tcp", port)
				if err != nil {
					return err
				}
				return srv.Serve(lis)
			}
		},
		OnStop: func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				srv.GracefulStop()
				return nil
			}
		},
	})
	return appServer
}

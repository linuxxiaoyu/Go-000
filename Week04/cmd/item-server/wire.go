//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/linuxxiaoyu/Go-000/Week04/pkg/app"
)

func InitializeApp(port string) *app.App {
	wire.Build(NewServer)
	return &app.App{}
}

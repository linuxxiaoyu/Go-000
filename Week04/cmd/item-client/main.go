package main

import (
	"context"
	"fmt"
	"log"

	"github.com/linuxxiaoyu/Go-000/Week04/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Printf("Dail: %+v\n", err)
		return
	}
	defer conn.Close()

	client := api.NewItemServiceClient(conn)
	resp, err := client.AddItem(context.Background(), &api.ItemRequest{
		// Name: "err",
		Name: "Go West",
		// Name:  "",
		Price: float32(3099),
	})
	if err != nil {
		fmt.Printf("AddItem: %+v\n", err)
		return
	}
	fmt.Println(resp.GetId())
}

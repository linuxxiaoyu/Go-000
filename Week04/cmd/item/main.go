package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	port = flag.Int("port", 8080, "service port")
)

func realMain() {
	flag.Parse()
	app := InitializeApp(fmt.Sprintf(":%d", *port))
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
	defer app.Stop()
}

func main() {
	realMain()
}

package main

import (
	"fmt"
	"go-games/get-fias-info/routers"
)

const (
	serverStarted = "server started"
	port          = "8888"
)

func main() {
	fmt.Println("Start")
	routers.StartWebServer(port)
}

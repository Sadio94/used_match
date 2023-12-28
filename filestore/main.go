package main

import (
	"filestore/service/http"
	"os"
	"os/signal"
)

func main() {
	go http.GinStart()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
}

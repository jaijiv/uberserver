package main

import (
	"os"
	"os/signal"

	srv1 "github.com/jaijiv/uberserver/srv1/handler"
	srv2 "github.com/jaijiv/uberserver/srv2/handler"
)

func main() {
	go srv1.RunMain()
	go srv2.RunMain()

	waitForExit()
}

func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	os.Exit(0)
}

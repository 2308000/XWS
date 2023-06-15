package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("Reservation service stopped")
		done <- true
		os.Exit(0)
	}()
	//config := cfg.NewConfig()
	//server := startup.NewServer(config)
	fmt.Println("Reservation service has started!")
	//server.Start()
	<-done
}

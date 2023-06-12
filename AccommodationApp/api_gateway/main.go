package main

import (
	"accommodation_booking/api_gateway/startup"
	cfg "accommodation_booking/api_gateway/startup/config"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		done <- true
		os.Exit(0)
	}()

	config := cfg.NewConfig()
	server := startup.NewServer(config)
	fmt.Printf("Gateway has been started!")
	server.Start()
	<-done
}
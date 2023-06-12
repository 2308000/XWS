package main

import (
	"accommodation_booking/profile_service/startup"
	cfg "accommodation_booking/profile_service/startup/config"
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
		fmt.Printf("Profile service stopped")
		done <- true
		os.Exit(0)
	}()
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	fmt.Printf("Profile service has started!")
	server.Start()
	<-done
}

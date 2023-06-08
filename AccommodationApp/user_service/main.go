package main

import (
	"accommodation_booking/user_service/startup"
	cfg "accommodation_booking/user_service/startup/config"
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
		done <- true
		os.Exit(0)
	}()

	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
	<-done
}

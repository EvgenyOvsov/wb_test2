package main

import (
	"draft/datastore"
	"draft/server"
	"os"
	"os/signal"
	"syscall"
)

// Дефолтная архитектура с обработкой прерываний
func main() {

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	ds := datastore.NewDatastore("datastore_service:8080")
	srv := server.NewServer(ds)
	go srv.Start()

	<-signalChan
	srv.Stop()
}

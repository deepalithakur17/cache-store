package main

import (
	"encoding/json"
	"fmt"
	filesave "gorm-gogo/filesave"
	"gorm-gogo/routes"
	"gorm-gogo/storage"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type RequestData struct {
	Data string `json:"data"`
	Time time.Time
}

func main() {
	// Set up signal channel to listen for shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	router := routes.Routes()

	if data, err := os.ReadFile("dump"); err == nil {
		if len(data) == 0 {
			log.Println("Dump is empty")
		} else if err := json.Unmarshal(data, &storage.Chache); err != nil {
			log.Printf("Failed to unmarshal data from dump: %v", err)
		} else {
			fmt.Println("Loaded previous data from dump.")
		}
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Run server in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	fmt.Println("Server running on port 8080")

	// shut down signal
	<-sigChan
	fmt.Println("Received shutdown signal. Initiating graceful shutdown...**")
	filesave.SaveDataToFile()

}

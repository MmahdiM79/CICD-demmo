package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cicd-demo/internal/demo"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "dev"
	}
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           demo.Handler(version, time.Now().UTC().Format(time.RFC3339)),
		ReadHeaderTimeout: 5 * time.Second,
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		shutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdown); err != nil {
			log.Print(err)
		}
		close(done)
	}()
	log.Printf("cicd-demo %s listening on http://localhost:%s", version, port)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
	<-done
}

package main

import (
	"log"
	"os"
	"time"

	"cicd-demo/internal/demo"
)

func main() {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "dev"
	}
	page, err := demo.RenderPage(version, time.Now().UTC().Format(time.RFC3339))
	if err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("dist", 0755); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("dist/index.html", page, 0644); err != nil {
		log.Fatal(err)
	}
	log.Print("wrote dist/index.html")
}

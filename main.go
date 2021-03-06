package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rodrigodealer/whiplash/server"
	"github.com/rodrigodealer/whiplash/tracing"
)

func main() {
	log.SetOutput(os.Stdout)
	tracing.StartTracing("localhost:8080", "whiplash")
	log.Print("Starting server on port 8080")
	err := http.ListenAndServe(":8080", tracing.TrackerHandler(server.Server()))
	if err != nil {
		log.Panic("Something is wrong : " + err.Error())
	}
}

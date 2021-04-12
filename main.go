package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", StatusHandler)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.ListenAndServe(":"+port, mux)
}

type Status struct {
	Name       string `json:"name"`
	StatusText string `json:"status_text"`
	StatusCode int    `json:"status_code"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := Status{
		Name:       "my bot",
		StatusText: "OK",
		StatusCode: 200,
	}
	response, _ := json.Marshal(status)
	w.Write(response)
}

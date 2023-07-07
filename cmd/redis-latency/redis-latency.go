package main

import (
	"log"
	"net/http"
	"redis-latency/internal/latency"
)

func main() {
	http.HandleFunc("/", latency.TestLatency)
	err := http.ListenAndServe(":8085", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

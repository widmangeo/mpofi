package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", HelloHandler)

	log.Println("Listening on port", port)
	go startTicker()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello from Koyeb\n")
}

func startTicker() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		resp, err := http.Get("http://67a5eb37c8b283181c3e.appwrite.global")
		if err != nil {
			log.Println("Error visiting URL:", err)
			continue
		}
		log.Println("Visited URL, status code:", resp.StatusCode)
		resp.Body.Close()
	}
}

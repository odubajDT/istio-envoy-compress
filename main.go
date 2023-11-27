package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ResponseData struct {
	Message string `json:"message"`
}

// KO_DOCKER_REPO=docker.io/odubajdt ko build . --bare --tags test
// kubectl label namespace istio-test istio-injection=enabled

func longPayloadHandler(w http.ResponseWriter, r *http.Request) {
	longData := "This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer. This is a really long payload that can be as long as needed for demonstration purposes. You can add more content here to make it longer."

	data := ResponseData{Message: longData}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func shortPayloadHandler(w http.ResponseWriter, r *http.Request) {
	shortData := "Short payload"
	_, err := fmt.Fprintf(w, shortData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Handling long payload on port 8080
	http.HandleFunc("/long", longPayloadHandler)
	http.HandleFunc("/short", shortPayloadHandler)

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("Short payload server error: ", err)
		}
	}()

	fmt.Println("Servers running on ports 8080 (long payload) and 8081 (short payload)")
	select {}
}

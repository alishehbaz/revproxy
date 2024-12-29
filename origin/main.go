package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	port := 8081

	originServerHandler := http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			fmt.Printf("[origin server] received request at: %s\n", time.Now())

			response := map[string]string{
				"source": "Origin Server",
				"time":   time.Now().UTC().Format(time.RFC3339),
			}

			fmt.Fprint(rw, "origin server response\n")
			json.NewEncoder(rw).Encode(response)
		})

	fmt.Printf("Origin server is up and running on port %d\n", port)
	log.Fatal(http.ListenAndServe(":8081", originServerHandler))

}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	port := 9081

	reverseProxyHandler := http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			fmt.Printf("[reverse proxy server] request received at: %s\n", time.Now())

			response := map[string]string{
				"source": "Reverse Proxy Server",
				"time":   time.Now().UTC().Format(time.RFC3339),
			}

			fmt.Fprint(rw, "reverse proxy server response\n")
			json.NewEncoder(rw).Encode(response)

		})

	fmt.Printf("Reverse proxy server is up and running on port %d\n", port)
	log.Fatal(http.ListenAndServe(":9081", reverseProxyHandler))

}

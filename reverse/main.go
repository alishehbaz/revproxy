package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {

	port := 9081

	originServerURL, err := url.Parse("http://127.0.0.1:8081")
	if err != nil {
		log.Fatal("there is something wrong with the origin server url")
	}

	reverseProxyHandler := http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			fmt.Printf("[reverse proxy server] request received at: %s\n", time.Now())

			req.Host = originServerURL.Host
			req.URL.Host = originServerURL.Host
			req.URL.Scheme = originServerURL.Scheme
			req.RequestURI = ""

			response := map[string]string{
				"source": "Reverse Proxy Server",
				"time":   time.Now().UTC().Format(time.RFC3339),
			}

			http.DefaultClient.Do(req)

			fmt.Fprint(rw, "reverse proxy server response\n")
			json.NewEncoder(rw).Encode(response)

		})

	fmt.Printf("Reverse proxy server is up and running on port %d\n", port)
	log.Fatal(http.ListenAndServe(":9081", reverseProxyHandler))

}

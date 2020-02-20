package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Save a copy of this request for debugging.
		requestDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(
			w,
			"Request: %v\n",
			string(requestDump),
		)
		log.Printf("Request: %v\n",
			string(requestDump))
	})

	httpServer := &http.Server{
		Addr: ":8080",
	}

	httpServer.ListenAndServe()
}

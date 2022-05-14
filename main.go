package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	statusCode int
)

func init() {
	c := os.Getenv("STATUS_CODE")
	if c == "" {
		statusCode = 200
	} else {
		var err error
		statusCode, err = strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	serveSimpleHTTP()
}

func serveSimpleHTTP() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(statusCode)
		default:
			http.Error(w, `{"status": "only GET allowed"}`, http.StatusMethodNotAllowed)
		}
	})
	log.Printf("Start simple api server on 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

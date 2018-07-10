package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Simple app running...")
	target := os.Getenv("SIMPLE_MSG")
	if target == "" {
		target = ":( SIMPLE_MSG variable not defined"
	}
	fmt.Fprintf(w, "Hello %s!\n", target)
}

func main() {
	flag.Parse()
	log.Print("Simple app server started...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

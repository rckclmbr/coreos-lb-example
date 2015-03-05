package main;

import (
    "net/http"
    "fmt"
    "os"
)

func hello(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    fmt.Fprintf(w, "Hello world from host: %s", hostname)
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":5000", nil)
}

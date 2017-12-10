package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    port :=8080
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world!")
    })
    http.HandleFunc("/toilet", toiletHandler)
    http.HandleFunc("/gedung", gedungHandler)
    log.Printf("Server starting on port %v\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}



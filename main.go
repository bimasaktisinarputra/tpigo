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
    http.HandleFunc("/t/", func(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		s:=r.URL.Path[len("/t/"):]
		if s!=""{
			GetT(w,r,s)
		}
	default:
		http.Error(w,"Invalid request method.", 405)
		fmt.Fprintf(w, "Failed to get DB")
	}
    })
    http.HandleFunc("/toilet", toiletHandler)
    http.HandleFunc("/gedung", gedungHandler)
    log.Printf("Server starting on port %v\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}



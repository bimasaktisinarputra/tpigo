package main

import (
    "fmt"
    "log"
    "net/http"

)

func main() {
    port :=8022
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    fmt.Fprintf(w, "Hello!\nUntuk melihat data seluruh Toilet, gunakan 167.205.67.246:8022/AllToilet/\nUntuk melihat data seluruh Gedung, gunakan 167.205.67.246:8022/AllGedung/\nUntuk melihat data toilet dengan id=n, gunakan 167.205.67.246:8022/Toilet/[n]\nUntuk melihat data gedung dengan id=n,gunakan 167.205.67.246:8022/Gedung/[n]")
    })
    /*http.HandleFunc("/t/", func(w http.ResponseWriter, r *http.Request){
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
    })*/

    //Panggil data satu gedung
    http.HandleFunc("/Gedung/", func(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case "GET":
        s:=r.URL.Path[len("/Gedung/"):]
        if s!=""{
            GetGedung(w,r,s)
        }
    default:
        http.Error(w,"Invalid request method.", 405)
        fmt.Fprintf(w, "Failed to get DB")
    }
    })
    
    //Panggil data semua gedung
    http.HandleFunc("/AllGedung/", func(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case "GET":
            GetAllGedung(w,r)
    default:
        http.Error(w,"Invalid request method.", 405)
        fmt.Fprintf(w, "Failed to get DB")
    }
    })

    //Panggil data satu toilet
    http.HandleFunc("/Toilet/", func(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case "GET":
        s:=r.URL.Path[len("/Toilet/"):]
        if s!=""{
            GetToilet(w,r,s)
        }
    default:
        http.Error(w,"Invalid request method.", 405)
        fmt.Fprintf(w, "Failed to get DB")
    }
    })

    //Panggil data semua gedung
    http.HandleFunc("/AllToilet/", func(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case "GET":
            GetAllToilet(w,r)
    default:
        http.Error(w,"Invalid request method.", 405)
        fmt.Fprintf(w, "Failed to get DB")
    }
    })

    log.Printf("Server starting on port %v\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}



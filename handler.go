package main

import(
	"net/http"
	"time"
	"encoding/json"
	"database/sql"
	"strconv"
	"log"

)

func gedungHandler(w http.ResponseWriter, r *http.Request) {
		mygedung:= MyGedung {
		Gedung{id: 1, Nama: "Labtek 5", Alias: "Labtek V"},
		Gedung{id: 2, Nama: "Labtek 8", Alias: "Labtek 8"},
	}
	json.NewEncoder(w).Encode(mygedung)
}
func toiletHandler(w http.ResponseWriter, r *http.Request) {
		mytoilet:=MyToilet {
		Toilet{id: 1, Buka: time.Date(2017,time.December,0,7,0,0,0, time.Local), Tutup: time.Date(2017,time.December,0,18,0,0,0,time.Local), Hari: "Weekdays", Nilai: 8, Tempat: "1", Lantai: 1},
		Toilet{id: 2, Buka: time.Date(2017,time.December,0,7,0,0,0,time.Local), Tutup: time.Date(2017,time.December,0,18,0,0,0,time.Local), Hari: "Semua", Nilai: 9, Tempat: "1", Lantai: 2},
	}
	json.NewEncoder(w).Encode(mytoilet)
}

func GetT(w http.ResponseWriter, r *http.Request,id string){
	myid, _ := strconv.Atoi(id)
	db, err := sql.Open("mysql","tamu:kosong@tcp(127.0.0.1:3306)/t")
	if err!= nil{log.Fatal(err)}
	defer db.Close()
	myt := MyT{}

	rows, err:= db.Query("select ID, Tempat from myt where id=?", myid)
	if err!=nil{log.Fatal(err)}
	defer rows.Close()
	for rows.Next(){
		err:= rows.Scan(&myt[myid].ID, &myt[myid].Tempat)
		if err!=nil{log.Fatal(err)}
		json.NewEncoder(w).Encode(&myt)
	}
	err=rows.Err()
	if err!=nil{log.Fatal(err)}
}

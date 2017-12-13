package main

import(
	"net/http"
	_ "time"
	"encoding/json"
	"database/sql"
	"strconv"
	"log"
	_ "github.com/go-sql-driver/mysql"
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
		//Toilet{id: 1, Buka: time.Date(2017,time.December,0,7,0,0,0, time.Local), Tutup: time.Date(2017,time.December,0,18,0,0,0,time.Local),  Nilai: 8, Tempat: "1", Lantai: 1},
		//Toilet{id: 2, Buka: time.Date(2017,time.December,0,7,0,0,0,time.Local), Tutup: time.Date(2017,time.December,0,18,0,0,0,time.Local),  Nilai: 9, Tempat: "1", Lantai: 2},
	}
	json.NewEncoder(w).Encode(mytoilet)
}

func GetT(w http.ResponseWriter, r *http.Request,id string){
	myid, _ := strconv.Atoi(id)
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/t")
	if err!= nil{log.Fatal(err)}
	defer db.Close()
	//myt := MyT{}
	tc:=Tcoba{}

	rows, err:= db.Query("select ID, Tempat from t where id=?", myid)
	if err!=nil{log.Fatal(err)}
	defer rows.Close()
	for rows.Next(){
		err:= rows.Scan(&tc.ID, &tc.Tempat)
		if err!=nil{log.Fatal(err)}
		json.NewEncoder(w).Encode(&tc)
	}

	err=rows.Err()
	if err!=nil{log.Fatal(err)}
}

func GetGedung(w http.ResponseWriter, r *http.Request, id string){
	myid, _ := strconv.Atoi(id)
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/tpi")
	if err!= nil{log.Fatal(err)}
	defer db.Close()
	//myt := MyT{}
	gd:=Gedung{}

	rows, err:= db.Query("select id, nama, nama_alt from gedung where id=?", myid)
	if err!=nil{log.Fatal(err)}
	defer rows.Close()
	for rows.Next(){
		err:= rows.Scan(&gd.id, &gd.Nama, &gd.Alias)
		if err!=nil{log.Fatal(err)}
		json.NewEncoder(w).Encode(&gd)
	}

	err=rows.Err()
	if err!=nil{log.Fatal(err)}
}

func GetToilet(w http.ResponseWriter, r *http.Request, id string){
	myid, _ := strconv.Atoi(id)
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/tpi")
	if err!= nil{log.Fatal(err)}
	defer db.Close()
	//myt := MyT{}
	to:=Toilet{}

	rows, err:= db.Query("select id, buka, tutup, nilai, tempat, lantai from toilet where id=?", myid)
	if err!=nil{log.Fatal(err)}
	defer rows.Close()
	for rows.Next(){
		err:= rows.Scan(&to.id, &to.Buka, &to.Tutup, &to.Nilai, &to.Tempat, &to.Lantai)
		if err!=nil{log.Fatal(err)}
		json.NewEncoder(w).Encode(&to)
	}

	err=rows.Err()
	if err!=nil{log.Fatal(err)}
}

func Test(w http.ResponseWriter, r *http.Request, id string){
	myid, _ := strconv.Atoi(id)
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/tpi")
	if err!= nil{log.Fatal(err)}
	defer db.Close()
	//myt := MyT{}
	//to:=Toilet{}

	rows, err:= db.Query("select buka from toilet where id=?", myid)
	if err!=nil{log.Fatal(err)}
	defer rows.Close()
	
		//err= rows.Scan(&to.Buka)
		if err!=nil{log.Fatal(err)}
		//log.Print(to.Buka)
	

	err=rows.Err()
	if err!=nil{log.Fatal(err)}
}

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
/*
func GetT(w http.ResponseWriter, r *http.Request,id string){
	myid, _ := strconv.Atoi(id)
	db, err := sql.Open("mysql","root:@tcp(167.205.67.251:3306)/t")
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
}*/

func GetGedung(w http.ResponseWriter, r *http.Request, id string){
	myid, _ := strconv.Atoi(id)
	db, err := sql.Open("mysql","root:@tcp(167.205.67.251:3306)/tpi")
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

func GetAllGedung(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql",
		"root:@tcp(167.205.67.251:3306)/tpi")
	if err !=nil {
		log.Fatal(err)
	}
	defer db.Close()

	gd := Gedung{}
	rows, err := db.Query("select id, nama, nama_alt from gedung")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&gd.id, &gd.Nama, &gd.Alias)
		if err !=nil{
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(&gd)
	}
	err=rows.Err()
}


func GetToilet(w http.ResponseWriter, r *http.Request, id string){
	myid, _ := strconv.Atoi(id)
	db, err := sql.Open("mysql","root:@tcp(167.205.67.251:3306)/tpi")
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

func GetAllToilet(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql",
		"root:@tcp(167.205.67.251:3306)/tpi")
	if err !=nil {
		log.Fatal(err)
	}
	defer db.Close()

	to := Toilet{}
	rows, err := db.Query("select id, buka, tutup, nilai, tempat, lantai from toilet")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&to.id, &to.Buka, &to.Tutup, &to.Nilai, &to.Tempat, &to.Lantai)
		if err !=nil{
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(&to)
	}
	err=rows.Err()
}

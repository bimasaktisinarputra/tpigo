package main

import _ "time"

type Gedung struct {
	id	int
	Nama	string
	Alias	string
}
type MyGedung []Gedung

type Toilet struct {
	id	int
	Buka	string
	Tutup	string
	Nilai	int
	Tempat	string
	Lantai	int
}
type MyToilet []Toilet

type Tcoba struct{
	ID	int
	Tempat	string
}
type MyT []Tcoba

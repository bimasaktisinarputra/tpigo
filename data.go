package main

import "time"

type Gedung struct {
	id	int
	Nama	string
	Alias	string
}
type MyGedung []Gedung

type Toilet struct {
	id	int
	Buka	time.Time
	Tutup	time.Time
	Hari	string
	Nilai	int
	Tempat	string
	Lantai	int
}
type MyToilet []Toilet


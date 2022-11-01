package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Interface")
	//Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan nama tertentu
	//Interface merupakan tipe data
	//Nilai objek bertipe interface zero value-nya adalah nill

	//Maka kita bisa 'obrol' dengan fungsi bangun datar (lingkaran dan persegi) via interface hitung
	//bangunData bertipe interface hitung
	var bangunDatar hitung

	//implement interface untuk persegi
	bangunDatar = persegi{10.0}
	fmt.Println("==========Persegi")
	fmt.Println("Luas     : ", bangunDatar.luas())
	fmt.Println("Keliling : ", bangunDatar.keliling())

	//implement interface untuk lingkaran
	bangunDatar = lingkaran{20.0}
	fmt.Println("==========Lingkaran")
	fmt.Println("Luas     : ", bangunDatar.luas())
	fmt.Println("Keliling : ", bangunDatar.keliling())

	//variabel-nya harus di-casting terlebih dahulu ke tipe asli variabel konkritnya (pada kasus ini tipenya lingkaran)
	//Cara casting objek interface sedikit unik, yaitu dengan menuliskan nama tipe tujuan dalam kurung,
	//ditempatkan setelah nama interface dengan menggunakan notasi titik (seperti cara mengakses property hanya saja ada tanda kurung nya)
	fmt.Println("Jari-jari: ", bangunDatar.(lingkaran).jariJari())

	//Embedded Interface
	//Interface bisa di-embed ke interface lain, sama seperti struct.
	//Cara penerapannya juga sama, cukup dengan menuliskan nama interface yang ingin di-embed ke dalam interface tujuan
	var bangunRuang2 hitungku = &kubus{4}

	fmt.Println("===== kubus")
	fmt.Println("luas      :", bangunRuang2.luas())
	fmt.Println("keliling  :", bangunRuang2.keliling())
	fmt.Println("volume    :", bangunRuang2.volume())
}

// Penerapan Interface
// Keyword type dan interface digunakan untuk pendefinisian interface.
// Dengan memanfaatkan interface hitung, perhitungan luas dan keliling bangun datar bisa dilakukan,
// tanpa perlu tahu jenis bangun datarnya sendiri itu apa.
// ini bisa digunakan pada bangun datar : lingkaran, kotak, persegi panjang, dll
type hitung interface {
	luas() float64
	keliling() float64
}

//Siapkan struct nya untuk masing-masing bidang datar

// LINGKARAN
type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

// PERSEGI
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

// EMBEDDED INTERFACE
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitungku interface {
	hitung2d
	hitung3d
}

// kubus
type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

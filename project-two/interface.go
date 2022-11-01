package main

import (
	"fmt"
	"math"
	"strings"
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

	//TIPE DATA Interface Kosong (any)
	//notasinya : interface{}
	//Variabel bertipe ini akan dapat menampung segala tipe data, bahkan array,pointer,apapun. Disebut sebagai Dynamic Typing.
	var secret interface{}

	secret = "ethan hunt" //berupa string
	fmt.Println(secret)   //ethan hunt

	secret = []string{"apple", "manggo", "banana"} //berupa slice string
	fmt.Println(secret)                            //[apple manggo banana]

	//Paling sering dalam bentuk map
	var data map[string]interface{} //data bertipe map, dengan key tipe string dan value tipe interface{}
	data = map[string]interface{}{"name": "bana  budhiana", "grade": 20, "breakfast": []string{"nasi", "susu", "telor"}}
	fmt.Println(data["name"])

	//tipe data any
	//interface kosong bisa dituliskan dengan 'any'
	var data2 map[string]any //interface{} bisa diganti any

	data2 = map[string]any{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}
	fmt.Println(data2["grade"])

	//Casting Variabel Interface Kosong
	var secret2 interface{}
	secret2 = 2                     //tipe data masih interface{}
	var number = secret2.(int) * 10 //bentuk casting nya variable.(nama_type_to_casting)
	fmt.Println(secret2, "multiplied by 10 is :", number)

	secret2 = []string{"apple", "manggo", "banana"}     //ingat secret2 adalah bertipe data interface{}
	var gruits = strings.Join(secret2.([]string), ", ") //maka harus di casting dengan ([]string) agar bisa dikasih operasi strings
	fmt.Println(gruits, "is my favorite fruits")

	//Kombinasi Slice, map, dan interface{}
	//Tipe []map[string]interface{} adalah salah satu tipe yang paling sering digunakan
	var person = []map[string]interface{}{ //slice dari tipe data map yang memiliki key bertipe string dan value bertipe interface{}
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}
	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

	//contoh lain bahwa interface sangat berguna
	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}
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

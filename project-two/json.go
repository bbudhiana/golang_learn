package main

import (
	"encoding/json"
	"fmt"
)

// 1.siapkan object struct user
// Struct User ini akan digunakan untuk buat variabel baru penampung hasil decode json string
// Proses decode dilakukan dengan fungsi json.Unmarshal
type User struct {
	Fullname string `json:"Name"`
	Age      int
}

func main() {
	fmt.Println("JSON")
	//JSON : Javascript Object Notation adalah notasi standar yang umum digunakan untuk komunikasi data dalam web
	//GO menyediakan encoding/json berisikan banyak fungsi untuk kebutuhan operasi json

	//DECODE JSON ke variabel Object Struct
	//dalam Go, data json dituliskan sebagai string. Dengan json.Unmarsha, Json string itu bisa dikonversi
	//menjadi bentuk object, bisa map[string]interface{} atau object struct
	//2. misal json string nya
	var jsonString = `{"Name":"Bana Budhiana", "Age":27}`
	var jsonData = []byte(jsonString) //konversi jadi slice byte

	//siapkan data tampungan
	var data User

	//lakukan decode dengan json.Unmarshal
	//Fungsi unmarshal hanya menerima data json dalam bentuk []byte
	//argument ke-2 fungsi unmarshal harus diisi dengan pointer dari objek yang nantinya akan menampung hasilnya
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("User :", data.Fullname)
	fmt.Println("Age :", data.Age)

	//Decode JSON Ke map[string]interface{} & interface{}
	//Tak hanya ke objek cetakan struct, target decoding data json juga bisa berupa variabel bertipe map[string]interface{}
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])

	//Variabel bertipe interface{} juga bisa digunakan untuk menampung hasil decode.
	//Dengan catatan pada pengaksesan nilai property, harus dilakukan casting terlebih dahulu ke map[string]interface{}.
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{}) //convert dulu data2 dari interface{} ke map[string]interface{}
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])

	//Decode Array JSON Ke Array Objek
	var jsonString2 = `[
		{"Name": "john wick", "Age": 27},
		{"Name": "ethan hunt", "Age": 32}
	]`
	//siapkan data tampungan bertipe slice User
	var data3 []User

	//lakukan unmarshal
	var err2 = json.Unmarshal([]byte(jsonString2), &data3)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	fmt.Println("user 1:", data3[0].Fullname)
	fmt.Println("user 2:", data3[1].Fullname)

	//Encode Objek Ke JSON String
	//encode data objek ke bentuk json string
	//Fungsi json.Marshal digunakan untuk encoding data ke json string.
	//Sumber data bisa berupa variabel objek cetakan struct, map[string]interface{}, atau slice
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData3, err3 = json.Marshal(object) //lakukan Marshal ke object untuk konversi ke string, return []byte, error
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}

	var jsonString3 = string(jsonData3) //Casting []byte ke string
	fmt.Println(jsonString3)
}

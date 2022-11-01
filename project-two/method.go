package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Method")
	//Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lain).
	//Method bisa diakses lewat variabel object
	//Keunggulan method dibanding fungsi adalah memiliki akses ke property struct hingga level private.

	//TEST PENGGUNAAN
	var m1 = murid{"Bana Budhiana", 20}
	m1.sayHello()

	var name = m1.getNameAt(1)
	fmt.Println("nama panggilan :", name)

	//TEST METHOD POINTER
	m1.changeName1("John Rambo")
	fmt.Println("After changeName1 :", m1.name) // After changeName1 : Bana Budhiana (walau diganti, name tidak berubah di object struct nya)

	m1.changeName2("Jason Bourne")
	fmt.Println("After changeName2 :", m1.name) //After changeName2 : Jason Bourne (berubah karena pointer nya ke object struct nya)

	//Keistimewaan lain method pointer adalah, method itu sendiri bisa dipanggil dari objek pointer maupun objek biasa.
	// akses method dari object biasa
	var m2 = murid{"Gajah Mada", 33}
	m2.sayHello()

	//akses method dari variabel object pointer
	var m3 = &murid{"Jane Rambu", 22}
	m3.sayHello()
}

// IMPLEMENTASI
// ketika deklarasi ditentukan juga pemilik method itu.
type murid struct {
	name  string
	grade int
}

// deklarasi method ditentukan dulu siapa yg punya method ini
// Struct yang digunakan akan menjadi pemilik method
// func (s student) sayHello() maksudnya adalah fungsi sayHello dideklarasikan sebagai method milik struct student
func (m murid) sayHello() {
	fmt.Println("Hello My Man !", m.name)
}

// ambil name base on position
func (m murid) getNameAt(i int) string {
	//strings.Split Fungsi ini berguna untuk memisah string menggunakan pemisah yang ditentukan sendiri.
	//Hasilnya adalah array berisikan kumpulan substring
	/*
		strings.Split("ethan hunt", " ") // ["ethan", "hunt"]
	*/
	return strings.Split(m.name, " ")[i-1]
}

// METHOD POINTER
// method pointer adalah method yang variabel object pemilik method berupa pointer nya
func (m murid) changeName1(name string) {
	fmt.Println("-----> on changeName1, name changed to :", name)
	m.name = name
	fmt.Println("-----> hanya berubah di dalam method changeName1 saja :", m.name)
}

// method pointer nya
func (m *murid) changeName2(name string) {
	fmt.Println("-----> on changeName2, name changed to :", name)
	m.name = name
}

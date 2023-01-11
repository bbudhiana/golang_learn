package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	fmt.Println("Tentang Struct")

	//GO tidak punya OOP, tapi punya struct
	//STRUCT adalah kumpulan definisi variabel (disebut properti) dan atau fungsi (atau method), yang
	//semuanya dibungkus sebagai sebuah TIPE DATA BARU dengan nama tertentu.
	//Dari struct yang dibuat, kita bisa buat variabel baru dengan atribut sesuai dengan skema struct
	//variabel dari struct bisa kita sebut dengan OBJECT or OBJECT STRUCT

	//Keyword 'type' digunakan untuk deklarasi struct
	/*
			type student struct {
		        	name string
				grade int
			}
	*/
	var s1 student
	s1.name = "Bana"
	s1.grade = 10

	//ubah properti nya
	s1.grade = 11
	fmt.Printf("Namanya : %v dengan grade: %v\n", s1.name, s1.grade)

	//inisialisasi dengan memberikan tambahan {}
	var s2 = student{}
	s2.name = "John Wich"
	s2.grade = 12
	fmt.Printf("Namanya : %v dengan grade: %v\n", s2.name, s2.grade)

	//dengan cara horizontal
	var s3 = student{"Rambo", 9}
	fmt.Printf("Namanya : %v dengan grade: %v\n", s3.name, s3.grade)

	//dengan menggunakan nama properti nya
	var s4 = student{name: "Jenny", grade: 13}
	fmt.Printf("Namanya : %v dengan grade: %v\n", s4.name, s4.grade)

	//variabel struct bisa diambil pointernya
	var s5 *student
	s5 = &s4
	s5.name = "Nama Berubah"
	fmt.Printf("Alamat %v adalah %v\n", s4, s5)

	//EMBEDDED STRUCT
	//Struct dalam struct, struct sebagai property struct lain
	var s6 = pegawai{}      //inisialisasi
	s6.person.name = "Jaja" //melalui person
	s6.age = 20             //bisa langsung
	s6.id = 9324
	fmt.Printf("Data s6 : %v %v %v\n", s6.name, s6.age, s6.id)

	//secara langsung bisa juga di input variabelnya
	var p1 = person{name: "abdul", age: 12}
	var s7 = pegawai{id: 3423432, person: p1}
	fmt.Printf("Data s7 : %v\n", s7)

	//Anonymous Struct, struct sekali pakai, tanpa pengisian property (pake {})
	var s8 = struct {
		id int
		person
	}{}
	s8.person = person{name: "Noah", age: 90}
	s8.id = 324324
	/*s8.person.name = "Noah"
	s8.person.age = 90*/
	fmt.Printf("Data s8 : %v\n", s8)

	//Anonymous Struct, struct sekali pakai, tanpa pengisian property (pake {})
	var s9 = struct {
		id int
		person
	}{
		person: person{name: "Noah", age: 70},
		id:     1212,
	}
	/*s8.person.name = "Noah"
	s8.person.age = 90*/
	fmt.Printf("Data s9 : %v\n", s9)

	//SLICE and STRUCT
	var allStudents = []person{
		{name: "bana", age: 20},
		{name: "joni", age: 32},
		{name: "jeni", age: 22},
		{name: "rambo", age: 30},
	}
	for _, student := range allStudents {
		fmt.Printf("nama %v - age %v\n", student.name, student.age)
	}

	//SLICE ANONYMOUS STRUCT
	var semuaMurid = []struct {
		person
		grade int
	}{
		{person: person{name: "bana", age: 20}, grade: 1},
		{person: person{name: "jeni", age: 33}, grade: 2},
		{person: person{name: "jony", age: 21}, grade: 1},
	}
	for _, murid := range semuaMurid {
		fmt.Printf("nama %v dan age %v - grade %v\n", murid.person.name, murid.person.age, murid.grade)
	}
}

type person struct {
	name string
	age  int
}

type pegawai struct {
	id int
	person
}

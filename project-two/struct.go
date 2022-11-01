package main

import "fmt"

func main() {
	fmt.Println("Lebih jauh tentang Go")
	fmt.Println("=====================")
	//GO tidak mendukung OOP namun memiliki tipe data terstruktur disebut dengan Struct
	//Struct : kumpulan definisi variabel(property) dan fungsi (method), dibungkus sebagai tipe data baru dengan nama tertentu
	//dari struct itu kita bisa buat variabel baru yang punya atribut sesuai skema struct.

	//DEKLARASI
	//gunakan 'type' untuk deklarasi struct
	type student struct {
		name  string
		grade int
	}

	//IMPLEMENTASI
	var s1 student
	s1.name = "Bana Budhiana"
	s1.grade = 10

	fmt.Println("Name of Student\t:", s1.name)
	fmt.Println("Grade\t\t:", s1.grade)

	//INISIALISASI
	//Cara 1
	var s2 = student{} //inisialisasi dulu ke object struct nya
	s2.name = "Bana Budhiana 2"
	s2.grade = 10

	//Cara 2
	var s3 = student{"Bana Budhiana 3", 8} //inisialisasi langsung

	fmt.Println("student 2 :", s2.name) //student 2 : Bana Budhiana 2
	fmt.Println("student 3 :", s3)      //student 3 : {Bana Budhiana 3 8}

	//Cara 3
	var s4 = student{name: "John"} //grade tidak diisi, defaultnya 0
	fmt.Println("student 4", s4)   //student 4 {John 0}

	//VARIABEL OBJECT POINTER
	//object yang dibuat dari tipe struct dapat diambil pointer nya dan disimpan di variabel object bertipe struct pointer
	var s5 = student{name: "john", grade: 3}

	//adalah variabel pointer hasil cetakan struct student  *student
	var s6 = &s5                              //s5 adalah variabel pointer dengan nilainya adalah alamat dari s5 (&s5) atau s6 menampung nilai referensi s5
	fmt.Println("student 5, name :", s5.name) //student 5, name : john
	fmt.Println("student 6, name :", s6.name) //student 6, name : john
	s6.name = "Bana"                          //coba diubah di referensi nya
	fmt.Println("student 5, name :", s5.name) //student 5, name : Bana //variabel yg direferensi ikut berubah
	fmt.Println("student 6, name :", s6.name) //student 6, name : Bana

	//EMBEDDED STRUCT
	//embeded struct adalah mekanisme menempelkan sebuah struct sebagai properti struct lain
	//Embedded struct adalah mutable, nilai property-nya nya bisa diubah.
	type person struct {
		name string
		age  int
	}

	type murid struct {
		grade int
		age   int
		person
	}
	//implementasi
	var s7 = murid{}
	s7.name = "John Rambo" //s7.person.name bisa juga
	s7.age = 40            //age ini punya struct murid
	s7.person.age = 41     //age ini punya struct person
	s7.grade = 10

	fmt.Println("Data : ", s7.name, s7.age, "grade:", s7.grade, "age person", s7.person.age)
	//Embedded Struct Dengan Nama Property Yang Sama
	//pengaksesan property-nya harus dilakukan secara eksplisit atau jelas

	//Pengisian Nilai Sub-Struct
	var p1 = person{name: "wick", age: 21}
	var s8 = murid{person: p1, grade: 2, age: 22} //person di isi dengan variabel struct p1

	fmt.Println("name  :", s8.name)
	fmt.Println("age   :", s8.age)
	fmt.Println("grade :", s8.grade)

	//ANONYMOUS STRUCT
	//Anonymous struct adalah struct yang tidak dideklarasikan di awal sebagai tipe data baru, melainkan langsung ketika pembuatan objek
	//Dilakukan karena hanya sekali pakai
	//Salah satu aturan yang perlu diingat dalam pembuatan anonymous struct adalah, deklarasi harus diikuti dengan inisialisasi
	var pak_rt = struct {
		person
		jabatan string
	}{} //deklarasi harus diikuti inisialisasi yaitu dengan memberikan {}, jadi {} adalah wajib buat anonymous

	pak_rt.person = person{"John Rambo", 42}
	pak_rt.jabatan = "Ketua RT"
	fmt.Println("Pak RT :", pak_rt.person.name, pak_rt.jabatan) //Pak RT : John Rambo Ketua RT

	//KOMBINASI SLICE dan STRUCT
	//Slice dan struct bisa dikombinasikan seperti pada slice dan map
	//implementasi
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}
	//maka bisa di iterasi kan
	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	//Inisialisasi Slice Anonymous Struct
	//Anonymous struct bisa dijadikan sebagai tipe sebuah slice
	var allStudents2 = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}
	for _, student := range allStudents2 {
		fmt.Println(student)
	}

	//Deklarasi Anonymous Struct Menggunakan Keyword var
	//Cara lain untuk deklarasi anonymous struct adalah dengan menggunakan keyword var
	//akan sangat berbeda 'var' dengan 'type', dengan var akan membuat sebuah anonymous struct, kemudian disimpan dalam var student2
	var student2 struct {
		person
		grade int
	}

	student2.person = person{"wick", 21}
	student2.grade = 2

	//Deklarasi anonymous struct menggunakan metode ini juga bisa dilakukan beserta inisialisasi-nya.
	// hanya deklarasi
	var student3 struct {
		grade int
	}

	// deklarasi sekaligus inisialisasi
	var student4 = struct {
		grade int
	}{
		12,
	}
	fmt.Println(student3, student4) //{0} {12}

	//NESTED STRUCT atau STRUCT BERSARANG
	//Nested struct adalah anonymous struct yang di-embed ke sebuah struct
	//Deklarasinya langsung di dalam struct peng-embed
	//Teknik ini biasa digunakan ketika decoding data json yang struktur datanya cukup kompleks dengan proses decode hanya sekali
	type student5 struct {
		person struct { //anonymous struct
			name string
			age  int
		}
		grade   int
		hobbies []string
	}

	//Deklarasi Dan Inisialisasi Struct Secara Horizontal
	//Tanda semi-colon (;) digunakan sebagai pembatas deklarasi poperty
	//type person2 struct { name string; age int; hobbies []string }
	//inisialisasi nya :
	//var p3 = struct { name string; age int } { age: 22, name: "wick" }
	//var p4 = struct { name string; age int } { "ethan", 23 }
	//fmt.Println(p3, p4)

	//Tag property dalam struct
	//Tag merupakan informasi opsional yang bisa ditambahkan pada masing-masing property struct.
	//Tag biasa dimanfaatkan untuk keperluan encode/decode data json
	//Informasi tag juga bisa diakses lewat reflect
	type person3 struct {
		name string `tag1`
		age  int    `tag2`
	}

	//Type Alias
	//Sebuah tipe data, seperti struct, bisa dibuatkan alias baru, caranya dengan type NamaAlias = TargetStruct
	type Person4 struct {
		name string
		age  int
	}
	type People = Person4 //People adalah alias dari struct Person4

	var p3 = Person4{"wick", 21}
	fmt.Println(p3)

	var p4 = People{"wick", 21}
	fmt.Println(p4)
}

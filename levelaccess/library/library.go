package library

import "fmt"

// level aksesnya adalah PUBLIC, ditandai dengan nama fungsi diawali huruf besar
func SayHello() {
	fmt.Println("hello manusia, apa kabar?")
	//introduce("Johny")
}

// level akses PRIVATE, ditandai oleh huruf kecil di awal nama fungsi
func introduce(name string) {
	fmt.Println("nama saya", name)
}

// INGAT : agar dapat di exported maka harus kapital huruf pertamanya
type Murid struct {
	Name  string //agar bisa diakses maka harus kapital
	Grade int
}

// dengan cara anonymous struct
var Murid2 = struct {
	Name  string
	Grade int
}{}

// dipanggil lebih dulu sebelum fungsi main() dieksekusi.
func init() {
	Murid2.Name = "John Wick"
	Murid2.Grade = 2

	fmt.Println("--> library/library.go imported")
}

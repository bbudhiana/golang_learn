package main

import "fmt"

func main() {
	fmt.Println("Property")
	//Properti Public dan Private
	//Dalam Go tidak ada istilah public modifier dan private modifier. Yang adalah
	//exported yg ekivalen dengan public modifier dan unexported ekivalen dengan private modifier
	//1.Exported Package dan Unexported Package
	//- Di Go, setiap folder atau sub-folder adalah satu package, file-file yang ada di dalam sebuah folder package-nya harus sama
	//- package pada file-file tersebut harus berbeda dengan package pada file-file lainnya yang berada pada folder berbeda
	//- Dalam sebuah package, biasanya kita menulis sangat banyak komponen, entah itu fungsi, struct, variabel, atau lainnya
	//2 Jenis Hak akses di Go
	//1. Exported atau public. Menandakan komponen tersebut diperbolehkan untuk diakses dari package lain yang berbeda
	//2. Unexported atau private. Berarti komponen hanya bisa diakses dalam package yang sama, bisa dalam satu file saja atau dalam beberapa file yang masih 1 folder

	//menentukan level akses atau modifier sangat mudah, penandanya adalah character case huruf pertama nama fungsi, struct, variabel, atau lainnya
	// - Ketika namanya diawali dengan huruf kapital menandakan kalau exported (atau public)
	// - Jika diawali huruf kecil, berarti unexported (atau private)
	// Lihat di levelaccess

	//Import Dengan Prefix Tanda Titik
	//Di Go, komponen yang berada di package lain yang di-import bisa dijadikan se-level dengan komponen package peng-import,
	//caranya dengan menambahkan tanda titik (.) setelah penulisan keyword import

	//Pemanfaatan Alias Ketika Import Package
	//misal :
	//import ( f "fmt")
	//maka panggilnya cukup : f.Println("test")

	//Fungsi init()
	//Fungsi ini otomatis dipanggil pertama kali ketika aplikasi di-run.
	//Jika fungsi ini berada dalam package main, maka dipanggil lebih dulu sebelum fungsi main() dieksekusi.
}

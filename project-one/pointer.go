package main

import "fmt"

func main() {
	fmt.Println("POINTER")

	//variabel pointer ditandai dengan * sebelum tipe data
	var number_pointer *int
	var kata_pointer *string

	//variabel biasa dapat diambil pointer valuenya dengan prefix & (disebut referencing)
	var number = 100
	number_pointer = &number
	var kata = "Jaja Miharja"
	kata_pointer = &kata

	//variabel pointer juga bisa diambil nilai nya dengan prefix * (disebut dereferencing)
	var nilai_number = *number_pointer
	var nilai_kata = *kata_pointer

	fmt.Printf("Variabel %v punya alamat memory/pointer: %v\n", number, number_pointer)
	fmt.Printf("Variabel %v punya alamat memory/pointer: %v\n", kata, kata_pointer)
	fmt.Printf("Nilai-nilai pointer %v dan %v\n", nilai_number, nilai_kata)

	var nilai_utama = 13
	changeValue(&nilai_utama, 100)
	fmt.Printf("Nilainya sekarang adalah : %v\n", nilai_utama) //100

}

func changeValue(original *int, new_value int) {
	//ambil value nya dan ubah
	*original = new_value
}

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//untuk recover panic error
	defer catch()

	fmt.Println("Error")
	//Error merupakan sebuah tipe
	//Error memiliki 1 buah properti berupa method error(), yg mengembalikan detail error berupa string

	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input) //capture input user dan masukkan ke dalam variabel input

	var number int
	var err error
	number, err = strconv.Atoi(input) //strconv.Atoi mengembalikan 2 nilai, hasil konversi dan error (bisa nill or ada)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error()) //jika ada maka print pesan errornya (error())
	}

	//IMPLEMENTASI CUSTOM ERROR
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error()) //panic akan menampilkan stact trace error
		fmt.Println(err.Error())
	}
}

// CUSTOM ERROR
func validate(input string) (bool, error) {
	//strings.TrimSpace() digunakan untuk menghilangkan karakter spasi sebelum dan sesudah string
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty") //errors.New adalah mekanisme utama custom error
	}
	return true, nil
}

// RECOVER : handle panic error
// jadi setelah ketemu panic error, recover akan kembalikan ke kondisi awal ketika panic menghentikan alur program
// gunakan defer untuk jalankan recover di akhir eksekusi program, karena defer di eksekusi saat akhir statement
func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

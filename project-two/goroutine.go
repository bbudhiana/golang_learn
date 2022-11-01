package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	fmt.Println("Go Routine")

	//menentukan jumlah core yang digunakan untuk eksekusi
	runtime.GOMAXPROCS(2)

	//jalankan dengan goroutine
	go print(5, "HAllooo")

	//jalankan tanpa goroutine
	print(5, "apa kabarrrr")

	//Fungsi fmt.Scanln() mengakibatkan proses jalannya aplikasi berhenti di baris ini (blocking) hingga user menekan tombol enter
	var input string
	fmt.Scanln(input)
}

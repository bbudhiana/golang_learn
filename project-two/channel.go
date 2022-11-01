package main

import (
	"fmt"
	"runtime"
)

func main() {
	//Channel digunakan untuk menghubungkan goroutine satu dengan goroutine lain.
	//Pengiriman dan penerimaan data pada channel bersifat blocking atau synchronous
	fmt.Println("Channel")
	runtime.GOMAXPROCS(2)

	//pembuatan channel
	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)

		//variabel data dikirim lewat channel messages
		messages <- data
	}

	//sayHelloTo di-eksekusi secara goroutine atau asynchronous atau tidak saling tunggu
	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	//proses penerimaan data dari channel yang di kanan, untuk disimpan ke variabel yang di kiri
	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}

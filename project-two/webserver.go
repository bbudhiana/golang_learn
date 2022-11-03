package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pie mas, sae?")
}

func main() {
	fmt.Println("Menggunakan Web Server")

	//Fungsi http.HandleFunc() digunakan untuk routing aplikasi web
	//Fungsi http.HandleFunc() memiliki 2 buah parameter yang harus diisi
	//1.Parameter pertama adalah rute yang diinginkan
	//2.Parameter kedua adalah callback atau aksi ketika rute tersebut diakses
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hallo gan, apa kabar?")
	})

	http.HandleFunc("/index", index)

	fmt.Println("Bismillah, starting server at http://localhost:8081")

	//with parameter
	http.HandleFunc("/parameter", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "john wick",
			"Message": "have a nice day",
		}

		//Fungsi template.ParseFiles() digunakan untuk parsing template, mengembalikan 2 data yaitu instance template-nya dan error (jika ada).
		var t, err = template.ParseFiles("project-two/template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	//Fungsi http.listenAndServe() digunakan untuk menghidupkan server sekaligus menjalankan aplikasi menggunakan server tersebut
	http.ListenAndServe(":8081", nil)
}

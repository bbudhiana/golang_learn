package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 1. Buat dulu struct yang dibutuhkan dan data sample nya
type student2 struct {
	ID    string
	Name  string
	Grade int
}

var data = []student2{
	{"E001", "ethan", 21},
	{"W001", "wick", 22},
	{"B001", "bourne", 23},
	{"B002", "bond", 23},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //hanya terima json

	if r.Method == "GET" {
		var result, err = json.Marshal(data) // data yang di-encode ke JSON

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result) //r.Write() digunakan untuk mendaftarkan data sebagai response
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		//Method r.FormValue() digunakan untuk mengambil data form yang dikirim dari client, pada konteks ini data yang dimaksud adalah ID
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	fmt.Println("Web service")

	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}

package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {
	fmt.Println("HASH")
	//Hash adalah algoritma enkripsi untuk mengubah text menjadi deretan karakter acak
	//Hash termasuk one-way encryption, membuat hasil dari hash tidak bisa dikembalikan ke text asli.

	//Implementasi hash
	//Go menyediakan cypto/sha1 untuk keperluan hashing
	var text = "this is secret"
	var sha = sha1.New() //buat variabel sha nya

	//Method Write() digunakan untuk menge-set data yang akan di-hash. Data harus dalam bentuk []byte
	sha.Write([]byte(text)) //sebelum dibuat sha nya, maka konversi dulu []byte

	//Method Sum() digunakan untuk eksekusi proses hash, menghasilkan data yang sudah di-hash dalam bentuk []byte
	//perlu parameter, bisa diisi dengan nil
	var encrypted = sha.Sum(nil)

	//%x untuk tulis format hexadecimal
	var encryptedString = fmt.Sprintf("%x", encrypted) //merupakan object bertipe hash.Hash

	fmt.Println(encryptedString)

	//Metode Salting Pada Hash SHA1
	//Salt dalam konteks kriptografi adalah data acak yang digabungkan pada data asli sebelum proses hash dilakukan
	var info, hasil_salt = doHashUsingSalt("jaja miharja")
	fmt.Printf("hasil hash+salt : %x, dengan salt: %x", info, hasil_salt)

}

// UNTUK SALT
func doHashUsingSalt(text string) (string, string) {
	//Hasilnya akan selalu unik setiap detiknya
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())

	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	fmt.Println("Text Salted :", saltedText)

	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

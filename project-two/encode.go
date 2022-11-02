package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("Encode/Decode")
	//Go punya package encoding/base64 untuk decode dan encode data ke base64.
	//data yg di encode harus type []byte, karenanya harus di casting agar sesuai
	var data = "bana budhiana"

	//proses encode string
	// harus di-casting terlebih dahulu ke dalam bentuk []byte
	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded :", encodedString) //YmFuYSBidWRoaWFuYQ==

	//proses decode ke string
	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("kembalikan encode string :", decodedString)

	//IMPLEMENTASI
	var data2 = "Bana Budhiana"

	//1. set dulu panjang encode nya
	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data2)))
	//2. lakukan encoded
	base64.StdEncoding.Encode(encoded, []byte(data2))
	//3. jadikan string
	var encodedString2 = string(encoded)
	fmt.Println(encodedString2)

}

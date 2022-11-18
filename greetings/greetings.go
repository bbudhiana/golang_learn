package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprintf(randomFormat())
	return message, nil
}

// fungsi Hellos mengembalikan sebuah Map berisi nama orang yg
// berhubungan dengan message nya
func Hellos(names []string) (map[string]string, error) {
	//sebuah map dengan nama dan pesannya
	messages := make(map[string]string)

	//lakukan looping untuk mendapatkan nama dari slice yg terinput
	//ambil pesan untuk setiap nama
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//tentukan pesan berdasarkan name yg diinput
		messages[name] = message
	}
	return messages, nil
}

// inisiasi untuk variabel yg digunakan dalam function
// Go executes init functions automatically at program startup, after global variables have been initialized
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat mengembalikan satu dari beberapa pesan
// pesan itu terpilih secara acak/random
func randomFormat() string {

	//Slice Message
	formats := []string{
		"Hi, %v. Welcome!",
		"Assalamu'alaikum %v, ahlan wasahlan",
		"Wilujeng enjing %v, selamat datang",
	}

	//kembalikan secara random dengan cara
	//menentukan index secara random dari slice data string nya
	return formats[rand.Intn(len(formats))]
}

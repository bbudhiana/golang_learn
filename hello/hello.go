package main

import (
	"fmt"
	"log"

	"bahaso.com/greetings"
)

func main() {

	//Set Properti untuk Log
	//isinya prefix log tentang apa dan flag untuk disable printing
	//waktu, sumber file, dan nomor baris program
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Reques a greeting message
	//message, err := greetings.Hello("Bana")

	//A slice of names
	names := []string{"bana", "budhiana", "Sultan Bogor"}

	messages, err := greetings.Hellos(names)

	//jika error maka print ke console
	//kemudian exit program
	if err != nil {
		log.Fatal(err)
	}

	//Get a greeting message and print it
	//message := greetings.Hello("Bana")
	fmt.Println(messages)
}

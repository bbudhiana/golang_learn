package main

import (
	//"fmt"
	f "fmt" //jadi alias "fmt" adalah f
	//"levelaccess/library"
	. "levelaccess/library" //diberi '.' agar pemanggilan function, struct, method nya tidak perlu pake 'library.' lagi
)

func main() {
	//library.SayHello()
	SayHello()

	//var murid1 = library.Murid{"Bana", 22}
	var murid1 = Murid{"Bana", 22}
	//fmt.Println("name :", murid1.Name)
	f.Println("name :", murid1.Name)

	//eksekusi dari partial.go, yang juga punya package main
	sayHellow("Johny Demonic")

	//hasil init() di library, Murid2 dideklarasikan di library.go
	f.Printf("Name   : %s \n", Murid2.Name)
	f.Printf("Grade  : %d \n", Murid2.Grade)
}

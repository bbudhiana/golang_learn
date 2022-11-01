package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Reflect")
	//Reflection adalah teknik untuk inspeksi variabel, mengambil informasi dari variabel tersebut atau bahkan memanipulasinya
	//2 fungsi yang paling penting untuk diketahui, yaitu reflect.ValueOf() dan reflect.TypeOf()
	//1. Fungsi reflect.ValueOf() akan mengembalikan objek dalam tipe reflect.Value, yang berisikan informasi
	//   yang berhubungan dengan NILAI pada variabel yang dicari
	//2. reflect.TypeOf() mengembalikan objek dalam tipe reflect.Type. Objek tersebut berisikan informasi
	//   yang berhubungan dengan TIPE DATA variabel yang dicari

	var number = 23
	var reflectValue = reflect.ValueOf(number)
	var nama = "abdullah"
	var reflectValue2 = reflect.ValueOf(nama)

	fmt.Println("tipe  variabel :", reflectValue.Type())
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel number :", reflectValue.Int()) //reflectValue.Int() untuk ambil data dalam int
	}
	if reflectValue2.Kind() == reflect.String {
		fmt.Println("nilai variabel nama :", reflectValue2.String())
	}

	//Pengaksesan Informasi Property Variabel Objek
	var s1 = &student{Name: "wick", Grade: 2}
	s1.getPropertyInfo()

	//Pengaksesan Informasi Method Variabel Objek
	var s2 = &student{Name: "john wick", Grade: 2}
	fmt.Println("nama :", s2.Name)

	var reflectValue3 = reflect.ValueOf(s2)
	var method = reflectValue3.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama :", s2.Name)
}

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	//pengecekan apakah variabel objek tersebut merupakan pointer atau tidak
	if reflectValue.Kind() == reflect.Ptr { //PTR is pointer
		//ka ternyata variabel memang pointer, maka perlu diambil objek reflect aslinya dengan cara memanggil method Elem()
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	//NumField() akan mengembalikan jumlah property publik yang ada dalam struct
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)         //nama :grade
		fmt.Println("tipe data :", reflectType.Field(i).Type)         //tipe data : int
		fmt.Println("nilai     :", reflectValue.Field(i).Interface()) //nilai : 2
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Konversi antar tipe data")

	//fungsi strconv.Atoi()  => string to int
	var str = "124"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num) // 124
	}

	//fungsi strconv.Itoa  => int to string
	var num2 = 124
	var str2 = strconv.Itoa(num2)

	fmt.Println(str2) // "124"

	//fungsi strconv.ParseInt()  => string to numeric non decimal
	var str3 = "124"
	var num3, err2 = strconv.ParseInt(str3, 10, 64) //str, base(10,2), bitsize (16,32,64)

	if err2 == nil {
		fmt.Println(num3) // 124
	}

	//fungsi strconv.FormatInt  => numeric int64 to string
	var num4 = int64(24)
	var str4 = strconv.FormatInt(num4, 8) //convert ke basis 8 dan ke dalam string

	fmt.Println(str4) // 30

	//fungsi strconv.ParseFloat()  => string to numeric decimal
	var str5 = "24.12"
	var num5, err5 = strconv.ParseFloat(str5, 32)

	if err5 == nil {
		fmt.Println(num5) // 24.1200008392334
	}

	//fungsi strconv.formatFloat()   =>  float64 to string, dengan eksponen, lebar decimal, dan lebar tipe bisa ditentukan
	var num6 = float64(24.12)
	var str6 = strconv.FormatFloat(num6, 'f', 6, 64)

	fmt.Println(str6) // 24.120000

	//fungsi strconv.ParseBool()   => string to boolean
	var str7 = "true"
	var bul, err7 = strconv.ParseBool(str7)

	if err7 == nil {
		fmt.Println(bul) // true
	}

	//fungsi strconv.FormatBool()   =>   boolean to string
	var bul2 = true
	var str8 = strconv.FormatBool(bul2)

	fmt.Println(str8) // true

	//Type Assertions Pada Interface Kosong
	var data = map[string]interface{}{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	fmt.Println(data["nama"].(string))      //menggunakan.(string)
	fmt.Println(data["grade"].(int))        //menggunakan .(int)
	fmt.Println(data["height"].(float64))   //menggunakan .(float64)
	fmt.Println(data["isMale"].(bool))      //menggunakan .(bool)
	fmt.Println(data["hobbies"].([]string)) //menggunakan .([]string)) untuk konvert ke slice berisi string
}

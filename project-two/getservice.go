package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// data adhan
type dataadhan struct {
	Day     string
	Fajr    string
	Sunrise string
	Zuhr    string
	Asr     string
	Maghrib string
	Isha    string
}

// dataadhaninmonth
type dataadhaninmonth struct {
	Timings dataadhan
	Date    map[string]interface{}
}

// Get Data Adhan from server
func fetchAdhan() (map[string]interface{}, error) {
	var err error

	//Statement &http.Client{} menghasilkan instance http.Client. Objek ini nantinya diperlukan untuk eksekusi request.
	var client = &http.Client{}
	var data []dataadhaninmonth

	output := map[string]interface{}{
		"code":      200,
		"status":    "OK",
		"latitude":  -6.178306,
		"longitude": 106.631889,
		"data":      data,
	}

	//Fungsi http.NewRequest() digunakan untuk membuat request baru
	// Parameter pertama, berisikan tipe request POST atau GET atau lainnya
	// Parameter kedua, adalah URL tujuan request
	// Parameter ketiga, form data request (jika ada)
	//Fungsi tersebut menghasilkan instance bertipe http.Request
	request, err := http.NewRequest("GET", baseURL+"/adhan?latitude=-6.178306&longitude=106.631889&month=10&year=2022", nil)
	if err != nil {
		return nil, err
	}

	//eksekusi request dengan Do
	//Method tersebut mengembalikan instance bertipe http.Response
	response, err := client.Do(request)
	fmt.Println(response)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// base URL
var baseURL = "http://localhost:8080"

func main() {
	fmt.Println("Get Data Sholat")

	var adhans, err = fetchAdhan()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	fmt.Println(adhans["data"])

	/*
		for _, each := range adhans {
			fmt.Printf("Code : %s\n", each["code"])
		}
	*/
}

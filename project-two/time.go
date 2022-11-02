package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("TIME")
	//time.Now() waktu sekarang
	//time.Date() waktu custom
	//keduanya bertipe time.Time
	var time1 = time.Now()
	fmt.Printf("time : %v\n", time1)
	fmt.Println("Tahun : ", time1.Year())

	var time2 = time.Date(2022, 10, 28, 10, 10, 10, 100, time.Local)
	fmt.Printf("time 2 : %v\n", time2)

	//PARSING dari string ke time.Time
	//butuh 2 parameter :
	//1. layout format dari data yg akan diparsing
	//2. data string yang ingin diparsing
	var layoutFormat, value string
	var date time.Time

	//dokumen time parsing : https://dasarpemrogramangolang.novalagung.com/A-time-parsing-format.html
	layoutFormat = "2006-01-02 15:04:05" //ini standarnya GO untuk yyyy-mm-dd
	value = "2001-11-11 11:11:11"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	//Predefined Layout Format Untuk Keperluan Parsing Time
	var date2, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	fmt.Println(date2.String()) //// 2015-09-02 08:00:00 +0700 WIB

	//Format dari time.Time ke string
	var date3, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")

	var dateS1 = date3.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1)
	// Wednesday 02, September 2015 08:00 WIB

	var dateS2 = date3.Format(time.RFC3339)
	fmt.Println("dateS2", dateS2)

	//time.Sleep(second)
	fmt.Println("start")
	time.Sleep(time.Second * 4)
	fmt.Println("after 4 seconds")

	// fungsi timer.After() akan mengembalikan data channel, sehingga perlu menggunakan tanda <- dalam penerapannya
	<-time.After(4 * time.Second)
	fmt.Println("expired")

	//duration
	start := time.Now()

	time.Sleep(5 * time.Second)
	fmt.Println("timsecond :", time.Second) //1s

	duration := time.Since(start)

	fmt.Println("time elapsed in seconds:", duration.Seconds())
	fmt.Println("time elapsed in minutes:", duration.Minutes())
	fmt.Println("time elapsed in hours:", duration.Hours())

	//kalkulasi durasi 2 waktu
	t1 := time.Now()
	time.Sleep(5 * time.Second)
	t2 := time.Now()

	duration2 := t2.Sub(t1)

	fmt.Println("time elapsed in seconds:", duration2.Seconds())
	fmt.Println("time elapsed in minutes:", duration2.Minutes())
	fmt.Println("time elapsed in hours:", duration2.Hours())
}

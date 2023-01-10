package main

import "fmt"

// definisi fungsi dengan output salah satunya adalah berupa function closure
func findMaxs(numbers []int, max int) (int, func() []int) {
	var res []int //tampungan
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	//function closure
	var fncClosure = func() []int {
		return res
	}
	return len(res), fncClosure
}

func main() {
	fmt.Println("Contoh Function")

	//Function Closure dan return multi value
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				min, max = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	var numbers = []int{1, 2, 3, 14, 5, 6, 7, 8, 9}
	var min, max = getMinMax(numbers)
	fmt.Printf("Min: %v dan Max: %v\n", min, max)

	//Immediately-Invoked Function Expression (IIFE)
	//fungsi closure yg langsung dieksekusi
	var angkaku = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var nomerBaru = func(min int) []int {
		var r []int //tampungan
		for _, e := range angkaku {
			if e < min {
				continue
			}
			r = append(r, e) //gabungkan dengan var tampungan
		}
		return r
	}(7) // langsung diisikan dengan 3 sebagai variabel min int
	fmt.Printf("Angka sebelummnya : %v\n", angkaku)
	fmt.Printf("Angka lebih gede or sama dari input : %v\n", nomerBaru)

	//PENGGUNAAN FUNGSI DENGAN RETURN CLOSURE
	var themax = 3
	var thenumbers = []int{2, 3, 4, 7, 3, 3, 5, 7, 9, 6, 5, 3, 2}
	var hoMany, getNumbers = findMaxs(thenumbers, themax)
	var theNumbers = getNumbers()

	fmt.Printf("Banyaknya item : %v\n", hoMany)
	fmt.Printf("Nomor-nomor itu : %v\n", theNumbers)

}

package main

import (
	"fmt"
	"math"
	"math/rand"
	"rsc.io/quote"
	"strings"
	"time"
)

func main() {
	//Penggunaan import fmt
	fmt.Println("Hello World")
	fmt.Println("Hello", "world!", "how", "are", "you")
	fmt.Println(quote.Go())

	/*
	  var <nama-variabel> <tipe-data>
	  var <nama-variabel> <tipe-data> = <nilai>
	*/
	var firstname = "Bana"
	var lastname string
	lastname = "Budhiana"
	fmt.Printf("Hello %s %s \n", firstname, lastname)

	//deklarasi variable yg ditentukan oleh value nya disebut sebagai type inference
	//Tanda := hanya digunakan sekali di awal pada saat deklarasi
	var namapertama = "Bana"
	namaakhir := "Budhiana" //otomatis namaakhir bertipe string - type inference
	fmt.Printf("Hello %s %s \n", namapertama, namaakhir)

	//deklarasi banyak variabel
	var firstWife, secondWife, thirdWife string
	firstWife = "ani"
	secondWife = "sri"
	thirdWife = "andriani"
	fmt.Printf("Wife %s %s %s\n", firstWife, secondWife, thirdWife)

	//deklarasi banyak variabel secara langsung
	var fourth, fifth, sixth = "empat", "lima", "enam"

	//ringkas
	//seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	fmt.Printf("Angka %s %s %s\n", fourth, fifth, sixth)

	//variable sampah '_' digunakan untuk tampung nilai yg tidak dipakai
	//Biasanya variabel underscore sering dimanfaatkan untuk menampung nilai balik fungsi yang tidak digunakan
	_ = "Golang itu mudah"

	//Keyword new digunakan untuk membuat variabel pointer dengan tipe data tertentu
	//Jika ditampilkan yang muncul bukanlah nilainya melainkan alamat memori nilai tersebut (dalam bentuk notasi heksadesimal)
	name := new(string)
	fmt.Println(name) // 0x20818a220

	//'make' hanya bisa digunakan untuk pembuatan beberapa jenis variabel saja, yaitu : channel, slice, map

	//Tipe data Numeric non-decimal
	//uint8 = 0-255
	//uint16= 0-65535
	//uint32= 0-4294967295
	//uint64= 0-18446744073709551615
	//uint = uint32 or uint64 (tergantung nilai)
	//byte = unit8
	//int8 = -128-127
	//int16 = -32768 - 32767
	//int32 = -2147483648 - 2147483647
	//int64 = -9223372036854775808 - 9223372036854775807
	//int = sama dengan int32 or int64 (tergantung nilai)
	//rune = int32
	var positiveNumber uint8 = 89
	var negatifNumber = -1232229
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negatifNumber)

	//Tipe data Numeric Decimal
	//Template %f digunakan untuk memformat data numerik desimal menjadi string
	var decimalNumber = 2.62
	fmt.Printf("ini adalah float number : %.3f\n", decimalNumber)   //.3 artinya adalah 3 angka belakang koma
	fmt.Printf("ini outputnya 10 decimal : %.10f\n", decimalNumber) //default 6 decimal, tapi bisa di set 10 decimal

	//Tipe data bool
	var exist = true
	fmt.Printf("ini boolean: %t\n", exist)

	//Tipe data yang bisa diisi dengan nill
	//1. pointer
	//2. tipe data fungsi
	//3. slice
	//4. map
	//5. channel
	//6. interface kosong atau 'any'

	//Konstanta
	//variable yang tidak bisa berubah nilainya setelah diinisialisasikan
	const namaOrtu string = "Budhi"
	fmt.Print("Halo ", namaOrtu, "\n")

	//Multi deklarasi
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	//perhitungan dan logika
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = value == 2
	fmt.Printf("nilai %d (%t) \n", value, isEqual) //%t digunakan untuk output boolean

	//kondisi, dalam if gak perlu kurung lagi
	var point = 8

	if point == 10 {
		fmt.Println("OK")
	} else {
		fmt.Println("Not Ok")
	}

	//variabel dalam if yg temporary, hanya bisa digunakan dalam block if
	var point2 = 8850.0
	if percent := point2 / 100; percent >= 100 { //percent adalah varibel temporer
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else {
		fmt.Printf("Not OK your %f", percent)
	}

	//switch
	// Di Go, ketika sebuah case terpenuhi, tidak akan dilanjutkan ke pengecekan case selanjutnya,
	// meskipun tidak ada keyword break di situ
	var point3 = 6

	switch point3 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	//Switch Dengan Gaya if - else
	//Keyword fallthrough digunakan untuk memaksa proses pengecekan diteruskan
	//ke satu case selanjutnya dengan tanpa menghiraukan nilai kondisinya
	var point4 = 6

	switch {
	case point4 == 8:
		fmt.Println("perfect")
	case (point4 < 8) && (point4 > 3):
		fmt.Println("awesome")
		fallthrough
	case point4 < 5:
		fmt.Println("you need to learn more") //apapun prosesnya, ini tetap dieksekusi
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//Di Go keyword perulangan hanya for saja
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	//for Dengan Argumen Hanya Kondisi, konsepnya mirip while
	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	//for Tanpa Argumen
	var j = 0

	for {
		fmt.Println("Angka", j)

		j++
		if j == 5 {
			break //tanda berhentinya pengulangan
		}
	}

	//Keyword break digunakan untuk menghentikan secara paksa sebuah perulangan,
	//sedangkan continue dipakai untuk memaksa maju ke perulangan berikutnya
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue //jika menghasilkan sisa bagi, maka lanjut ke pengulangan berikut
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	//perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	//ARRAY
	//Default nilai tiap elemen array pada awalnya tergantung dari tipe datanya
	//Jika int maka tiap element zero value-nya adalah 0
	//jika bool maka false
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	//langsung inisialisasi data, pakai {}
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	//Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	//cukup ganti dengan tanda 3 titik (...)
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	//Array Multidimensi
	var numbers1 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	//Perulangan Array dengan for
	var fruits2 = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits2); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits2[i])
	}

	//Perulangan Array dengan for - range
	var fruits3 = [4]string{"apple", "grape", "banana", "melon"}
	//fruits diambil berurutan dan ditampung oleh variabel fruit
	//sedangkan indeks nya ditampung variabel i
	for i, fruit := range fruits3 {
		fmt.Printf("elemen %d : %s\n", i, fruit) // 'i' harus dikeluarkan, karena go mewajibkan variabel digunakan
	}

	//kadangkala yg dibutuhkan adalah elemen-nya, index tidak perlu
	//Di sinilah salah satu kegunaan variabel pengangguran, atau underscore (_).
	//Tampung saja nilai yang tidak ingin digunakan ke underscore.
	var fruits4 = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruits4 {
		fmt.Printf("nama buah : %s\n", fruit) //jadi tidak perlu lagi ada 'i'
	}

	//Deklarasi sekaligus alokasi data array juga bisa dilakukan lewat keyword make.
	var fruits5 = make([]string, 2) //buat array string dengan jumlah elemen = 2
	fruits5[0] = "apple"
	fruits5[1] = "manggo"

	fmt.Println(fruits5) // [apple manggo]

	//SLICE
	//Slice adalah reference elemen array
	//Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya.
	//menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama
	var fruits6 = []string{"apple", "grape", "banana", "melon"} //ciri slice, tidak dituliskan jumlah elemen
	fmt.Println(fruits6[0])                                     // "apple"

	//perbedaan slice dan array bisa diketahui pada saat deklarasi variabel-nya,
	//jika jumlah elemen tidak dituliskan, maka variabel tersebut adalah slice
	var fruitsA = []string{"apple", "grape"}     // slice
	var fruitsB = [2]string{"banana", "melon"}   // array
	var fruitsC = [...]string{"papaya", "grape"} // array
	fmt.Println(fruitsA, fruitsB, fruitsC)

	//Array adalah kumpulan nilai atau elemen, sedang slice adalah referensi tiap elemen tersebut.
	var fruits7 = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits7[0:2] //akses elemen slice 'fruits7'dari index 0 hingga sebelum index ke-2

	fmt.Println(newFruits) // ["apple", "grape"] adalah slice baru

	//Slice merupakan tipe data reference atau referensi
	//jika ada slice baru yang terbentuk dari slice lama, maka data elemen slice yang baru
	//akan memiliki alamat memori yang sama dengan elemen slice lama
	var fruits8 = []string{"apple", "grape", "banana", "melon"}
	var bFruits = fruits8[1:4]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits8)  // [apple grape banana melon]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(baFruits) // [grape]

	baFruits[0] = "pinnaple" //ubah grape jadi pinnaple
	fmt.Println(fruits8)     // [apple pinnaple banana melon]

	//len and cap
	var fruits9 = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits9)) // len: 4
	fmt.Println(cap(fruits9)) // cap: 4

	//append
	//Fungsi append() digunakan untuk menambahkan elemen pada slice
	//Elemen baru tersebut diposisikan setelah indeks paling akhir
	var fruits10 = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits10, "papaya")
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]

	//MAP
	//tipe data asosiatif yang ada di Go, berbentuk key-value pair.
	//Default nilai variabel map adalah nil. perlu dilakukan inisialisasi nilai default di awal
	//map[string]int{}
	var chicken map[string]int //string=key , int= value
	chicken = map[string]int{} //inisialisasi, jika tidak dilakukan akan error

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	//Zero value dari map adalah nil, maka tiap variabel bertipe map harus di-inisialisasi
	//secara explisit nilai awalnya (agar tidak nil).
	var data map[string]int
	//data["one"] = 1 // akan muncul error! karena tidak diinisiasi dengan map[string]int{}

	data = map[string]int{}
	data["one"] = 1 // tidak ada error

	//Definisi dari awal, mirip seperti Slice
	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara vertical
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	fmt.Println(chicken1, chicken2)

	//iterasi Item Map Menggunakan for - range
	var chicken3 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken3 {
		fmt.Println(key, "  \t:", val) // \t untuk tabulasi
	}

	//Menghapus Item Map
	//Fungsi delete() digunakan untuk menghapus item dengan key tertentu pada variabel map
	var chicken4 = map[string]int{"januari": 50, "februari": 40}
	delete(chicken4, "januari") //delete dari chicken4 data map dengan key="januari"
	fmt.Println(chicken4)       //map[februari:40]

	//check elemen dalam map, var nilai, existensi = chicken4['nama_bulan'], jika ada maka 'true', selain itu 'false'
	var value2, isExist = chicken4["mei"]                      //menampung value dan key
	fmt.Printf("eksistensi : %t nilai: %d\n", isExist, value2) //isExist=false  value2=0

	//Kombinasi Slice & Map
	//[]map[string]int, artinya slice yang tipe tiap elemen-nya adalah map[string]int
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	//Go terbaru mempersingkat slice dan map
	var chickens5 = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	for index, chicken := range chickens5 {
		fmt.Println(chicken["gender"], chicken["name"], index)
	}

	//FUNGSI
	//Definisi fungsi sendiri adalah sekumpulan blok kode yang dibungkus dengan nama tertentu.
	//Fungsi main merupakan fungsi yang paling utama pada program Go
	var names2 = []string{"John", "Wick"}
	printMessage("halo", names2)

	//Fungsi Dengan Return Value / Nilai Balik
	rand.Seed(time.Now().Unix()) //Fungsi ini diperlukan untuk memastikan bahwa angka random yang akan di-generate benar-benar acak.
	var randomValue int
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	//return dalam fungsi berguna juga untuk menghentikan proses seperti break
	divideNumber(10, 2)
	divideNumber(4, 0)

	//FUNGSI MULTIPLE RETURN
	//Jika ada kebutuhan di mana data yang dikembalikan harus banyak,
	//biasanya digunakanlah tipe seperti map, slice, atau struct sebagai nilai balik.
	var diameter float64 = 15
	var area, circumference = calculate(diameter) //tampung 2 return dengan 2 variabel

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	//Fungsi Variadic
	//pembuatan fungsi dengan parameter sejenis yang tak terbatas
	//hampir sama dengan deklarasi biasa, namun ditambah (...)
	var avg = calculate3(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

	//Pengisian Parameter Fungsi Variadic Menggunakan Data Slice
	//Slice bisa digunakan sebagai parameter variadic.
	//Caranya dengan menambahkan tanda titik tiga kali, tepat setelah nama variabel yang dijadikan parameter
	var numbers3 = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg2 = calculate3(numbers3...) //slice sebagai parameter, tinggal tambah '...'
	var msg2 = fmt.Sprintf("Rata-rata : %.2f", avg2)

	fmt.Println(msg2)

	//Fungsi Dengan Parameter Biasa & Variadic
	//Parameter variadic bisa dikombinasikan dengan parameter biasa, dengan syarat parameter variadic-nya harus diposisikan di akhir.
	yourHobbies("wick", "sleeping", "eating", "ngising")

	//dengan slice maka gunakan '...' di akhiran parameter input nya
	var hobbies = []string{"sleeping", "eating", "drinking"}
	yourHobbies("wick", hobbies...)

	//Fungsi Closure
	//Definisi Closure adalah sebuah fungsi yang bisa disimpan dalam variabel
	//Closure merupakan anonymous function atau fungsi tanpa nama
	//1.Closure Disimpan Sebagai Variabel
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	var numbers4 = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers4)

	//Template %v digunakan untuk menampilkan segala jenis data. Bisa array, int, float, bool, dan lainnya
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers4, min, max)

	//2. Immediately-Invoked Function Expression (IIFE)
	//Closure jenis ini dieksekusi langsung pada saat deklarasinya
	var numbers5 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int //variabel untuk tampung return
		for _, e := range numbers5 {
			if e < min { //jika lebih kecil dari min yg ditentukan maka lewatkan
				continue
			}
			r = append(r, e) //selain itu gabung ke variabel tampungan
		}
		return r
	}(3)

	fmt.Println("original number :", numbers5)
	fmt.Println("filtered number :", newNumbers)

	//3. Closure Sebagai Nilai Kembalian
	//Salah satu keunikan closure lainnya adalah bisa dijadikan sebagai nilai balik fungsi
	var max2 = 3
	var numbers6 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numbers6, max2)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers6)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]

	//FUNCTION sebagai parameter
	//Cara membuat parameter fungsi adalah dengan langsung menuliskan skema fungsi nya sebagai tipe data
	var data2 = []string{"wick", "jason", "ethan"}

	//closure filter untuk cari yg contain huruf "o"
	var dataContainsO = filter(data2, func(each string) bool {
		return strings.Contains(each, "o") //apakah berisi "o"
	})

	//closure filter untuk cari yg panjangnya 5
	var dataLenght5 = filter(data2, func(each string) bool {
		return len(each) == 5 //apakah berjumlah 5 karakter
	})

	fmt.Println("data asli \t\t:", data2)
	// data asli : [wick jason ethan]

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// filter ada huruf "o" : [jason]

	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)
	// filter jumlah huruf "5" : [jason ethan]

	//Jika Function sebagai parameter mulai rumit, maka perlu Alias Skema Closure
	//Membuat alias fungsi berarti menjadikan skema fungsi tersebut menjadi tipe data baru. Caranya dengan menggunakan keyword type
	/*
		type FilterCallback func(string) bool 				//type untuk jadikan alias

		func filter(data []string, callback FilterCallback) []string {
		    // ...
		}
	*/

	//POINTER
	//Pointer adalah reference atau alamat memori. Variabel pointer berarti variabel yang berisi alamat memori suatu nilai.
	//Variabel-variabel yang memiliki reference atau alamat pointer yang sama, saling berhubungan satu sama lain dan nilainya pasti sama
	//Nilai default variabel pointer adalah nil (kosong). Variabel pointer tidak bisa menampung nilai yang bukan pointer,
	//dan sebaliknya variabel biasa tidak bisa menampung nilai pointer
	//1. Penerapan pointer
	//a. Variabel bertipe pointer ditandai dengan adanya tanda asterisk (*) tepat sebelum penulisan tipe data ketika deklarasi.
	//b. Variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan tanda ampersand (&) tepat sebelum nama variabel.
	//Metode ini disebut dengan referencing
	//c. Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda asterisk (*) tepat sebelum nama variabel.
	//Metode ini disebut dengan dereferencing.
	var numberA = 4
	var numberB = &numberA //numberB menampung alamat memory numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc00019c220 (alamat pointernya atau memori, pakai &)

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc00019c220

	//Ketika salah satu variabel pointer di ubah nilainya, sedang ada variabel lain yang memiliki referensi memori yang sama,
	//maka nilai variabel lain tersebut juga akan berubah
	var numberC = 4
	var numberD = &numberA

	fmt.Println("numberC (value)   :", numberC)
	fmt.Println("numberC (address) :", &numberC)
	fmt.Println("numberD (value)   :", *numberD)
	fmt.Println("numberD (address) :", numberD)

	fmt.Println("")

	numberC = 5

	fmt.Println("numberC (value)   :", numberC)
	fmt.Println("numberC (address) :", &numberC)
	fmt.Println("numberD (value)   :", *numberD)
	fmt.Println("numberD (address) :", numberD)

	//Parameter Pointer
	//Parameter bisa juga dirancang sebagai pointer
	var number = 4
	fmt.Println("before :", number) // 4

	change(&number, 10)
	fmt.Println("after  :", number) // 10

}

// data pertama adalah string "hallo" yang ditampung parameter message,
// dan parameter ke 2 adalah slice string names yang nilainya ditampung oleh parameter arr.
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ") //value array di join dengan operasi join
	fmt.Println(message, nameString)
}

// nilai baliknya int
func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value //return untuk kembalikan nilai
}

func divideNumber(m, n int) {
	//jika ternyata parameter 2 adalah 0 maka perhitungan dihentikan untuk menghindari error di proses berikutnya (pembagian 0)
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return //dengan return, proses fungsi tidak diteruskan ke bawahnya
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

// multiple return function
// Fungsi calculate() menerima satu buah parameter (diameter), dengan 2 return (float64, float64)
func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)

	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

// Fungsi Dengan Predefined Return Value
// di Go variabel yang digunakan sebagai nilai balik bisa didefinisikan di awal
// fungsi calculate2 adalah penyederhanaan dari fungsi calculate
func calculate2(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	//paling menarik disini, karena tipe return sudah ditentukan awal maka hanya perlu 'return'
	return
}

// Fuction Variadic, isian parameter bisa dinamis jumlahnya yaitu ada '...'
func calculate3(numbers ...int) float64 {
	var total = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

// Function biasa dan variadic
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

// return berupa closure
// fungsi ini digunakan untuk mencari banyaknya angka-angka yang nilainya di bawah atau sama dengan angka tertentu
func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	//return berupa banyaknya elemen input yg <= dan function ambil numbers tersebut
	return len(res), func() []int {
		return res
	}
}

// fungsi sebagai parameter
// Parameter callback merupakan sebuah closure yang dideklarasikan bertipe func(string) bool
// Fungsi filter() sendiri kita buat untuk filtering data array (yang datanya didapat dari parameter pertama),
// dengan kondisi filter bisa ditentukan sendiri.
func filter(data []string, callback func(string) bool) []string {
	var result []string //data tampungan untuk di return
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each) //jika bernilai true dari function parameternya, maka append ke tampungan
		}
	}
	return result
}

// buat ubah parameter pointer
func change(original *int, value int) {
	*original = value
}

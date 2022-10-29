package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello", "world!", "how", "are", "you")

	/*
	  var <nama-variabel> <tipe-data>
	  var <nama-variabel> <tipe-data> = <nilai>
	*/
	var firstname = "Bana"
	var lastname string
	lastname = "Budhiana"
	fmt.Printf("Hello %s %s \n", firstname, lastname)

	//deklarasi variable yg ditentukan oleh value nya - type inference
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
	fmt.Printf("bilangan positif: %d\n", positiveNumber)

	//Tipe data Numeric Decimal
	//Template %f digunakan untuk memformat data numerik desimal menjadi string
	var decimalNumber = 2.62
	fmt.Printf("ini adalah float number : %.3f\n", decimalNumber)

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
	fmt.Printf("nilai %d (%t) \n", value, isEqual)

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

}

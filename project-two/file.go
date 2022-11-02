package main

import (
	"fmt"
	"io"
	"os"
)

var path = "/home/bahaso/Pribadi/go/golang_learn/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err != nil
}

func createFile() {
	//deteksi dulu apakah file masih exist
	//os.Stat return 2 parameter : info file yg dicari dan error jika ada
	var _, err = os.Stat(path)

	//buat file jika belum ada file nya
	if os.IsNotExist(err) {
		//os.Create hasilkan 2 return : file dan error
		var file, err = os.Create(path)

		if isError(err) {
			return
		}

		//File yang baru terbuat statusnya adalah otomatis open, karenanya perlu di close
		defer file.Close()
	}
	fmt.Println("====> file berhasil dibuat <===== : ", path)
}

func writeFile() {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString("halo\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("mari belajar golang\n")
	if isError(err) {
		return
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di isi")
}

func readFile() {
	// buka file
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil dibaca")
	fmt.Println(string(text))
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di delete")
}
func main() {
	fmt.Println("Operasi file")
	createFile()
	writeFile()
	readFile()
	deleteFile()

}

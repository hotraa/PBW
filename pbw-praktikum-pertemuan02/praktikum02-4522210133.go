package main

import (
	"bufio"
	"fmt"
	"os"
)

type Mahasiswa struct {
	Nama    string
	NPM     string
	Jurusan string
}

func main() {
	// Membuat map kosong untuk menyimpan data mahasiswa
	mahasiswa := make(map[string]Mahasiswa)

	fmt.Print("\n\t========== Pencarian Data Mahasiswa ==========\n\n")
	// Memasukkan data mahasiswa
	inputDataMahasiswa(mahasiswa)

	// Menampilkan daftar nama mahasiswa
	tampilkanNamaMahasiswa(mahasiswa)

	// Mencari data mahasiswa berdasarkan NPM
	cariMahasiswaByNPM(mahasiswa)
}

// Function untuk memasukkan data mahasiswa
func inputDataMahasiswa(mahasiswa map[string]Mahasiswa) {
	var npm, jurusan string
	var jumlahMahasiswa int

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scanln(&jumlahMahasiswa)

	for i := 0; i < jumlahMahasiswa; i++ {
		fmt.Printf("\nMasukkan nama mahasiswa ke-%d\t: ", i + 1)
		scanner.Scan()
		nama := scanner.Text()
		fmt.Printf("Masukkan NPM mahasiswa ke-%d\t: ", i + 1)
		fmt.Scanln(&npm)
		fmt.Printf("Masukkan jurusan mahasiswa ke-%d\t: ", i+1)
		scanner.Scan()
		jurusan = scanner.Text()

		// Menambahkan data mahasiswa ke dalam map
		mahasiswa[npm] = Mahasiswa{Nama: nama, NPM: npm, Jurusan: jurusan}
	}
}

// Function untuk menampilkan daftar nama mahasiswa
func tampilkanNamaMahasiswa(mahasiswa map[string]Mahasiswa) {
	fmt.Println("\nDaftar Nama Mahasiswa:")
	i := 1
	for _, mhs := range mahasiswa {
		fmt.Printf("%d. %s\n", i, mhs.Nama)
		i++
	}
}

// Function untuk mencari data mahasiswa berdasarkan NPM
func cariMahasiswaByNPM(mahasiswa map[string]Mahasiswa) {
	var npm string

	fmt.Print("\nMasukkan NPM untuk mencari data mahasiswa: ")
	fmt.Scanln(&npm)

	mhs, ada := mahasiswa[npm]
	if ada {
		fmt.Printf("\nData Mahasiswa:\nNama\t: %s\nNPM\t: %s\nJurusan\t: %s\n", mhs.Nama, mhs.NPM, mhs.Jurusan)
		fmt.Print("\n")
	} else {
		fmt.Println("Data Mahasiswa tidak ada!")
		fmt.Print("\n")
	}
}

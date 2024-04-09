package main

import (
	"fmt"
	"os"
)

// Fungsi dibawah sama saja dengan
// 7 x (6 x (5 x (4 x (3 x (2 x (1 x 1))))))

func faktorial(value int) int {
	if value == 1 {
		return 1
	}
	return value * faktorial(value - 1)
}

func main() {
	var angkaFaktorial int

	fmt.Print("Masukkan angka faktorial: ")
	_, inputAngkaError := fmt.Scanln(&angkaFaktorial)

	if inputAngkaError != nil || angkaFaktorial <= 0 {
		fmt.Println("\nInputan salah! Masukkan angka bilangan bulat positif\n")
		os.Exit(1)
	}

	fmt.Printf("Hasil faktorial dari %d adalah ", angkaFaktorial)

	fmt.Println(faktorial(angkaFaktorial))
}

package main

import (
	"fmt"
	"os"
)

func main() {
	// Membuat direktori baru.
	err := os.Mkdir("Muhammad Abduh Harry Malhotra", 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Direktori 'Muhammad Abduh Harry Malhotra' berhasil dibuat.")
}

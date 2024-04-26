package main

import (
	"fmt"
	"os"
)

func main() {
	var err error

	// Mengubah izin direktori.
	err = os.Chmod("PEBEWE", 0133)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Izin 'PEBEWE' telah diubah menjadi 0133.")
}

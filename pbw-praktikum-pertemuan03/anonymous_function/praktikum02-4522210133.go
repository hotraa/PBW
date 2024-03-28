package main

import (
	"fmt"
	"math"
)

func hitungLingkaran(jariJari float64) (float64, float64) {
	luas := math.Pi * jariJari * jariJari
	keliling := 2 * math.Pi * jariJari
	return luas, keliling
}

func main() {
	luas, keliling := hitungLingkaran(5.0)

	pembulatanLuas := fmt.Sprintf("%.2f", luas)
	pembulatanKeliling := fmt.Sprintf("%.2f", keliling)

    fmt.Println("")
	fmt.Println("Luas lingkaran:", pembulatanLuas)
	fmt.Println("Keliling lingkaran:", pembulatanKeliling)
    fmt.Println("")
}

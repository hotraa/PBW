package main

import "fmt"

func main() {
    fmt.Println("Bubble Sort")
    var arrayNumber [20]float32

	fmt.Println("Masukkan angka yang akan diurutkan:\n")
    for i := 0; i < len(arrayNumber); i++ {
        fmt.Printf("Angka ke-%d: ", i + 1)
        fmt.Scan(&arrayNumber[i])
    }

    for i := 0; i < len(arrayNumber); i++ {
        for j := 0; j < len(arrayNumber)-i-1; j++ {
            if arrayNumber[j] > arrayNumber[j + 1] {
                x := arrayNumber[j]
                arrayNumber[j] = arrayNumber[j + 1]
                arrayNumber[j + 1] = x
            }
        }
    }

    fmt.Println("\nSetelah dilakukan pengurutan:")
    fmt.Println(arrayNumber)
}

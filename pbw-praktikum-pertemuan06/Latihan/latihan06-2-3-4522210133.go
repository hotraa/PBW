package main

import (
	"fmt"
	"os"
)

func main() {
	var err error

	fileInfo, err := os.Stat("PEBEWE")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if fileInfo.IsDir() {
		fmt.Println("PEBEWE adalah sebuah direktori.")
	} else {
		fmt.Println("PEBEWE adalah sebuah file.")
	}
}

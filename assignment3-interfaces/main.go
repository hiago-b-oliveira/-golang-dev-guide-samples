package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid number of arguments")
		os.Exit(1)
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}

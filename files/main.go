package main

import (
	"fmt"
	"io"
	"os"
)

type File struct{}

func main() {

	fmt.Println(os.Args[1])

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	io.Copy(os.Stdout, file)

}

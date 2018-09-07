package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type fileWriter struct{}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	fw := fileWriter{}

	io.Copy(fw, file)
}

func (fileWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote out this many bytes", len(bs))

	return len(bs), nil
}

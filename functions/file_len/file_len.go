package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println(fileLen("teste.txt"))
}

func fileLen(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data := make([]byte, 2048)
	total := 0

	for {
		count, err := file.Read(data)
		total += count
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
	return total, nil
}

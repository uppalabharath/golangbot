package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fpath := flag.String("fpath", "../default.txt", "Path of the file to read the contents from")
	flag.Parse()
	file, err := os.Open(*fpath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal("Error occurred while closing the file", err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal("Error occurred while reading the file", err)
	}
}

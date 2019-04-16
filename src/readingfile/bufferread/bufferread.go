package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fPath := flag.String("fpath", "../default.txt", "Path of the text file ro read")
	flag.Parse() // Parses the fpath command arg
	f, err := os.Open(*fPath)
	if err != nil {
		log.Fatal("Error occurred while opening file", err)
		return
	}
	reader := bufio.NewReader(f)
	buffer := make([]byte, 3)
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal("error while closing the file", err)
		}
	}()
	for {
		_, err := reader.Read(buffer)
		if err != nil {
			log.Fatal("Error occurred while reading file contents", err)
		}
		fmt.Println(string(buffer))
	}
}

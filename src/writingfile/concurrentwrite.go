package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
)

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	writer, err := os.Create("concurrent.txt")
	if err != nil {
		log.Fatal("Error while creating the file", err)
		done <- false
		return
	}
	for randomInt := range data {
		_, err = fmt.Fprintln(writer, randomInt)
		if err != nil {
			log.Fatal(err)
			writer.Close()
			done <- false
			return
		}
	}
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
		done <- false
	}
	done <- true
}

func main() {

	data := make(chan int)
	done := make(chan bool)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}

	go consume(data, done)

	go func() {
		wg.Wait()
		close(data)
	}()

	d := <-done

	if d {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("Some error occurred")
	}
}

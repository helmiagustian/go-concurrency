package main

import (
	"fmt"
	"strings"
)

func main() {
	uni()
}

func simple1() {
	foo := make(chan string, 2) // bidirectional buffered channel
	foo <- "Hello World"        // send
	bar := <-foo                // receive
	fmt.Println(bar)
}

func simple2() {
	foo := make(chan string) // bidirectional unbuffered channel
	// foo <- "Hello World"  // deadlock
	go func() { foo <- "Hello World" }() // send
	bar := <-foo                         // receive
	fmt.Println(bar)
}

func uni() {
	ch := make(chan int, 10)
	makeEvenNums(4, ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	receiveNum(ch)
	receiveNum(ch)
}

func makeEvenNums(count int, in chan<- int) { // send only channel
	for i := 0; i < count; i++ {
		in <- 2 * i // send
	}
}

func receiveNum(out <-chan int) { // receive only channel
	fmt.Println(<-out) // receive
}

func synchronization() {
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink ...",
		"The dark bird bird of prey lands on a small ...",
	}
	histogram := make(map[string]int)
	done := make(chan bool)
	// splits and count words
	go func() {
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				histogram[word]++
			}
		}
		done <- true
	}()
	if <-done {
		for k, v := range histogram {
			fmt.Printf("%s\t(%d)\n", k, v)
		}
	}
}

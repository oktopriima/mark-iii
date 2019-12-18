package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	go f("goroutine")

	f("direct")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second * 5)
	fmt.Println("done")
}

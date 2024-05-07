package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "heesu", "larry"}
	for _, person := range people {
		go isCool(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

}

func isCool(person string, c chan string) {
	time.Sleep(time.Second)
	c <- person + " is Cool"
}

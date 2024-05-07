package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("Done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return true
}

func pointer() {
	a := 2
	b := &a
	*b = 20
	fmt.Println(a)
}

func arraysAndSlices() {
	names := []string{"nico", "heesu", "querensys"}
	names = append(names, "flynn")
	fmt.Println(names)
}

func maps() {
	heesu := map[string]string{"name": "heesu", "age": "25"}
	for key, value := range heesu {
		fmt.Println(key, value)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

func printPerson() {
	favFood := []string{"burger", "chicken"}
	heesu := person{name: "heesu", age: 25, favFood: favFood}
	fmt.Println(heesu)
}

func main() {
	totalLength, up := lenAndUpper("heesu")
	fmt.Println(totalLength, up)
	fmt.Println(multiply(2, 3))
	repeatMe("heesu", "choonsik", "querensys")
	result := superAdd(1, 2, 3, 4, 5)
	fmt.Println(result)
	fmt.Println(canIDrink(25))
	pointer()
	arraysAndSlices()
	maps()
	printPerson()
}

package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	taehoon := person{"taehoon", 28, favFood}
	// taehoon := person{name: "taehoon", age: 28, favFood: favFood}
	// 2가지 방법으로 명시 가능
	fmt.Println(taehoon)
}

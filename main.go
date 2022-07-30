package main

import "fmt"

// func lenAndUpper(name string) (lenght int, uppercase string) {
// 	defer fmt.Println("I'm done")
// 	lenght = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// func main() {
// 	totalLenght, up := lenAndUpper("taehoon")
// 	fmt.Println(totalLenght, up)
// }

// func superAdd(numbers ...int) int {
// 	fmt.Println(numbers)
// 	return 1
// }

// func main() {
// 	superAdd(1, 2, 3, 4, 5, 6)
// }

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

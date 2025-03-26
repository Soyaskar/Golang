package main

import (
	"fmt"
	"log"
	"soyaskar/initialGoProject/auth"
)

func main() {
	fmt.Println("Hello world!")

	var a = 10
	log.Println("Value of a: ")
	log.Printf("a =%d", a)

	b := "Apple"
	log.Printf("b =%s", b)
	b = "Ball"
	log.Printf("b =%s", b)

	check := true
	if check {
		fmt.Println("checked")
	}
	c := 50
	if c < 50 {
		fmt.Println("c is less than 50")
	} else if c == 50 {
		fmt.Println("c is equal to 50")
	} else {
		fmt.Println("c is not less than 50")
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("Index =%d", i+1)
	}

	numbers := []int{1, 5, 6, 9}
	for index, number := range numbers {
		result := fmt.Sprintf("index = %d, number = %d", index, number)
		log.Printf("Result =%s", result)
	}
	if res, val := checkEven(23); res {
		fmt.Printf("Value=%s", val)
	} else {
		fmt.Printf("Value=%s", val)
	}

	fmt.Printf("\n")
	sum := add(19, 20)
	fmt.Println(sum)

	summ, length, msg := addd(9, 10, 50, 60, 40)
	fmt.Println(summ, length, msg)

	func() {
		fmt.Println("This is anonymous")
	}()
	display := func(name string) {
		fmt.Printf("Name: %s", name)
	}
	display("Soyuz")
	fmt.Printf("\n")

	next := sequenceGenerator()
	aa := next()
	fmt.Println(aa)
	bb := next()
	fmt.Println(bb)
	for i := 0; i < 10; i++ {
		fmt.Println(next())
	}

	resultt := calculate("multiply", 10, 10)
	fmt.Println(resultt)

	msgg := auth.Login("Soyuz")
	fmt.Println(msgg)

}

func checkEven(a int) (bool, string) {
	if a%2 == 0 {
		return true, "even"
	} else {
		return false, "odd"
	}

}

// Variadic function
// REST OPERATOR(...): combine variable parameters to an array
func add(a int, b int) int {
	return a + b
}
func addd(numbers ...int) (int, int, string) {
	fmt.Println(numbers)
	addition := 0
	for _, number := range numbers {
		addition += number
	}
	return addition, len(numbers), "Addition successful"
}

// ANONYMOUS & CLOSURE FUNCTION:
// cant be made outside main func
// closure: Capture and Retain
func sequenceGenerator() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

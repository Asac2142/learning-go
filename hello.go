package main

import (
	"fmt"
	"math"
)

func main() {
	age := 70
	id := "23rsv167-4"
	salary := 5000.75

	fmt.Printf("Value: %d \t\t Type: %T \n", age, age)
	fmt.Printf("Value: %s \t Type: %T \n", id, id)
	fmt.Printf("Value: %f \t Type %T \n\n", salary, salary)

	num1, num2, num3 := 747, 911, 90210

	fmt.Printf("Decimal: %d \t Binary: %b \t\t Hexadecimal: %x \n", num1, num1, num1)
	fmt.Printf("Decimal: %d \t Binary: %b \t\t Hexadecimal: %x \n", num2, num2, num2)
	fmt.Printf("Decimal: %d \t Binary: %b \t Hexadecimal: %x \n\n", num3, num3, num3)

	var maxInt8 int8 = math.MaxInt8
	var maxUInt8 uint8 = math.MaxUint8

	fmt.Println("Max int8: ", maxInt8)
	fmt.Println("Max uint8: ", maxUInt8)

	var p1 = person{"James"}
	fmt.Println(p1.first)
	p1.sayHello()
}

type person struct {
	first string
}

func (p person) sayHello() {
	fmt.Println("Hi, my first name is", p.first)
}

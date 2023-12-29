package main

import "fmt"

// * you can specify the data type at the end for both arguments or ...
func add(num1, num2 int) int {
	return num1 + num2
}

// * like this, specifying the data type for each
func substract(num1 int, num2 int) int {
	return num1 - num2
}

func greet(text string) {
	fmt.Println(text)
}

// * functions can return more than one value
// * NOTICE the return for swap is specifying the TYPE of return, so in this case is string, string
func swap(text1, text2 string) (string, string) {
	return text2, text1
}

/*
- Go's return values may be named. If so, they are treated as variables defined at the top of the function.
- These names should be used to document the meaning of the return values.
- A return statement without arguments returns the named return values. This is known as a "naked" return.
- Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9 // * implicitly, x will have a rounded value since the return type for "x" is "int"
	y = sum - x     // * same for "y"
	return
}

func main() {
	var result = add(3, 9)
	fmt.Printf("Add result %d \n", result)
	fmt.Printf("Substract result %d \n", substract(9, 3))
	greet("I am within a function")

	msg1, msg2 := swap("Hello", "World") // * the variables msg1 and msg2 will hold the return strings from "swap()"
	fmt.Println(msg1, msg2)

	val1, val2 := split(100)

	fmt.Println(val1, val2)

	// * iota is an identifier that starts with 0 and it is only used with constants
	// * so in this case, I am creating a constant section I think is called... and I am attaching the value of iota to each day (monday, tuesday...)
	const (
		_       = iota // * throwing away the first iota value
		Monday         // * 1
		Tuesday        // * 2 and so on...
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
		numberOfDays // * if I were to create my own module where exports the days of the week, "numberOfDays" would not be exported since starts with a lower case letter
		// * "numberOfDays" can only be accessed within the same package, kind of like having public (exported) and private (not exported)
	)

	const myIndex = iota      // * zero here
	const anotherIndex = iota // * zero here as well

	fmt.Println("Monday", Monday)
	fmt.Println("Tuesday", Tuesday)
	fmt.Println("Wednesday", Wednesday)
	fmt.Println("Thursday", Thursday)
	fmt.Println("Friday", Friday)
	fmt.Println("Saturday", Saturday)
	fmt.Println("Sunday", Sunday)
	fmt.Println("My index", myIndex)
	fmt.Println("My index", anotherIndex)
}

package main

import "fmt"

const foo = "something"

// ! this "month" variable is not valid because is using the shortcut which is only valid within functions
// month := 'May'

func main() {
	var message string = "Welcome to GO!"
	var id = 44
	age := 13 // * this is a shorcut for declaring and assigning variables but only available within functions

	fmt.Printf("Message %s \n", message)
	fmt.Println(id)
	fmt.Println(age)
	fmt.Println(foo)

	a, b, c, d, e, f := 1, 2, 3, 4, 5, 6

	// * %v default format
	// * %d base 10 format
	// * %b binary format
	// * %x hexadecimal format

	fmt.Printf("a: default: %v  base10: %d  binary: %b  hexa:%x \n", a, a, a, a)
	fmt.Printf("b: default: %v  base10: %d  binary: %b  hexa:%x \n", b, b, b, b)
	fmt.Printf("c: default: %v  base10: %d  binary: %b  hexa:%x \n", c, c, c, c)
	fmt.Printf("d: default: %v  base10: %d  binary: %b  hexa:%x \n", d, d, d, d)
	fmt.Printf("e: default: %v  base10: %d  binary: %b  hexa:%x \n", e, e, e, e)
	fmt.Printf("f: default: %v  base10: %d  binary: %b  hexa:%x \n", f, f, f, f)

	fmt.Print(`
		all
		wrong
		And I am asked—ask myself (I, too, covered   
		with the gurry of it) where
		shall we go from here, what can we do
		when even the public conveyances
		sing?
		how can we go anywhere,
		even cross-town
		how get out of anywhere (the bodies   
		all buried
		in shallow graves?
	`)

	salary := 5800.50     // GO infers salary as float64 type
	bonus := 100          // as int
	netWorth := 70000.589 // as float64

	subTotal := salary + float64(bonus) + netWorth // subtotal is also a float64, you have to convert bonus (int) to float64 since netWorth and salary are both float64
	var total float32 = float32(subTotal)          // total would be a float32 so since subTotal is float64 the convertion needs to happen

	fmt.Printf("Total: %v \n", total)
	fmt.Printf("Type of total is: %T \n", total) // * with %T we can print the type of data for total (float32 in this case)

	const hello = "Hello"
	fmt.Printf("Type of hello: %T \n", hello)
}

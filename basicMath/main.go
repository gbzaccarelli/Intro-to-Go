package main

import "fmt"

// In this problem, you will create the basicMath function that takes two integers as input parameters
// Your function should return four values: the sum, difference, product, and quotient of the two integers.
func basicMath(a int, b int) (int, int, int, float64) {

}

func main() {
	good := true
	a, b, c, d := basicMath(10, 5)
	fmt.Printf("Sum: %d, Difference: %d, Product: %d, Quotient: %.2f\n", a, b, c, d)
	if !(a == 15 && b == 5 && c == 50 && d == 2.00) {
		good = false
	}
	a, b, c, d = basicMath(20, 4)
	fmt.Printf("Sum: %d, Difference: %d, Product: %d, Quotient: %.2f\n", a, b, c, d)
	if !(a == 24 && b == 16 && c == 80 && d == 5.00) {
		good = false
	}
	if good {
		fmt.Println("Correct output. Well done!")
	} else {
		fmt.Println("Wrong output. Please try again.")
	}

}

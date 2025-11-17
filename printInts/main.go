package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// In this problem, you will modify the printInts function to print out numbers 1 to 100.
// Each number should be separated by a space, and for every 10 numbers printed out, you should make a new line.
// Example Output:
// 1 2 3 4 5 6 7 8 9 10
// 11 12 13 14 15 16 17 18 19 20
// ...
// 91 92 93 94 95 96 97 98 99 100
func printInts() {
	// Your code goes here
}

func testUserOutput(output string) {
	numbers := true
	for i := 1; i <= 100; i++ {

		if !strings.Contains(output, fmt.Sprintf("%d", i)) {
			numbers = false
			break
		}
	}

	newlines := true
	for i := 10; i <= 100; i += 10 {
		if !strings.Contains(output, fmt.Sprintf("%d", i)) {
			newlines = false
			break
		}
	}

	if numbers && newlines {
		fmt.Println("Correct output. Well done!")
	} else {
		fmt.Println("Incorrect output. Please try again.")

	}
}

func main() {
	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printInts()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = origStdout

	output := buf.String()

	fmt.Print(output)
	testUserOutput(output)
}

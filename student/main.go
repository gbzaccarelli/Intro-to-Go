package main

import "fmt"

// In this problem, you will implement functions that manipulate a Student struct using pointers.
type Student struct {
	Name  string
	Grade int
}

//Implement increaseGrade that takes a *Student
func increaseGrade(s *Student) {
	// Your code here
}

//Implement swapGrades that takes two *Student
func swapGrades(s1, s2 *Student) {
	// Your code here
}

//Implement createStudent that returns *Student
func createStudent(name string, grade int) *Student {

}

func main() {
	student1 := createStudent("Alice", 85)
	student2 := createStudent("Bob", 90)

	fmt.Printf("Before: %s (%d), %s (%d)\n",
		student1.Name, student1.Grade, student2.Name, student2.Grade)

	increaseGrade(student1)
	swapGrades(student1, student2)

	fmt.Printf("After: %s (%d), %s (%d)\n",
		student1.Name, student1.Grade, student2.Name, student2.Grade)
}

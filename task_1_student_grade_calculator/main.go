package main

import (
	"fmt"
)

var grades = map[string]int{}

func main() {
	// fmt.Println("Hello, World!")
	var studentName string
	var noOfSubjects int

	fmt.Print("Enter the your student name : ")
	fmt.Scanln(&studentName)

	fmt.Print("How many subjects do  you take ? : ")
	fmt.Scanln(&noOfSubjects)

	getSubjects(noOfSubjects)

	// getAverageGrade();
	fmt.Println("------------------------------------------------------")

	fmt.Printf("Your Student name is : %v \nNo of subjeect : %v \n", studentName, noOfSubjects)

	fmt.Println("\nYour Subjects and Grades are : ")
	for subject, grade := range grades {
		fmt.Printf("Subject : %s \t Grade : %d \n", subject, grade)
	}

	fmt.Printf("\nYour Average Grade is : %v", getAverageGrade())

	fmt.Println("\n------------------------------------------------------")

}

func getSubjects(noOfSubjects int) map[string]int {
	var subjectName string
	var grade int

	if noOfSubjects < 1 {
		fmt.Println("You have to take at least one subject")
		return nil
	}
	for i := 1; i <= noOfSubjects; i++ {
		fmt.Printf("Enter the name of subject %d : ", i)
		fmt.Scanln(&subjectName)

		fmt.Printf("Enter the grade for %s : ", subjectName)
		fmt.Scanln(&grade)

		grades[subjectName] = grade
	}
	return grades
	// fmt.Println("\n number of the subjects :", noOfSubjects)
}

func getAverageGrade() float64 {
	var count, total int

	for _, grade := range grades {
		count++
		total += grade
	}

	return float64(total) / float64(count)

}

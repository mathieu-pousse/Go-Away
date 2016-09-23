/**
 * Filename: struct.go
 * Author: Nyah Check
 * Usage: Experiments on constructors.
 * Licence: GNU PL 2016
 */
package main

import (

	"fmt"
)

type student_info struct {
	Name string
	St_ID string
	Dept string
	Age int
	height float32
}

func (s *student_info) add_name() int{
	s.Name = "Nyah Check"
	return 0
} 

func (s *student_info) add_id() int {

	s.St_ID = "FE12A025"
	s.Dept = "Computer Engineering"
	return 0
}

func (s *student_info) add_bio() int {

	s.Age = 25
	s.height = 1.92
	
	return 0
}

func (s * student_info) print_info() {

	fmt.Println("Student information: ")
	fmt.Printf("Name: %s", s.Name)
	fmt.Printf("\nID: %s", s.St_ID)
	fmt.Printf("\nDepartment: %s", s.Dept)
	fmt.Printf("\nAge: %d\t Height: %.2f\n", s.Age, s.height);

}

/**
 * Adding constructor for student_info object.
 */
func new_student(name string, id string, dept string, age int, ht float32) *student_info {

	return &student_info {
	
		Name: name,
		St_ID: id,
		Dept: dept,
		Age: age,
		height: ht,
	}

	
}
func main() {

	var student student_info
	fmt.Println("Student details before constructors.")
	student.print_info()
	student.add_name()
	student.add_id()
	student.add_bio()
	fmt.Println("Student details after constructors.")
	student.print_info()
	
	fmt.Println("New Student added with constructor.")
	new_st := new(student_info)
	new_st = new_student("Tekang Check", "SM120154", "Accounting", 23, 1.89)
	new_st.print_info();
	
}

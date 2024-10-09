package main

import "fmt"

// func main() {
// 	age := 24
// 	legalAge := 18
// 	fmt.Println("Age: ", age)

// 	//sends the copy of the age as a parameter.
// 	//the age copy is then deleted by the garbage collector
// 	// adultYears := getAdultYears(age, legalAge)
// 	// fmt.Println("Adult years: ", adultYears)

// 	var agePointer *int
// 	agePointer = &age
// 	fmt.Println("Address of the age variable: ", agePointer)
// 	fmt.Println("Value of the age variable by dereferencing: ", *agePointer)

// 	adultYears := getAdultYears(agePointer, legalAge)
// 	fmt.Println("Adult years: ", adultYears)

// }

// // func getAdultYears (currentage int, legalAge int) int {
// // 	return currentage - legalAge;
// // }

// func getAdultYears (currentage *int, legalAge int) int {
// 	return *currentage - legalAge; //get the value of the adress in the *int type currentage pointer and subtrackt legalAge from it

// }

func main() {
	var i int //creates a variable on stack named i, which is the type of int and has an assigned address. Doesn't have a value yet.
	i = 32 //the value of the i is now 32

	var p *int //creates a variable on stack named p, which is the type of int pointer and has an assigned address. Doesn't have a value yet.
	p = &i //the value of p is the address of i

	var j *int //creates a variable on stack named j, which is the type of int pointer and has an assigned address. Doesn't have a value yet.
	j = new (int) //allocates a space on heap with type int and an assigned address. then this assigned address is assigned to the value of j
	//the value of j is the newly allocated heap memory space
	*j = 10 //the value of the allocated space on heap is 10
}
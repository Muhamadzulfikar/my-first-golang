package main

import (
	"fmt"
)

func main() {
	// write in golang
	fmt.Println("Hello world")

	var tes string = "Hello world"
	fmt.Println(tes)

	// declare variables
	var nilai int = 12
	fmt.Println(nilai)

	var gaji float64 = 7.5
	fmt.Println(gaji)

	var isTrue bool = true
	fmt.Println(isTrue)

	//array asociates
	//define key and value types
	var person map[string]string = map[string]string{
		"id":   "1",
		"name": "John",
	}

	// call array asociates with key
	fmt.Println(person["id"])

	// define array
	var fruits []string = []string{"apple", "banana", "grape", "orange"}

	// call array with index array
	fmt.Println(fruits[0])

	// conditional statement
	if 12 > 10 {
		fmt.Println("betul")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("loop", i)
	}

	//defnine multiple variables
	var a, b, c, d string = "a", "b", "c", "d"
	fmt.Printf("%s, %s, %s, %s\n", a, b, c, d)

	//define variables with any type
	name := person["name"]

	fmt.Println(name)

	_ = "Golang itu mudah"

	//declare pointer
	pointer := new(string)

	// redeclarate pointer value
	*pointer = "jao"

	fmt.Println(pointer)  //print pointer address
	fmt.Println(*pointer) //print pointer value

	var decimalNumber float32 = 2.62 //declarated decimal variable with length value

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber) //print decimal with 3 value afeter comma

	const firstName string = "muhamad Zulfikar"

	const (
		myname   string = "muhamad"
		yourname string = "Zulfikar"
		myage    uint8  = 20
	)

	fmt.Print(myage)

	var myProfile map[string]int = map[string]int{
		"idpersonal":    1,
		"idtransaction": 2,
	}

	for key, val := range myProfile {
		fmt.Println(key, val)
	}

	fmt.Println(len(myProfile))

	delete(myProfile, "idpersonal")

	var month []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	for _, month := range month {
		fmt.Printf("sekarang bulan %d \n", month)
	}

	var day []string = []string{"jan", "feb", "mar", "apr", "may", "jun"}

	var newDay = day[0:2]

	fmt.Println(newDay)

	gabungan := []any{"12", 0}

	fmt.Println(gabungan)

	obj := map[string]any{
		"fruits": []string{
			"apple", "orange", "grape",
		},
		"person": []string{
			"zulfikar", "anton",
		},

	}

	newfruits, newPerson := obj["fruits"], obj["person"]

	fmt.Println(newfruits, newPerson)

	var cars []string = []string{"volvo", "BMW", "supra"}

	var newCar []string = cars[1:2]

	fmt.Println(newCar)

	for key,car := range cars {
		fmt.Println(key,car)
	}
}

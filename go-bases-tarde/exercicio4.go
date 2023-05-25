package main

import "fmt"

func Exercicio4() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])

	var count int
	for _, employee := range employees {
		if employee > 20 {
			count += 1
		}
	}
	employees["Thamyris"] = 34
	delete(employees, "Pedro")

	fmt.Println(employees)

}

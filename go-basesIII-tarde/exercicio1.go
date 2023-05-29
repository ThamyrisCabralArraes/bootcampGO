package main

import "fmt"

// type users struct {
// 	Nome      string
// 	Sobrenome string
// 	Email     string
// 	Idade     int
// 	Senha     int
// }

func exercicio1(n, s, e *string, i *int, se *int) {
	*n, *s, *e, *i, *se = "Carlos", "Arraes", "carlos@test.com", 37, 12345

	fmt.Println(*n, *s, *i, *e, *se)
}

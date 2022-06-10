package main

import "fmt"

func main() {
	var a *int // обявляем переменную А как указатель на int
	var b = 13
	a = &b // присваиваем переменной А указатель на В

	fmt.Println(*a) // разименовываем указатель А

	*a = 25 // присваиваем переменной В новое значение, через указатель А
	fmt.Println(b)
}

package main

import "fmt"

func main() {
	//Простейший вывод на консоль. Println - это вывод аргумента + `\n` (новая строка)
	fmt.Println("Hello world")
	fmt.Println("Second line")
	//Функция Print - это простой вывод аргумента
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Print("Third")
	// Форматированный вывод: Printf - стандартный вывод os.Stdout  c флагами форматирования
	fmt.Printf("\nHello, my name is %s\nMy age is %d\n", "Bob", 42)
}

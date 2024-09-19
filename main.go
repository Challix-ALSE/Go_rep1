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
	//////////////////////
	//////////////////////
	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after initialisation:", age)


	// Деклврироывние и инициализация пользовательским значением
	var height int = 183
	fmt.Println("My height is:", height)

	//В чем полустрогость типизации
	var uid = 12345
	fmt.Println("My uid:", uid)
	//Декларирование и инициализация переменных одного типа (множественный слуяай)
	var firstVar, secondVar int = 20, 30
	fmt.Printf("firstVar:%d secondVar:%d\n", firstVar, secondVar)
	// Декларирование блока переменных 
	var(
		personName string = "Bob"
		personAge int = 42
		personUid int
	)

	fmt.Printf("Name: %s\nAge %d\nUid: %d\n", personName, personAge, personUid)
}

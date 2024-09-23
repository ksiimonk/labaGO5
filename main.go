package main

import (
	"fmt"
	"math"
)

// Структура Person
type Person struct {
	name string
	age  int
}

// Метод для вывода информации о человеке
func (p Person) Info() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}

// Метод birthday для увеличения возраста на 1 год
func (p *Person) birthday() {
	p.age++
}

// Интерфейс Shape с методом Area
type Shape interface {
	Area() float64
}

// Структура Circle с полем radius
type Circle struct {
	radius float64
}

// Реализация метода Area для Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Структура Rectangle
type Rectangle struct {
	width  float64
	height float64
}

// Реализация метода Area для Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Функция для вывода площади каждого объекта
func PrintAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}

// Интерфейс Stringer
type Stringer interface {
	String() string
}

// Структура Book для хранения информации о книге
type Book struct {
	title  string
	author string
}

// Реализация метода String для Book
func (b Book) String() string {
	return fmt.Sprintf("Title: %s, Author: %s", b.title, b.author)
}

func main() {
	// Работа с Person
	person := Person{name: "Alice", age: 30}
	fmt.Println(person.Info())
	person.birthday()
	fmt.Println("After birthday:", person.Info())

	// Работа с Shape
	shapes := []Shape{
		Circle{radius: 5},
		Rectangle{width: 4, height: 6},
	}

	PrintAreas(shapes)

	// Работа с Stringer
	book := Book{title: "Go Programming", author: "John Doe"}
	fmt.Println(book.String())
}

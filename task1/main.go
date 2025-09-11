package main

import "fmt"

// Структура родителя
type Human struct {
	Name string
	Age  int
}

// Метод реализующий структуру родителя
func (h Human) speak() string {
	return fmt.Sprintf("My name is %s, I am %v years old. Yes, I can speak!", h.Name, h.Age)
}

// Метод реализующий структуру родителя
func (h Human) reader() string {
	return fmt.Sprintf("I can reader!")
}

// дочерняя стрруктура
type Action struct {
	Human
	Profession string
}

// Метод реализующий дочернюю структуру
func (a Action) work() string {
	return fmt.Sprintf("I'm %s, my work as a %s!", a.Name, a.Profession)
}

// переопределение родительского метода, на дочерний метод, если это нужно
func (a Action) reader() string {
	return fmt.Sprintf("I (%s) can reader books!", a.Name)
}

func main() {
	// экземпляр родительской структуры
	helen := Human{
		Name: "Helen",
		Age:  33,
	}

	// экземпляр дочерней структуры
	dmitrii := Action{
		Human: Human{
			Name: "Dmitrii",
			Age:  35,
		},
		Profession: "programmer",
	}

	//методы родительской структуры
	fmt.Println(helen.speak())
	fmt.Println(helen.reader())

	//собственный метод дочерней структуры
	fmt.Println(dmitrii.work())

	// методы родителя для дочерней структуры
	fmt.Println(dmitrii.speak())
	fmt.Println(dmitrii.reader()) // данный метод переопределен для дочернего типа, но можно убрать

}

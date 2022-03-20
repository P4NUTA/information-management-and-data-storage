package main

import (
	"fmt"
	"lab2/stack"
)

func main() {
	// Задача: Реализовать 2 стека и поменять их местами

	// Демонстрация работы стека
	fmt.Println("Демонстрация работы стека")
	stack1 := stack.New()
	stack1.Push("1")
	stack1.Push("2")
	stack1.Push("3")

	fmt.Println(stack1.Pop())
	fmt.Println(stack1.Pop())
	fmt.Println(stack1.Pop())

	fmt.Println("Начальные условия")
	stack1.Push(1.2)
	stack1.Push(2.3)
	stack1.Push(3.4)

	stack2 := stack.New()
	stack2.Push("4")
	stack2.Push("5")
	stack2.Push("6")

	fmt.Print("stack 1 = ")
	stack1.Print()
	fmt.Print("stack 2 = ")
	stack2.Print()

	// Способ 1
	stack1, stack2 = stack2, stack1
	fmt.Println("Поменяли первым способом")
	fmt.Print("stack 1 = ")
	stack1.Print()
	fmt.Print("stack 2 = ")
	stack2.Print()

	// Способ 2
	stack1, stack2 = stack.ChangeStacksByReturn(stack1, stack2)
	fmt.Println("Поменяли вторым способом")
	fmt.Print("stack 1 = ")
	stack1.Print()
	fmt.Print("stack 2 = ")
	stack2.Print()

	// Способ 3
	stack1, stack2 = stack.ChangeByPop(stack1, stack2)
	fmt.Println("Поменяли третьим способом")
	fmt.Print("stack 1 = ")
	stack1.Print()
	fmt.Print("stack 2 = ")
	stack2.Print()
}

package stack

import "fmt"

type Stack []any

// New Создание стека
func New() Stack {
	return Stack{}
}

// Push добавление элемента в стек
func (s *Stack) Push(text any) {
	*s = append(*s, text)
}

//Pop удаление последнего элемента из стека
func (s *Stack) Pop() any {
	if s.Empty() {
		return ""
	}
	length := len(*s)
	text := (*s)[length-1]
	*s = (*s)[:length-1]
	return text
}

// Print Отображение всех элементов в стеке
func (s Stack) Print() {
	fmt.Println(s)
}

// Empty Проверяет стек на пустоту
func (s Stack) Empty() bool {
	return len(s) == 0
}

// ChangeStacksByReturn поменять местами с помощью функции
func ChangeStacksByReturn(stack1, stack2 Stack) (Stack, Stack) {
	return stack2, stack1
}

// ChangeByPop поменять местами с помощью функции Pop()
func ChangeByPop(stack1, stack2 Stack) (Stack, Stack) {
	tempStack := New()
	for len(stack1) > 0 {
		tempStack.Push(stack1.Pop())
	}

	for len(stack2) > 0 {
		stack1.Push(stack2.Pop())
	}

	for len(tempStack) > 0 {
		stack2.Push(tempStack.Pop())
	}

	return stack1, stack2
}

package main

import "fmt"

func removeDuplicates(s string) string {
	var stack = new(Stack)
	for _, value := range s {
		if stack.IsEmpty() || int(value) != stack.Peek() {
			stack.Push(int(value))
		} else {
			stack.Pop()
		}
	}

	var result = ""
	for !stack.IsEmpty() {
		result = string(stack.Pop()) + result
	}

	return result
}

func main() {
	fmt.Print(removeDuplicates("abbaca"))
}

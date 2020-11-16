package main

import "fmt"

/**
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
通过次数458,074提交次数1,058,787
*/
func main() {
	fmt.Println(isOk("()"))
	fmt.Println(isOk("()[]{}"))
	fmt.Println(isOk("(]"))
	fmt.Println(isOk("([)]"))
	fmt.Println(isOk("()"))
	fmt.Println(isOk("()"))

}

func isOk(needValidateString string) bool {
	var stack = new(Stack)
	for _, value := range needValidateString {
		if value == '{' {
			stack.Push('}')
		} else if value == '(' {
			stack.Push(')')
		} else if value == '[' {
			stack.Push(']')
		} else if !stack.IsEmpty() && int(value) != stack.Pop() {
			return false
		}
	}
	return stack.IsEmpty()
}

type Stack struct {
	index int
	data  [16]int
}

func (s *Stack) Push(k int) bool {
	if s.index == len(s.data) {
		return false
	}
	s.data[s.index] = k
	s.index++
	return true
}

func (s *Stack) Pop() (ret int) {
	if s.IsEmpty() {
		return -1
	}
	s.index--
	return s.data[s.index]
}

func (s *Stack) Peek() (ret int) {
	if s.IsEmpty() {
		return -1
	}
	tempIndex := s.index
	tempIndex--
	return s.data[tempIndex]
}

func (s *Stack) IsEmpty() bool {
	return s.index <= 0
}

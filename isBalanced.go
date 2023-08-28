package main

import "fmt"

func isBalanced(staples string) bool {
	var stack []string

	for _, ch := range staples {
		switch ch {
		case '(':
			stack = append(stack, ")")
		case '{':
			stack = append(stack, "}")
		case '[':
			stack = append(stack, "]")
		case ')', '}', ']':
			//	s := len(stack) - 1
			s := stack[len(stack)-1]
			if string(ch) == s {
				stack = stack[:len(stack)-1]
			} else {
				fmt.Println("Строка не сбалансирована")
				return false
			}
		}
	}

	if len(stack) == 0 {
		fmt.Println("Строка сбалансирована")
		return true
	} else {
		fmt.Println("Строка не сбалансирована")
		return false
	}
}

func main() {
	var intro string
	intro = ""
	fmt.Println("Введите строку со скобками:")
	fmt.Scan(&intro)
	isBalanced(intro)

}

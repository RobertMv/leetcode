package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("]"))      // false
}

func isValid(s string) bool {
	closeToOpen := map[byte]byte{')': '(', ']': '[', '}': '{'}
	var stack stack
	for _, c := range s {
		if v, ok := closeToOpen[byte(c)]; ok {
			if stack.last() == v {
				stack.pop()
			} else {
				return false
			}
		} else {
			stack.push(byte(c))
		}
	}
	return len(stack) == 0
}

type stack []byte

func (s *stack) push(b byte) {
	*s = append(*s, b)
}

func (s *stack) pop() byte {
	idx := len(*s) - 1
	el := (*s)[idx]
	*s = (*s)[:idx]
	return el
}

func (s *stack) last() byte {
	if len(*s) == 0 {
		return 0
	}
	idx := len(*s) - 1
	el := (*s)[idx]
	return el
}

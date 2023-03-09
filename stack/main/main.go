package main

import (
	"data-structure/stack"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	expr := infix2Postfix("((40 / (7 / 2)) * 1.25) - (2 + (1 + 1))")
	fmt.Println(expr)
	fmt.Printf("%.2f", evalPostfix(expr))
}

func bracketCheck(str string) bool {
	rs := []rune(str)
	s := stack.NewSequenceStack[rune]()
	for i := 0; i < len(str); i++ {
		if rs[i] == '(' || rs[i] == '[' || rs[i] == '{' {
			s.Push(rs[i])
		}
		if rs[i] == ')' || rs[i] == ']' || rs[i] == '}' {
			if s.IsEmpty() {
				return false
			}
			top, _ := s.Pop()
			if rs[i] == ')' && top != '(' {
				return false
			}
			if rs[i] == ']' && top != '[' {
				return false
			}
			if rs[i] == '}' && top != '{' {
				return false
			}
		}
	}
	return s.IsEmpty()
}

func isInfixExpr(expr string) bool {
	matched, _ := regexp.MatchString("[-+*/]", expr)
	matched1, _ := regexp.MatchString("\\d$|\\)$", expr)
	if !matched || !matched1 {
		return false
	}
	s := stack.NewSequenceStack[rune]()
	bs := stack.NewSequenceStack[rune]()
	exprRune := []rune(expr)
	first := exprRune[0]
	s.Push(first)
	if first == '(' {
		bs.Push(first)
	}
	for i := 1; i < len(expr); i++ {
		switch exprRune[i] {
		case ' ':
			continue
		case '+':
			pre := s.GetTop()
			if !unicode.IsNumber(pre) && pre != ')' {
				return false
			}
		case '-':
			pre := s.GetTop()
			if !unicode.IsNumber(pre) && pre != ')' {
				return false
			}
		case '*':
			pre := s.GetTop()
			if !unicode.IsNumber(pre) && pre != ')' {
				return false
			}
		case '/':
			pre := s.GetTop()
			if !unicode.IsNumber(pre) && pre != ')' {
				return false
			}
		case '(':
			pre := s.GetTop()
			if unicode.IsNumber(pre) || pre == ')' {
				return false
			}
			bs.Push(exprRune[i])
		case ')':
			pop, err := bs.Pop()
			if pop != '(' || err != nil {
				return false
			}
			pre := s.GetTop()
			if !unicode.IsNumber(pre) && pre != ')' {
				return false
			}
		case '.':
			if !unicode.IsNumber(s.GetTop()) {
				return false
			}
			matched, _ := regexp.MatchString("\\.\\d+$", string(s.Slice))
			if matched {
				return false
			}
		default:
			if unicode.IsNumber(exprRune[i]) {
				if s.GetTop() == ')' {
					return false
				}
			} else {
				return false
			}

		}
		s.Push(exprRune[i])
	}
	if !bs.IsEmpty() {
		return false
	}
	return true
}

func subNum(expr string) (numStr string, num float64) {
	for _, v := range expr {
		if unicode.IsNumber(v) || v == '.' {
			numStr += string(v)
		} else {
			break
		}
	}
	var err error
	if num, err = strconv.ParseFloat(numStr, 2); err != nil {
		numStr = ""
	}
	return
}

func infix2Postfix(expr string) (result string) {
	if !isInfixExpr(expr) {
		return
	}
	s := stack.NewSequenceStack[rune]()
	for expr != "" {
		switch expr[0] {
		case '(':
			s.Push(rune(expr[0]))
			expr = expr[1:]
		case ')':
			for !s.IsEmpty() {
				pop, _ := s.Pop()
				if pop == '(' {
					break
				}
				result += string(pop)
				result += " "
			}
			expr = expr[1:]
		case '+':
			for top := s.GetTop(); top != '(' && !s.IsEmpty(); {
				pop, _ := s.Pop()
				result += string(pop)
				result += " "
			}
			s.Push(rune(expr[0]))
			expr = expr[1:]
		case '-':
			for top := s.GetTop(); top != '(' && !s.IsEmpty(); {
				pop, _ := s.Pop()
				result += string(pop)
				result += " "
			}
			s.Push(rune(expr[0]))
			expr = expr[1:]
		case '*':
			for top := s.GetTop(); (top == '*' || top == '/') && !s.IsEmpty(); {
				pop, _ := s.Pop()
				result += string(pop)
				result += " "
			}
			s.Push(rune(expr[0]))
			expr = expr[1:]
		case '/':
			for top := s.GetTop(); (top == '*' || top == '/') && s.IsEmpty(); {
				pop, _ := s.Pop()
				result += string(pop)
				result += " "
			}
			s.Push(rune(expr[0]))
			expr = expr[1:]
		case ' ':
			expr = expr[1:]
		default:
			if numStr, _ := subNum(expr); numStr != "" {
				result += numStr
				result += " "
				expr = strings.TrimPrefix(expr, numStr)
			} else {
				return ""
			}
		}
	}
	for !s.IsEmpty() {
		pop, _ := s.Pop()
		if pop != '(' {
			result += string(pop)
			result += " "
		}
	}
	return
}

func evalPostfix(expr string) float64 {
	s := stack.NewSequenceStack[float64]()
	for expr != "" {
		switch expr[0] {
		case '+':
			post, _ := s.Pop()
			pre, _ := s.Pop()
			s.Push(pre + post)
			expr = expr[1:]
		case '-':
			post, _ := s.Pop()
			pre, _ := s.Pop()
			s.Push(pre - post)
			expr = expr[1:]
		case '*':
			post, _ := s.Pop()
			pre, _ := s.Pop()
			s.Push(pre * post)
			expr = expr[1:]
		case '/':
			post, _ := s.Pop()
			pre, _ := s.Pop()
			s.Push(pre / post)
			expr = expr[1:]
		case ' ':
			expr = expr[1:]
		default:
			numStr, num := subNum(expr)
			expr = strings.TrimPrefix(expr, numStr)
			s.Push(num)
		}
	}
	pop, _ := s.Pop()
	return pop
}

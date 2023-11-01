package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	data []interface{}
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(val interface{}) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() interface{} {
	if s.isEmpty() {
		return nil
	}
	lastIndex := len(s.data) - 1
	lastData := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return lastData
}

func (s *Stack) Peek() interface{} {
	if s.isEmpty() {
		return nil
	}
	lastIndex := len(s.data) - 1
	lastData := s.data[lastIndex]
	return lastData
}

func (s *Stack) Reverse() {
	for i, j := 0, len(s.data)-1; i < j; i, j = i+1, j-1 {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
}

func (s *Stack) Swap() {
	lastIndex := len(s.data) - 1
	s.data[lastIndex], s.data[lastIndex-1] = s.data[lastIndex-1], s.data[lastIndex]
}

func showError(err error) {
	if err != nil {
		fmt.Println("[Invalid Command]", err)
	}
}

func main() {
	dc := new(Stack)
	reader := bufio.NewReader(os.Stdin)

	for true {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}
		switch input {
		// Arithmetic operations
		case "+":
			x, y := dc.Pop().(float64), dc.Pop().(float64)
			dc.Push(x + y)
		case "-":
			x, y := dc.Pop().(float64), dc.Pop().(float64)
			dc.Push(y - x)
		case "*":
			x, y := dc.Pop().(float64), dc.Pop().(float64)
			dc.Push(x * y)
		case "/":
			x, y := dc.Pop().(float64), dc.Pop().(float64)
			dc.Push(y / x)
		case "%":
			x, y := dc.Pop().(float64), dc.Pop().(float64)
			dc.Push(math.Mod(y, x))
		case "~":
			x, y := dc.Pop().(float64), dc.Pop().(float64)
			dc.Push(y / x)
			dc.Push(math.Mod(y, x))
		case "^":
			x, y := dc.Pop().(float64), dc.Pop().(float64)
			dc.Push(math.Pow(y, x))
		case "v":
			x := dc.Pop().(float64)
			dc.Push(math.Sqrt(x))

			// Printing Commands
		case "f":
			fmt.Println(dc.data...)
		case "p":
			fmt.Println(dc.Peek())
		case "n":
			fmt.Println(dc.Pop())
		case "q":
			os.Exit(0)

			// stack control
		case "c":
			dc.data = nil
		case "d":
			dc.Push(dc.Peek())
		case "r":
			dc.Swap()
		case "R":
			dc.Reverse()

			// default case (push the number to the stack)
		default:
			num, err := strconv.ParseFloat(input, 64)
			if err != nil {
				showError(err)
			} else {
				dc.Push(num)
			}
		}
	}
}

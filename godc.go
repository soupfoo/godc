package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/exec"
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
	if len(s.data) > 1 {
		lastIndex := len(s.data) - 1
		s.data[lastIndex], s.data[lastIndex-1] = s.data[lastIndex-1], s.data[lastIndex]
	} else {
		fmt.Println("[Not enough items to swap]")
	}
}

func showError(err error) {
	if err != nil {
		fmt.Println("[Invalid Command]", err)
	}
}

func msg() {
	fmt.Println("[Not enough operands]")
}

func main() {
	dc := new(Stack)
	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print(">  ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		expression := strings.Split(input, " ")

		for _, element := range expression {
			switch {
			// Arithmetic operations
			case element == "+":
				if len(dc.data) > 1 {
					x, y := dc.Pop().(float64), dc.Pop().(float64)
					dc.Push(x + y)
				} else {
					msg()
				}
			case element == "-":
				if len(dc.data) > 1 {
					x, y := dc.Pop().(float64), dc.Pop().(float64)
					dc.Push(y - x)
				} else {
					msg()
				}
			case element == "*":
				if len(dc.data) > 1 {
					x, y := dc.Pop().(float64), dc.Pop().(float64)
					dc.Push(x * y)
				} else {
					msg()
				}
			case element == "/":
				if len(dc.data) > 1 {
					x, y := dc.Pop().(float64), dc.Pop().(float64)
					dc.Push(y / x)
				} else {
					msg()
				}
			case element == "%":
				if len(dc.data) > 1 {
					x, y := dc.Pop().(float64), dc.Pop().(float64)
					dc.Push(math.Mod(y, x))
				} else {
					msg()
				}
			case element == "~":
				if len(dc.data) > 1 {
					x, y := dc.Pop().(float64), dc.Pop().(float64)
					dc.Push(y / x)
					dc.Push(math.Mod(y, x))
				} else {
					msg()
				}
			case element == "^":
				if len(dc.data) > 1 {
					x, y := dc.Pop().(float64), dc.Pop().(float64)
					dc.Push(math.Pow(y, x))
				} else {
					msg()
				}
			case element == "v":
				if len(dc.data) > 0 {
					x := dc.Pop().(float64)
					dc.Push(math.Sqrt(x))
				} else {
					msg()
				}

				// Printing commands
			case element == "f":
				fmt.Println(dc.data...)
			case element == "p":
				fmt.Println(dc.Peek())
			case element == "n":
				fmt.Println(dc.Pop())

				// Stack control
			case element == "c":
				dc.data = nil
			case element == "d":
				dc.Push(dc.Peek())
			case element == "r":
				dc.Swap()
			case element == "R":
				dc.Reverse()
			case element == "z":
				dc.Push(len(dc.data))

				// Miscellaneous
			case element == "#" || strings.HasPrefix(element, "#"):
				continue
			case element == "clear":
				cmd := exec.Command("clear")
				cmd.Stdout = os.Stdout
				cmd.Run()
			case element == "q":
				os.Exit(0)

				// Default case (push the number to the stack)
			default:
				num, err := strconv.ParseFloat(element, 64)
				if err != nil {
					showError(err)
				} else {
					dc.Push(num)
				}
			}
		}
	}
}

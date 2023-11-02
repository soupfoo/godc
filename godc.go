package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
)

type Stack struct {
	data []any
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(val any) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() any {
	if s.isEmpty() {
		return nil
	}
	lastIndex := len(s.data) - 1
	lastData := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return lastData
}

func (s *Stack) Peek() any {
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

// Go v1.21 introduced slices pkg
// slices.Index() can replace this function
func elementIndex(xs []string, s string) int {
	for i, value := range xs {
		if value == s {
			return i
		}
	}
	return -1
}

func showError(err error) {
	if err != nil {
		fmt.Println("[Invalid Command]", err)
	}
}

func msg() {
	fmt.Println("[Not enough operands]")
}

func help() {
	fmt.Println(`(godc)	
<number> : pushes a number to the stack

[Arithematic operations]
+,-,*,/ : sum, difference, product, quotient
% : remainder
~ : quotient and remainder
^ : exponentiation
| : modular exponentiation
v : square root

[Printing commands]
f : prints the entire contents of the stack
p : prints the value at the top of the stack
n : pops the value at the top of the stack and prints it

[Stack control commands]
c : clears the stack
d : duplicates the value at the top of the stack
r : swaps top two values
R : inverts the entire stack
z : pushes the length of the stack

[Registers]
sx : pops the value at the top and saves it in register x
lx : pushes the value of x

[Strings/Macros]
Macros can be implemented by storing strings in register.
Anything between [ and ] is a string.
e.g. [ 2 ^ ] sm
     this expression stores the string "2 ^" in register "m"
x : executes a macro
Q : quits a macro

[Miscellaneous]
help : shows help text
#    : comment line
clear: clears the screen
q    : quits godc
 `)
}

func main() {
	dc := new(Stack)                 // main stack
	register := make(map[string]any) // register

	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print(">  ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		expression := strings.Split(input, " ")
		var element string

		for i := 0; i < len(expression); i++ {
			element = expression[i]

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
			case element == "|":
				if len(dc.data) > 1 {
					x, y, z := dc.Pop().(float64), dc.Pop().(float64), dc.Pop().(float64)
					dc.Push(math.Mod(math.Pow(z, y), x))
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

			// Registers
			case strings.HasPrefix(element, "s"):
				if len(element) < 2 {
					fmt.Println("[invalid command]")
					break
				}
				value := dc.Pop()
				register[string(element[1])] = value
			case strings.HasPrefix(element, "l"):
				registerKey := string(element[1])
				registerValue := register[registerKey]
				if registerValue != nil {
					dc.data = append(dc.data, registerValue)
				} else {
					dc.data = append(dc.data, 0)
				}

			// Macros
			case element == "x":
				if dc.isEmpty() {
					fmt.Println("[Empty stack]")
					break
				}
				stackLen := len(dc.data)

				elementType := reflect.TypeOf(dc.data[stackLen-1]).Kind()

				if elementType == reflect.Slice {
					var temp any
					temp = dc.data[stackLen-1]
					expression = temp.([]string)
					dc.data = dc.data[:stackLen-1]
					i = -1
					continue
				} else {
					fmt.Println("[Invalid macro]")
				}
			case element == "Q":
				break

			// Strings
			case element == "[":
				endStr := elementIndex(expression, "]")
				if endStr == -1 {
					fmt.Println("[Invalid input : end of string not found]")
					break
				}
				stringExpression := expression[i+1 : endStr]
				if (len(expression) > endStr+1) && strings.HasPrefix(expression[endStr+1], "s") {
					macroKey := string(expression[endStr+1][1])
					register[macroKey] = stringExpression
					i = endStr + 1
				} else {
					expression = stringExpression
					i = -1
				}
			case element == "]":
				continue

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
			case element == "help":
				help()
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

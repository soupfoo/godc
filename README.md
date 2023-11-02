# godc
A reverse-polish desk calculator

---
godc is a [dc](https://en.wikipedia.org/wiki/Dc_(computer_program)) clone written in go. dc is one of the oldest surviving unix program, written even before the invention of C programming language.

# Commands
\<number> : pushes a number to the stack

## [Arithematic operations]
- +,-,*,/ : sum, difference, product, quotient
- % : remainder
- ~ : quotient and remainder
- ^ : exponentiation
- | : modular exponentiation
- v : square root

## [Printing commands]
- f : prints the entire contents of the stack
- p : prints the value at the top of the stack
- n : pops the value at the top of the stack and prints it

## [Stack control commands]
- c : clears the stack
- d : duplicates the value at the top of the stack
- r : swaps top two values
- R : inverts the entire stack
- z : pushes the length of the stack

## [Registers]
- sx : pops the value at the top and saves it in register x
- lx : pushes the value of x

## [Strings/Macros]
Macros can be implemented by storing strings in register.
Anything between [ and ] is a string.

e.g. `[ 2 ^ ] sm` stores the string "2 ^" in register "m"

- x : executes a macro
- Q : quits a macro

## [Miscellaneous]
- help : shows help text
- \#    : comment line
- clear: clears the screen
- q    : quits godc

# License
godc is licensed under the MIT license. See [LICENSE](./LICENSE) for details.
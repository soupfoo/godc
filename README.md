
# godc

> A reverse-polish desk calculator

godc is a [dc](https://en.wikipedia.org/wiki/Dc_(computer_program)) clone written in go. dc is one of the oldest surviving unix program, written even before the invention of C programming language.

## Installation

**Build from source:**

Required: Go compiler (>=v1.21)

```
git clone https://github.com/soupfoo/godc.git
cd godc
go build
./godc
```

You can download a prebuilt binary from [releases](https://github.com/soupfoo/godc/releases) page as well.

## Usage

godc uses postfix notation which means operators come after operands.

examples:
```
sum of 2 and 4

2 4 + p
```

```
evaluate (a + b)^2 ; a = 5, b = 7

5 7 sb sa la lb + 2 ^ p
```

```
quadratic roots; a = 2, b = 12, c = 1

2 12 1 sc sb sa lb d * 4 la lc * * - v sd
lb -1 * sx 2 la * sy
lx ld + ly / lx ld - ly / f
```

## Commands

```
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

[Conditionals]
=, !=, >, >=, <, <=
usage: =m runs the macro m if top two elements of the stack are equal.

[Miscellaneous]
help : shows help text
#    : comment line
clear: clears the screen
q    : quits godc
```

## Limitations

- Certain dc commands like `P,S,:,;` not implemented
- Precision control not implemented
- Script mode `-e` and file mode `-f` not implemented
- A register can store only one value at a time

## Known bugs

- Entering more than one strings in a single line causes runtime error

## License

godc is licensed under the MIT license. See [LICENSE](./LICENSE) for details.

## Acknowledgement

[MetaStag's godc](https://github.com/MetaStag/godc)


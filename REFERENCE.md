# Go run commands

Run these commands from the project root. The examples use the current program in `0.hello_world/hello.go`.

```sh
# Check the installed Go version
go version

# Run the program
go run ./0.hello_world/hello.go

# Build an executable in the project root
go build -o hello ./0.hello_world/hello.go

# Format the source file
gofmt -w ./0.hello_world/hello.go

# Build an executable program
go build hello.go

# Run an executable program
./hello.exe
```

On Windows, the built executable is named `hello.exe` and can be run with `./hello.exe` in PowerShell (or `hello.exe` in Command Prompt).

## Go topics

### For loops

Go uses `for` for repetition. Its three-part form (`for init; condition; post`) works like a counter loop, while `for condition` works like a while loop. `for range` iterates over collections such as slices, arrays, and maps. In the fruit example, `range` supplies each item's index and value.

### If/else

Go conditions do not need parentheses. An `if` can declare a short-lived variable in its initializer (`if num := 9; ...`); that variable is available throughout the related `if`, `else if`, and `else` branches. The example checks the number and prints whether it is negative, one digit, or multiple digits.

### Switch

A `switch` selects a matching `case` and can list multiple values in one case. The day example distinguishes weekends, Monday, and other weekdays. Cases do not fall through automatically; the `fallthrough` keyword explicitly continues into the next case body. The role example uses it to print the permissions available to an admin.

### Basic data types

Go variables can be declared with an explicit type using `var`, as in the example's `string`, `int`, and `bool` values. Go also has signed and unsigned integer types, floating-point types (`float32` and `float64`), and complex types (`complex64` and `complex128`).

## Functions

### Basic function

Declare a function with `func`, its name, parameter list, and optional result type. Call it by name:

```go
func greet() {
	fmt.Println("Hello!")
}

greet()
```

### Function parameters

Parameters are named and typed. Parameters with the same type can share a type declaration:

```go
func add(a, b int) int {
	return a + b
}
```

### Return values

Specify a result type after the parameter list and return a value with `return`:

```go
func square(n int) int {
	return n * n
}
```

### Multiple returned values

A function can return more than one value. Callers can assign each result, or use `_` to ignore one:

```go
func divide(a, b int) (int, int) {
	return a / b, a % b
}

quotient, remainder := divide(10, 3)
```

### Named return values

Return values can be named in the function signature. Those names act as local variables, and a bare `return` returns their current values:

```go
func split(total int) (whole, half int) {
	whole = total / 2
	half = total - whole
	return
}
```

### Variadic functions

A variadic parameter (`...Type`) accepts zero or more values of that type. Inside the function it behaves like a slice, and it must be the final parameter:

```go
func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

result := sum(1, 2, 3)
```

### Anonymous functions

An anonymous function has no declared name. It can be assigned to a variable and called through that variable:

```go
double := func(n int) int {
	return n * 2
}

result := double(4)
```

### Closures

A closure is a function value that can use variables from its surrounding scope. The returned function below captures `next` and retains its value between calls:

```go
func makeCounter() func() int {
	next := 0
	return func() int {
		next++
		return next
	}
}

counter := makeCounter()
first := counter()  // 1
second := counter() // 2
```

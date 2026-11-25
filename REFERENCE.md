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

### Arrays

An array stores a fixed number of elements of one type. Its length is part of its type, and its elements can be accessed by zero-based index. Declare and initialize an array with its length and element type:

```go
numbers := [3]int{1, 3, 5}
fmt.Println(numbers[0]) // 1
```

Use a `for` loop with an index or `for range` to visit its elements. Arrays can also be declared with `var numbers [5]int`; the elements are initialized to the type's zero value.

### Slices

A slice is a flexible view over an underlying array. Unlike an array, a slice's length can grow when elements are appended. Create one with a literal or `make`; `len` reports its current length and `cap` reports how much it can grow before needing a larger underlying array:

```go
numbers := make([]int, 3, 5) // length 3, capacity 5
numbers = append(numbers, 10, 20)
fmt.Println(len(numbers), cap(numbers)) // 5 5
```

Use a slice expression such as `numbers[1:3]` to get the elements from index 1 up to, but not including, index 3. Slices can be traversed with `for range`.

Assigning one slice to another copies its slice header, so both may refer to the same underlying array. Use the built-in `copy(destination, source)` to copy elements into a separately allocated slice:

```go
source := []int{10, 20, 30}
destination := make([]int, len(source))
copy(destination, source)
```

### Maps

A map stores values under keys. Its type is written `map[KeyType]ValueType`; initialize it with a map literal or `make` before adding entries. A nil map can be read, but assigning to it causes a runtime panic:

```go
user := map[string]string{
	"name": "Soe",
	"city": "Yangon",
}

user["email"] = "soe@example.com" // add a key
user["city"] = "Mandalay"        // update a key
fmt.Println(user["name"])

scores := make(map[string]int)
scores["Go"] = 10
delete(scores, "Go")
```

Use `for range` to iterate over a map. Map iteration order is unspecified:

```go
for key, value := range user {
	fmt.Println(key, value)
}
```

### Comma-ok idiom

Reading a map key returns its value. To distinguish a missing key from a key holding the value type's zero value, use the two-result form: the second result is `true` when the key exists and `false` otherwise. This pattern is called the comma-ok idiom:

```go
value, ok := user["name"]
if ok {
	fmt.Println("Name:", value)
} else {
	fmt.Println("Name not found")
}
```

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

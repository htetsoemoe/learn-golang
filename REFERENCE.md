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

# Golang Practice CLI

This repository contains three sections for practicing Golang:

- **DSA**: Contains implementations of Data Structures and Algorithms (DSA) problems.
- **Leetcode**: Includes solutions to Leetcode problems.
- **Pattern**: Focuses on solving pattern-based problems.

Each section is structured as a Cobra CLI application, allowing execution of different commands.

## Project Structure

```
root/
│── dsa/
│   ├── main.go
│   ├── cmd/
│   ├── go.mod
│── leetcode/
│   ├── main.go
│   ├── cmd/
│   ├── go.mod
│── pattern/
│   ├── main.go
│   ├── cmd/
│   ├── go.mod
│── README.md
│── LICENSE
```

## Installation

Ensure you have **Golang** installed (Go 1.18 or later).

Clone the repository:
```sh
$ git clone <repository-url>
$ cd <repository-folder>
```

## Running Each CLI

Each section (DSA, Leetcode, Pattern) has its own **Cobra CLI**, and they can be executed separately.

### DSA CLI
```sh
$ cd dsa
$ go run main.go
```
You can extend the CLI by adding new commands in the `cmd/` folder.

### Leetcode CLI
```sh
$ cd leetcode
$ go run main.go
```

### Pattern CLI
```sh
$ cd pattern
$ go run main.go
```

## Adding New Commands
Each CLI follows Cobra's command structure. To add a new command:
1. Navigate to the respective section (`dsa`, `leetcode`, or `pattern`).
2. Inside `cmd/`, create a new Go file (e.g., `example.go`).
3. Register the command in `cmd/root.go`.
4. Build and test the CLI.

Example command file (`cmd/example.go`):
```go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var exampleCmd = &cobra.Command{
    Use:   "example",
    Short: "An example command",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("This is an example command")
    },
}

func init() {
    rootCmd.AddCommand(exampleCmd)
}
```

## Contributing
Feel free to contribute by adding new commands or optimizing existing ones. Follow standard Go conventions and use `golangci-lint` for code quality.

## License
This is free and unencumbered software released into the public domain.

